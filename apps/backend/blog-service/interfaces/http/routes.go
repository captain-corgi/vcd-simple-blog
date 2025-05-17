package http

import (
	"github.com/labstack/echo/v4"
	"github.com/vcd-simple-blog/apps/backend/blog-service/interfaces/http/handlers"
	"github.com/vcd-simple-blog/apps/backend/blog-service/interfaces/http/middleware"
	"github.com/vcd-simple-blog/apps/backend/blog-service/usecases"
	"os"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(e *echo.Echo, blogUseCase *usecases.BlogUseCase) {
	// Create handlers
	blogHandler := handlers.NewBlogHandler(blogUseCase)

	// Create middleware
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev_secret_key"
	}
	authMiddleware := middleware.NewAuthMiddleware(jwtSecret)

	// API v1 group
	v1 := e.Group("/api/v1")

	// Blog routes
	blogs := v1.Group("/blogs")
	blogs.GET("", blogHandler.GetBlogs)
	blogs.GET("/:id", blogHandler.GetBlog)
	blogs.POST("", blogHandler.CreateBlog, authMiddleware.Authenticate)
	blogs.PUT("/:id", blogHandler.UpdateBlog, authMiddleware.Authenticate)
	blogs.POST("/:id/publish", blogHandler.PublishBlog, authMiddleware.Authenticate)
	blogs.DELETE("/:id", blogHandler.DeleteBlog, authMiddleware.Authenticate)
}
