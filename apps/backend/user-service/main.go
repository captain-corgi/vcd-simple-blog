package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vcd-simple-blog/apps/backend/user-service/config"
	"github.com/vcd-simple-blog/apps/backend/user-service/infrastructure/database"
	"github.com/vcd-simple-blog/apps/backend/user-service/infrastructure/repository"
	"github.com/vcd-simple-blog/apps/backend/user-service/interfaces/http"
	"github.com/vcd-simple-blog/apps/backend/user-service/usecases"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db, err := database.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)

	// Initialize use cases
	userUseCase := usecases.NewUserUseCase(userRepo)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize API routes
	http.RegisterRoutes(e, userUseCase)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
