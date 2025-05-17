package entity

import (
	"errors"
	"time"

	"github.com/vcd-simple-blog/apps/backend/user-service/domain/valueobject"
)

// User represents a user entity
type User struct {
	ID            string
	UserID        string // Reference to auth service user ID
	Username      string
	Email         string
	DisplayName   string
	Bio           string
	AvatarURL     string
	ProfileStatus valueobject.ProfileStatus
	Role          valueobject.UserRole
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// NewUser creates a new user entity
func NewUser(id, userID, username, email string) (*User, error) {
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	now := time.Now()
	return &User{
		ID:            id,
		UserID:        userID,
		Username:      username,
		Email:         email,
		DisplayName:   username, // Default to username
		ProfileStatus: valueobject.ProfileStatusPublic,
		Role:          valueobject.RoleUser,
		CreatedAt:     now,
		UpdatedAt:     now,
	}, nil
}

// UpdateProfile updates the user's profile information
func (u *User) UpdateProfile(displayName, bio, avatarURL string) {
	if displayName != "" {
		u.DisplayName = displayName
	}

	u.Bio = bio
	u.AvatarURL = avatarURL
	u.UpdatedAt = time.Now()
}

// SetProfileStatus sets the user's profile status
func (u *User) SetProfileStatus(status valueobject.ProfileStatus) {
	u.ProfileStatus = status
	u.UpdatedAt = time.Now()
}

// SetRole sets the user's role
func (u *User) SetRole(role valueobject.UserRole) {
	u.Role = role
	u.UpdatedAt = time.Now()
}
