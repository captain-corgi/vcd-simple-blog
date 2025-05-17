package entity

import (
	"errors"
	"time"
)

// Token represents a refresh token entity
type Token struct {
	ID        string
	UserID    string
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

// NewToken creates a new token entity
func NewToken(id, userID, token string, expiresAt time.Time) (*Token, error) {
	if userID == "" {
		return nil, errors.New("user ID cannot be empty")
	}

	if token == "" {
		return nil, errors.New("token cannot be empty")
	}

	if expiresAt.Before(time.Now()) {
		return nil, errors.New("expiration time must be in the future")
	}

	return &Token{
		ID:        id,
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}, nil
}

// IsExpired checks if the token is expired
func (t *Token) IsExpired() bool {
	return time.Now().After(t.ExpiresAt)
}
