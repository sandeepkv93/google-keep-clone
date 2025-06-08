package main

import (
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"google-keep-clone/internal/config"
	"google-keep-clone/internal/handlers"
	"google-keep-clone/internal/middleware"
	"google-keep-clone/internal/models"
	"google-keep-clone/internal/repositories"
	"google-keep-clone/internal/services"
	wsocket "google-keep-clone/internal/websocket"
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

	// Initialize WebSocket hub
	hub := wsocket.NewHub()
	go hub.Run()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	noteRepo := repositories.NewNoteRepository(db)
	labelRepo := repositories.NewLabelRepository(db)

	// Initialize services
	authService := services.NewAuthService(userRepo, cfg)
	noteService := services.NewNoteService(noteRepo, userRepo, hub)
	labelService := services.NewLabelService(labelRepo, noteRepo, userRepo, hub)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	noteHandler := handlers.NewNoteHandler(noteService)
	labelHandler := handlers.NewLabelHandler(labelService)

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
			"status":  "ok",
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

	// Notes routes (protected)
	notes := app.Group("/notes", middleware.AuthMiddleware(authService))

	// CRUD operations
	notes.Get("/", noteHandler.GetNotes)
	notes.Post("/", noteHandler.CreateNote)
	notes.Get("/:id", noteHandler.GetNoteByID)
	notes.Put("/:id", noteHandler.UpdateNote)
	notes.Delete("/:id", noteHandler.DeleteNote)

	// Note actions
	notes.Patch("/:id/pin", noteHandler.TogglePin)
	notes.Patch("/:id/archive", noteHandler.ToggleArchive)
	notes.Patch("/:id/color", noteHandler.UpdateColor)

	// Special views
	notes.Get("/search", noteHandler.SearchNotes)
	notes.Post("/search/advanced", noteHandler.SearchNotesAdvanced)
	notes.Get("/pinned", noteHandler.GetPinnedNotes)
	notes.Get("/archived", noteHandler.GetArchivedNotes)

	// Note label operations
	notes.Post("/:note_id/labels", labelHandler.AttachLabelToNote)
	notes.Delete("/:note_id/labels/:label_id", labelHandler.DetachLabelFromNote)

	// Labels routes (protected)
	labels := app.Group("/labels", middleware.AuthMiddleware(authService))

	// CRUD operations for labels
	labels.Get("/", labelHandler.GetLabels)
	labels.Post("/", labelHandler.CreateLabel)
	labels.Get("/:id", labelHandler.GetLabelByID)
	labels.Put("/:id", labelHandler.UpdateLabel)
	labels.Delete("/:id", labelHandler.DeleteLabel)

	// Label-specific views
	labels.Get("/:id/notes", labelHandler.GetNotesByLabel)

	// API routes (for future extensions)
	api := app.Group("/api", middleware.AuthMiddleware(authService))
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "API routes for future features",
			"user_id": c.Locals("userID"),
		})
	})

	// WebSocket routes
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		// Get user ID from query parameter (token validation should be done here)
		token := c.Query("token")
		if token == "" {
			_ = c.Close()
			return
		}

		// Validate token and get user ID
		claims, err := authService.ValidateToken(token)
		if err != nil {
			_ = c.Close()
			return
		}

		userID, err := uuid.Parse(claims.UserID)
		if err != nil {
			_ = c.Close()
			return
		}

		// Handle WebSocket connection
		hub.HandleWebSocket(c, userID)
	}))

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
