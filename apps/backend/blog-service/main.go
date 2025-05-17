package main

import (
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vcd-simple-blog/blog-service/config"
	"github.com/vcd-simple-blog/blog-service/infrastructure/database"
	"github.com/vcd-simple-blog/blog-service/infrastructure/repository"
	"github.com/vcd-simple-blog/blog-service/interfaces/http"
	"github.com/vcd-simple-blog/blog-service/usecases"
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
	blogRepo := repository.NewBlogRepository(db)

	// Initialize use cases
	blogUseCase := usecases.NewBlogUseCase(blogRepo)

	// Create Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize API routes
	http.RegisterRoutes(e, blogUseCase)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
