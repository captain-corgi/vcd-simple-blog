package http

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/vcd-simple-blog/apps/backend/user-service/interfaces/http/handlers"
	"github.com/vcd-simple-blog/apps/backend/user-service/interfaces/http/middleware"
	"github.com/vcd-simple-blog/apps/backend/user-service/usecases"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(e *echo.Echo, userUseCase *usecases.UserUseCase) {
	// Create handlers
	userHandler := handlers.NewUserHandler(userUseCase)

	// Create middleware
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev_secret_key"
	}
	authMiddleware := middleware.NewAuthMiddleware(jwtSecret)

	// API v1 group
	v1 := e.Group("/api/v1")

	// Public routes
	v1.GET("/users", userHandler.GetAllUsers)
	v1.GET("/users/:id", userHandler.GetUserByID)
	v1.GET("/users/username/:username", userHandler.GetUserByUsername)

	// Protected routes
	users := v1.Group("/users")
	users.Use(authMiddleware.Authenticate)
	users.GET("/me", userHandler.GetCurrentUser)
	users.PUT("/me/profile", userHandler.UpdateProfile)
	users.PUT("/me/profile-status", userHandler.UpdateProfileStatus)

	// Admin routes for user creation (typically handled by auth service)
	admin := v1.Group("/admin/users")
	admin.Use(authMiddleware.Authenticate)
	admin.POST("", userHandler.CreateUser)
}
