package usecases

import (
	"context"
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/vcd-simple-blog/apps/backend/auth-service/config"
	"github.com/vcd-simple-blog/apps/backend/auth-service/domain/entity"
	"github.com/vcd-simple-blog/apps/backend/auth-service/domain/repository"
	"github.com/vcd-simple-blog/apps/backend/auth-service/interfaces/http/dto"
)

// AuthUseCase implements authentication use cases
type AuthUseCase struct {
	userRepo  repository.UserRepository
	tokenRepo repository.TokenRepository
	jwtConfig config.JWTConfig
}

// NewAuthUseCase creates a new auth use case
func NewAuthUseCase(userRepo repository.UserRepository, tokenRepo repository.TokenRepository, jwtConfig config.JWTConfig) *AuthUseCase {
	return &AuthUseCase{
		userRepo:  userRepo,
		tokenRepo: tokenRepo,
		jwtConfig: jwtConfig,
	}
}

// Register registers a new user
func (uc *AuthUseCase) Register(ctx context.Context, email, username, password string) (*entity.User, error) {
	// Check if email already exists
	_, err := uc.userRepo.FindByEmail(ctx, email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	// Check if username already exists
	_, err = uc.userRepo.FindByUsername(ctx, username)
	if err == nil {
		return nil, errors.New("username already exists")
	}

	// Create new user
	id := uuid.New().String()
	user, err := entity.NewUser(id, email, username, password)
	if err != nil {
		return nil, err
	}

	// Save user to database
	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Login authenticates a user and returns tokens
func (uc *AuthUseCase) Login(ctx context.Context, email, password string) (*dto.TokenResponse, error) {
	// Find user by email
	user, err := uc.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Verify password
	if !user.VerifyPassword(password) {
		return nil, errors.New("invalid email or password")
	}

	// Generate tokens
	return uc.generateTokens(ctx, user)
}

// RefreshToken refreshes an access token using a refresh token
func (uc *AuthUseCase) RefreshToken(ctx context.Context, refreshToken string) (*dto.TokenResponse, error) {
	// Find token in database
	token, err := uc.tokenRepo.FindByToken(ctx, refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Check if token is expired
	if token.ExpiresAt.Before(time.Now()) {
		// Delete expired token
		_ = uc.tokenRepo.Delete(ctx, token.ID)
		return nil, errors.New("refresh token expired")
	}

	// Find user
	user, err := uc.userRepo.FindByID(ctx, token.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Delete old token
	if err := uc.tokenRepo.Delete(ctx, token.ID); err != nil {
		return nil, err
	}

	// Generate new tokens
	return uc.generateTokens(ctx, user)
}

// Logout invalidates a refresh token
func (uc *AuthUseCase) Logout(ctx context.Context, refreshToken string) error {
	// Find token in database
	token, err := uc.tokenRepo.FindByToken(ctx, refreshToken)
	if err != nil {
		return errors.New("invalid refresh token")
	}

	// Delete token
	return uc.tokenRepo.Delete(ctx, token.ID)
}

// generateTokens generates access and refresh tokens
func (uc *AuthUseCase) generateTokens(ctx context.Context, user *entity.User) (*dto.TokenResponse, error) {
	// Generate access token
	accessToken, err := uc.generateAccessToken(user)
	if err != nil {
		return nil, err
	}

	// Generate refresh token
	refreshToken := uuid.New().String()
	expiresAt := time.Now().Add(uc.jwtConfig.RefreshTokenTTL)

	// Create token entity
	token, err := entity.NewToken(uuid.New().String(), user.ID, refreshToken, expiresAt)
	if err != nil {
		return nil, err
	}

	// Save token to database
	if err := uc.tokenRepo.Create(ctx, token); err != nil {
		return nil, err
	}

	return &dto.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresIn:    int(uc.jwtConfig.AccessTokenTTL.Seconds()),
	}, nil
}

// generateAccessToken generates a JWT access token
func (uc *AuthUseCase) generateAccessToken(user *entity.User) (string, error) {
	// Create claims
	claims := jwt.MapClaims{
		"sub":  user.ID,
		"exp":  time.Now().Add(uc.jwtConfig.AccessTokenTTL).Unix(),
		"iat":  time.Now().Unix(),
		"iss":  uc.jwtConfig.Issuer,
		"role": user.Role,
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token
	return token.SignedString([]byte(uc.jwtConfig.Secret))
}
