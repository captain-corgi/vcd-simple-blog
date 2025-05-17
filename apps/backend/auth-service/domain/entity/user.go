package entity

import (
	"errors"
	"time"

	"github.com/vcd-simple-blog/auth-service/domain/valueobject"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user entity
type User struct {
	ID             string
	Email          string
	Username       string
	HashedPassword string
	Role           valueobject.UserRole
	Verified       bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// NewUser creates a new user entity
func NewUser(id, email, username, password string) (*User, error) {
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}

	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &User{
		ID:             id,
		Email:          email,
		Username:       username,
		HashedPassword: string(hashedPassword),
		Role:           valueobject.RoleUser,
		Verified:       false,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

// VerifyPassword checks if the provided password matches the stored hashed password
func (u *User) VerifyPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return err == nil
}

// ChangePassword changes the user's password
func (u *User) ChangePassword(newPassword string) error {
	if newPassword == "" {
		return errors.New("new password cannot be empty")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.HashedPassword = string(hashedPassword)
	u.UpdatedAt = time.Now()
	return nil
}

// VerifyEmail marks the user as verified
func (u *User) VerifyEmail() {
	u.Verified = true
	u.UpdatedAt = time.Now()
}

// PromoteToAdmin promotes the user to admin role
func (u *User) PromoteToAdmin() {
	u.Role = valueobject.RoleAdmin
	u.UpdatedAt = time.Now()
}
