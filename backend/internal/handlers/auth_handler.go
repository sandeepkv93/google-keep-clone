package handlers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/google/uuid"
    "google-keep-clone/backend/internal/services"
    "google-keep-clone/backend/internal/validators"
)

type AuthHandler struct {
    authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
    return &AuthHandler{authService: authService}
}

// @Summary Register user
// @Description Register a new user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body validators.RegisterRequest true "User registration data"
// @Success 201 {object} map[string]interface{}
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
    var req validators.RegisterRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
    }

    if err := validators.ValidateStruct(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    user, token, err := h.authService.Register(&req)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(201).JSON(fiber.Map{
        "user":  user,
        "token": token,
    })
}

// @Summary Login user
// @Description Login with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param request body validators.LoginRequest true "User login data"
// @Success 200 {object} map[string]interface{}
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
    var req validators.LoginRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
    }

    if err := validators.ValidateStruct(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": err.Error()})
    }

    user, token, err := h.authService.Login(&req)
    if err != nil {
        return c.Status(401).JSON(fiber.Map{"error": err.Error()})
    }

    return c.JSON(fiber.Map{
        "user":  user,
        "token": token,
    })
}

// @Summary Get current user
// @Description Get the currently authenticated user
// @Tags auth
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} models.User
// @Router /auth/me [get]
func (h *AuthHandler) GetCurrentUser(c *fiber.Ctx) error {
    userID := c.Locals("userID").(string)
    
    // Parse userID to UUID
    id, err := uuid.Parse(userID)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid user ID"})
    }

    // Get user from database (you'll need to implement this)
    // For now, return the userID
    return c.JSON(fiber.Map{
        "user_id": id,
        "message": "User authenticated successfully",
    })
}

// @Summary Logout user
// @Description Logout the current user
// @Tags auth
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *fiber.Ctx) error {
    // In a JWT implementation, logout is typically handled client-side
    // by removing the token from storage
    return c.JSON(fiber.Map{
        "message": "Logged out successfully",
    })
}