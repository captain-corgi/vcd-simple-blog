package config

import (
	"os"
)

// Config holds all configuration for the API Gateway
type Config struct {
	Environment    string
	AuthServiceURL string
	BlogServiceURL string
	UserServiceURL string
	JWTSecret      string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	authServiceURL := os.Getenv("AUTH_SERVICE_URL")
	if authServiceURL == "" {
		authServiceURL = "http://localhost:8081"
	}

	blogServiceURL := os.Getenv("BLOG_SERVICE_URL")
	if blogServiceURL == "" {
		blogServiceURL = "http://localhost:8082"
	}

	userServiceURL := os.Getenv("USER_SERVICE_URL")
	if userServiceURL == "" {
		userServiceURL = "http://localhost:8083"
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev_secret_key"
	}

	return &Config{
		Environment:    env,
		AuthServiceURL: authServiceURL,
		BlogServiceURL: blogServiceURL,
		UserServiceURL: userServiceURL,
		JWTSecret:      jwtSecret,
	}, nil
}
