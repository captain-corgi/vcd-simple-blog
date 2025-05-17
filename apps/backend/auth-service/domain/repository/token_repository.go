package repository

import (
	"context"

	"github.com/vcd-simple-blog/auth-service/domain/entity"
)

// TokenRepository defines the interface for token data access
type TokenRepository interface {
	FindByToken(ctx context.Context, token string) (*entity.Token, error)
	FindByUserID(ctx context.Context, userID string) ([]*entity.Token, error)
	Create(ctx context.Context, token *entity.Token) error
	Delete(ctx context.Context, id string) error
	DeleteByUserID(ctx context.Context, userID string) error
	DeleteExpired(ctx context.Context) error
}
