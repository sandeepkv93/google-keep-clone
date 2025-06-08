package middleware

import (
    "strings"
    "github.com/gofiber/fiber/v2"
    "google-keep-clone/internal/services"
)

func AuthMiddleware(authService *services.AuthService) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Get Authorization header
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Status(401).JSON(fiber.Map{
                "error": "Authorization header required",
            })
        }

        // Check if it starts with "Bearer "
        if !strings.HasPrefix(authHeader, "Bearer ") {
            return c.Status(401).JSON(fiber.Map{
                "error": "Invalid authorization header format",
            })
        }

        // Extract token
        token := strings.TrimPrefix(authHeader, "Bearer ")
        if token == "" {
            return c.Status(401).JSON(fiber.Map{
                "error": "Token required",
            })
        }

        // Validate token
        claims, err := authService.ValidateToken(token)
        if err != nil {
            return c.Status(401).JSON(fiber.Map{
                "error": "Invalid or expired token",
            })
        }

        // Set user info in context
        c.Locals("userID", claims.UserID)
        c.Locals("email", claims.Email)

        return c.Next()
    }
}

func OptionalAuthMiddleware(authService *services.AuthService) fiber.Handler {
    return func(c *fiber.Ctx) error {
        // Get Authorization header
        authHeader := c.Get("Authorization")
        if authHeader == "" {
            return c.Next()
        }

        // Check if it starts with "Bearer "
        if !strings.HasPrefix(authHeader, "Bearer ") {
            return c.Next()
        }

        // Extract token
        token := strings.TrimPrefix(authHeader, "Bearer ")
        if token == "" {
            return c.Next()
        }

        // Validate token
        claims, err := authService.ValidateToken(token)
        if err != nil {
            return c.Next()
        }

        // Set user info in context if valid
        c.Locals("userID", claims.UserID)
        c.Locals("email", claims.Email)

        return c.Next()
    }
}