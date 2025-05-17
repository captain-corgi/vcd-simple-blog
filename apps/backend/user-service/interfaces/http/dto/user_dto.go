package dto

import (
	"time"

	"github.com/vcd-simple-blog/apps/backend/user-service/domain/entity"
	"github.com/vcd-simple-blog/apps/backend/user-service/domain/valueobject"
)

// CreateUserRequest represents a request to create a new user
type CreateUserRequest struct {
	UserID   string `json:"user_id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

// UpdateProfileRequest represents a request to update a user's profile
type UpdateProfileRequest struct {
	DisplayName string `json:"display_name"`
	Bio         string `json:"bio"`
	AvatarURL   string `json:"avatar_url"`
}

// UpdateProfileStatusRequest represents a request to update a user's profile status
type UpdateProfileStatusRequest struct {
	Status string `json:"status" validate:"required,oneof=public private limited"`
}

// UserResponse represents a user response
type UserResponse struct {
	ID            string    `json:"id"`
	UserID        string    `json:"user_id"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	DisplayName   string    `json:"display_name"`
	Bio           string    `json:"bio"`
	AvatarURL     string    `json:"avatar_url"`
	ProfileStatus string    `json:"profile_status"`
	Role          string    `json:"role"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// NewUserResponse creates a new user response from a user entity
func NewUserResponse(user *entity.User) *UserResponse {
	return &UserResponse{
		ID:            user.ID,
		UserID:        user.UserID,
		Username:      user.Username,
		Email:         user.Email,
		DisplayName:   user.DisplayName,
		Bio:           user.Bio,
		AvatarURL:     user.AvatarURL,
		ProfileStatus: string(user.ProfileStatus),
		Role:          string(user.Role),
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
}

// ToProfileStatus converts a string to a ProfileStatus value object
func ToProfileStatus(status string) valueobject.ProfileStatus {
	switch status {
	case "public":
		return valueobject.ProfileStatusPublic
	case "private":
		return valueobject.ProfileStatusPrivate
	case "limited":
		return valueobject.ProfileStatusLimited
	default:
		return valueobject.ProfileStatusPublic
	}
}
