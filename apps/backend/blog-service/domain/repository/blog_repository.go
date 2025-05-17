package repository

import (
	"context"

	"github.com/vcd-simple-blog/apps/backend/blog-service/domain/entity"
)

// BlogRepository defines the interface for blog data access
type BlogRepository interface {
	FindAll(ctx context.Context, limit, offset int) ([]*entity.Blog, error)
	FindByID(ctx context.Context, id string) (*entity.Blog, error)
	FindByAuthorID(ctx context.Context, authorID string, limit, offset int) ([]*entity.Blog, error)
	Create(ctx context.Context, blog *entity.Blog) error
	Update(ctx context.Context, blog *entity.Blog) error
	Delete(ctx context.Context, id string) error
}
