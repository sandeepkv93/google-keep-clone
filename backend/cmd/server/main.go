package main

import (
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/cors"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "google-keep-clone/backend/internal/config"
    "google-keep-clone/backend/internal/handlers"
    "google-keep-clone/backend/internal/middleware"
    "google-keep-clone/backend/internal/models"
    "google-keep-clone/backend/internal/repositories"
    "google-keep-clone/backend/internal/services"
)

// @title Google Keep Clone API
// @version 1.0
// @description A Google Keep clone REST API
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Printf("No .env file found: %v", err)
    }

    // Load configuration
    cfg := config.Load()

    // Initialize database
    db, err := initDatabase(cfg.DatabaseURL)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto-migrate database schemas
    if err := db.AutoMigrate(
        &models.User{},
        &models.Note{},
        &models.Label{},
        &models.Attachment{},
    ); err != nil {
        log.Fatal("Failed to migrate database:", err)
    }

    // Initialize repositories
    userRepo := repositories.NewUserRepository(db)

    // Initialize services
    authService := services.NewAuthService(userRepo, cfg)

    // Initialize handlers
    authHandler := handlers.NewAuthHandler(authService)

    // Initialize Fiber app
    app := fiber.New(fiber.Config{
        ErrorHandler: func(c *fiber.Ctx, err error) error {
            code := fiber.StatusInternalServerError
            if e, ok := err.(*fiber.Error); ok {
                code = e.Code
            }
            return c.Status(code).JSON(fiber.Map{
                "error": err.Error(),
            })
        },
    })

    // Middleware
    app.Use(logger.New())
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
        AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
        AllowHeaders: "Origin,Content-Type,Accept,Authorization",
    }))

    // Health check
    app.Get("/health", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "status": "ok",
            "message": "Google Keep Clone API is running",
        })
    })

    // Auth routes
    authRoutes := app.Group("/auth")
    authRoutes.Post("/register", authHandler.Register)
    authRoutes.Post("/login", authHandler.Login)
    authRoutes.Post("/logout", authHandler.Logout)
    
    // Protected auth routes
    authRoutes.Get("/me", middleware.AuthMiddleware(authService), authHandler.GetCurrentUser)

    // API routes (will be protected)
    api := app.Group("/api", middleware.AuthMiddleware(authService))
    
    // Placeholder for future API routes
    api.Get("/", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "message": "API routes will be implemented here",
            "user_id": c.Locals("userID"),
        })
    })

    // Start server
    port := cfg.Port
    if port == "" {
        port = "8080"
    }

    log.Printf("🚀 Server starting on port %s", port)
    log.Printf("🏥 Health check: http://localhost:%s/health", port)
    log.Printf("📚 Environment: %s", cfg.Environment)

    if err := app.Listen(":" + port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}

func initDatabase(databaseURL string) (*gorm.DB, error) {
    if databaseURL == "" {
        log.Fatal("DATABASE_URL environment variable is required")
    }

    db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    log.Println("✅ Database connected successfully")
    return db, nil
}