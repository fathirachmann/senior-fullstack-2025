package main

import (
	"senior-fullstack-2025/config"
	"senior-fullstack-2025/internal/handler"
	"senior-fullstack-2025/internal/middleware"
	"senior-fullstack-2025/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()
	redisClient := config.NewRedisConfig(cfg.RedisAddr, cfg.RedisPassword)
	qDrantClient := config.NewQDrantConfig(cfg.QDrantURI, cfg.QDrantAPIKey, cfg.QDrantPort)

	vectorService := service.NewVectorService(qDrantClient)
	vectorHandler := handler.NewVectorHandler(vectorService)
	authService := service.NewAuthService(redisClient, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(authService)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to the Senior Fullstack 2025 API")
	})

	// Auth Routes
	authRoute := app.Group("/auth")
	authRoute.Post("/register", func(c *fiber.Ctx) error {
		return authHandler.RegisterUser(c)
	})
	authRoute.Post("/login", func(c *fiber.Ctx) error {
		return authHandler.LoginUser(c)
	})

	// Vector Routes
	vectorRoute := app.Group("/vector")
	vectorRoute.Use(middleware.Authentication(cfg.JWTSecret))
	vectorRoute.Post("/init", func(c *fiber.Ctx) error {
		return vectorHandler.InitData(c)
	})
	vectorRoute.Post("/search", func(c *fiber.Ctx) error {
		return vectorHandler.Search(c)
	})
	vectorRoute.Post("/seed", func(c *fiber.Ctx) error {
		return vectorHandler.BulkInsert(c)
	})

	// The Routes can be made modular but for simplicity, they are defined inline here.

	app.Listen(":3000")
}
