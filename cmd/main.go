package main

import (
	"senior-fullstack-2025/config"
	"senior-fullstack-2025/internal/handler"
	"senior-fullstack-2025/internal/service"

	"github.com/gofiber/fiber"
)

func main() {
	cfg := config.LoadConfig()
	redisClient := config.NewRedisConfig(cfg.RedisAddr, cfg.RedisPassword)

	authService := service.NewAuthService(redisClient, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(authService)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	// Auth Routes
	authRoute := app.Group("/auth")
	authRoute.Post("/register", func(c *fiber.Ctx) {
		authHandler.RegisterUser(c)
	})
	authRoute.Post("/login", func(c *fiber.Ctx) {
		authHandler.LoginUser(c)
	})

	app.Listen(":3000")
}
