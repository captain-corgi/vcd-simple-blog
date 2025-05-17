package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vcd-simple-blog/apps/backend/auth-service/config"
	"github.com/vcd-simple-blog/apps/backend/auth-service/infrastructure/database"
	"github.com/vcd-simple-blog/apps/backend/auth-service/infrastructure/repository"
	"github.com/vcd-simple-blog/apps/backend/auth-service/interfaces/http"
	"github.com/vcd-simple-blog/apps/backend/auth-service/usecases"
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
	tokenRepo := repository.NewTokenRepository(db)

	// Initialize use cases
	authUseCase := usecases.NewAuthUseCase(userRepo, tokenRepo, cfg.JWT)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize API routes
	http.RegisterRoutes(e, authUseCase)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
