package http

import (
	"github.com/labstack/echo/v4"
	"github.com/vcd-simple-blog/apps/backend/api-gateway/config"
	"github.com/vcd-simple-blog/apps/backend/api-gateway/interfaces/http/handlers"
	"github.com/vcd-simple-blog/apps/backend/api-gateway/interfaces/http/middleware"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(e *echo.Echo, cfg *config.Config) {
	// Create middleware
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTSecret)

	// Create handlers
	authHandler := handlers.NewAuthHandler(cfg.AuthServiceURL)
	blogHandler := handlers.NewBlogHandler(cfg.BlogServiceURL)
	userHandler := handlers.NewUserHandler(cfg.UserServiceURL)

	// API v1 group
	v1 := e.Group("/api/v1")

	// Auth routes
	auth := v1.Group("/auth")
	auth.POST("/login", authHandler.Login)
	auth.POST("/register", authHandler.Register)
	auth.POST("/refresh", authHandler.RefreshToken)
	auth.POST("/logout", authHandler.Logout)

	// Blog routes
	blog := v1.Group("/blogs")
	blog.GET("", blogHandler.GetAllBlogs)
	blog.GET("/:id", blogHandler.GetBlogByID)
	blog.POST("", blogHandler.CreateBlog, authMiddleware.Authenticate)
	blog.PUT("/:id", blogHandler.UpdateBlog, authMiddleware.Authenticate)
	blog.DELETE("/:id", blogHandler.DeleteBlog, authMiddleware.Authenticate)

	// User routes
	user := v1.Group("/users", authMiddleware.Authenticate)
	user.GET("/me", userHandler.GetCurrentUser)
	user.PUT("/me", userHandler.UpdateCurrentUser)
	user.GET("/:id", userHandler.GetUserByID)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
}
