package handler

import (
	"fmt"
	"senior-fullstack-2025/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	Auth *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		Auth: authService,
	}
}

func (h *AuthHandler) RegisterUser(c *fiber.Ctx) error {
	type RegisterRequest struct {
		Username string `json:"username"`
		RealName string `json:"realname"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Invalid request body",
			"detail": err.Error(),
		})
	}

	err := h.Auth.RegisterUser(req.Username, req.RealName, req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  "Failed to register user",
			"detail": err.Error(),
		})
	}

	message := fmt.Sprintf("User %s registered successfully", req.RealName)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": message,
	})
}

func (h *AuthHandler) LoginUser(c *fiber.Ctx) error {
	type LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Invalid request body",
			"detail": err.Error(),
		})
	}

	token, err := h.Auth.LoginUser(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
