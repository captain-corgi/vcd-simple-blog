package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"github.com/vcd-simple-blog/apps/backend/auth-service/domain/entity"
	"time"
)

// TokenRepository implements the domain.repository.TokenRepository interface
type TokenRepository struct {
	db *gorm.DB
}

// NewTokenRepository creates a new token repository
func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{
		db: db,
	}
}

// FindByToken finds a token by token string
func (r *TokenRepository) FindByToken(ctx context.Context, token string) (*entity.Token, error) {
	var t entity.Token
	result := r.db.WithContext(ctx).First(&t, "token = ?", token)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("token not found")
		}
		return nil, result.Error
	}
	return &t, nil
}

// FindByUserID finds tokens by user ID
func (r *TokenRepository) FindByUserID(ctx context.Context, userID string) ([]*entity.Token, error) {
	var tokens []*entity.Token
	result := r.db.WithContext(ctx).Find(&tokens, "user_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return tokens, nil
}

// Create creates a new token
func (r *TokenRepository) Create(ctx context.Context, token *entity.Token) error {
	return r.db.WithContext(ctx).Create(token).Error
}

// Delete deletes a token
func (r *TokenRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&entity.Token{}, "id = ?", id).Error
}

// DeleteByUserID deletes all tokens for a user
func (r *TokenRepository) DeleteByUserID(ctx context.Context, userID string) error {
	return r.db.WithContext(ctx).Delete(&entity.Token{}, "user_id = ?", userID).Error
}

// DeleteExpired deletes all expired tokens
func (r *TokenRepository) DeleteExpired(ctx context.Context) error {
	return r.db.WithContext(ctx).Delete(&entity.Token{}, "expires_at < ?", time.Now()).Error
}
