package http

import (
	"github.com/labstack/echo/v4"
	"github.com/vcd-simple-blog/apps/backend/auth-service/interfaces/http/handlers"
	"github.com/vcd-simple-blog/apps/backend/auth-service/usecases"
)

// RegisterRoutes registers all API routes
func RegisterRoutes(e *echo.Echo, authUseCase *usecases.AuthUseCase) {
	// Create handlers
	authHandler := handlers.NewAuthHandler(authUseCase)

	// API v1 group
	v1 := e.Group("/api/v1")

	// Auth routes
	auth := v1.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)
	auth.POST("/refresh", authHandler.RefreshToken)
	auth.POST("/logout", authHandler.Logout)
}
