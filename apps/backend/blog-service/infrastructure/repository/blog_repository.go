package repository

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"github.com/vcd-simple-blog/apps/backend/blog-service/domain/entity"
)

// BlogRepository implements the domain.repository.BlogRepository interface
type BlogRepository struct {
	db *gorm.DB
}

// NewBlogRepository creates a new blog repository
func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{
		db: db,
	}
}

// FindAll finds all blogs with pagination
func (r *BlogRepository) FindAll(ctx context.Context, limit, offset int) ([]*entity.Blog, error) {
	var blogs []*entity.Blog
	result := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}
	return blogs, nil
}

// FindByID finds a blog by ID
func (r *BlogRepository) FindByID(ctx context.Context, id string) (*entity.Blog, error) {
	var blog entity.Blog
	result := r.db.WithContext(ctx).First(&blog, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("blog not found")
		}
		return nil, result.Error
	}
	return &blog, nil
}

// FindByAuthorID finds blogs by author ID with pagination
func (r *BlogRepository) FindByAuthorID(ctx context.Context, authorID string, limit, offset int) ([]*entity.Blog, error) {
	var blogs []*entity.Blog
	result := r.db.WithContext(ctx).Where("author_id = ?", authorID).Limit(limit).Offset(offset).Find(&blogs)
	if result.Error != nil {
		return nil, result.Error
	}
	return blogs, nil
}

// Create creates a new blog
func (r *BlogRepository) Create(ctx context.Context, blog *entity.Blog) error {
	return r.db.WithContext(ctx).Create(blog).Error
}

// Update updates a blog
func (r *BlogRepository) Update(ctx context.Context, blog *entity.Blog) error {
	return r.db.WithContext(ctx).Save(blog).Error
}

// Delete deletes a blog
func (r *BlogRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&entity.Blog{}, "id = ?", id).Error
}
