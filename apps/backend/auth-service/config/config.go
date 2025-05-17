package config

import (
	"os"
	"strconv"
	"time"
)

// Config holds all configuration for the Auth Service
type Config struct {
	Environment string
	Database    DatabaseConfig
	JWT         JWTConfig
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
	SSLMode  string
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	Secret           string
	AccessTokenTTL   time.Duration
	RefreshTokenTTL  time.Duration
	Issuer           string
	AccessTokenCookie  string
	RefreshTokenCookie string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	// Database config
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "localhost"
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil || dbPort == 0 {
		dbPort = 5432
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "postgres"
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		dbPassword = "postgres"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "auth_db"
	}

	dbSSLMode := os.Getenv("DB_SSLMODE")
	if dbSSLMode == "" {
		dbSSLMode = "disable"
	}

	// JWT config
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "dev_secret_key"
	}

	accessTokenTTL, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_TTL"))
	if err != nil || accessTokenTTL == 0 {
		accessTokenTTL = 15 // 15 minutes
	}

	refreshTokenTTL, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TTL"))
	if err != nil || refreshTokenTTL == 0 {
		refreshTokenTTL = 7 * 24 * 60 // 7 days
	}

	jwtIssuer := os.Getenv("JWT_ISSUER")
	if jwtIssuer == "" {
		jwtIssuer = "vcd-simple-blog"
	}

	accessTokenCookie := os.Getenv("ACCESS_TOKEN_COOKIE")
	if accessTokenCookie == "" {
		accessTokenCookie = "access_token"
	}

	refreshTokenCookie := os.Getenv("REFRESH_TOKEN_COOKIE")
	if refreshTokenCookie == "" {
		refreshTokenCookie = "refresh_token"
	}

	return &Config{
		Environment: env,
		Database: DatabaseConfig{
			Host:     dbHost,
			Port:     dbPort,
			Username: dbUser,
			Password: dbPassword,
			DBName:   dbName,
			SSLMode:  dbSSLMode,
		},
		JWT: JWTConfig{
			Secret:             jwtSecret,
			AccessTokenTTL:     time.Duration(accessTokenTTL) * time.Minute,
			RefreshTokenTTL:    time.Duration(refreshTokenTTL) * time.Minute,
			Issuer:             jwtIssuer,
			AccessTokenCookie:  accessTokenCookie,
			RefreshTokenCookie: refreshTokenCookie,
		},
	}, nil
}
