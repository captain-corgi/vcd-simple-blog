package usecases

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/vcd-simple-blog/blog-service/domain/entity"
	"github.com/vcd-simple-blog/blog-service/domain/repository"
)

// BlogUseCase implements the blog use cases
type BlogUseCase struct {
	blogRepo repository.BlogRepository
}

// NewBlogUseCase creates a new blog use case
func NewBlogUseCase(blogRepo repository.BlogRepository) *BlogUseCase {
	return &BlogUseCase{
		blogRepo: blogRepo,
	}
}

// GetAllBlogs retrieves all blogs with pagination
func (uc *BlogUseCase) GetAllBlogs(ctx context.Context, limit, offset int) ([]*entity.Blog, error) {
	return uc.blogRepo.FindAll(ctx, limit, offset)
}

// GetBlogByID retrieves a blog by ID
func (uc *BlogUseCase) GetBlogByID(ctx context.Context, id string) (*entity.Blog, error) {
	return uc.blogRepo.FindByID(ctx, id)
}

// GetBlogsByAuthor retrieves blogs by author ID
func (uc *BlogUseCase) GetBlogsByAuthor(ctx context.Context, authorID string, limit, offset int) ([]*entity.Blog, error) {
	return uc.blogRepo.FindByAuthorID(ctx, authorID, limit, offset)
}

// CreateBlog creates a new blog
func (uc *BlogUseCase) CreateBlog(ctx context.Context, title, content, authorID string, tags []string) (*entity.Blog, error) {
	id := uuid.New().String()
	blog, err := entity.NewBlog(id, title, content, authorID, tags)
	if err != nil {
		return nil, err
	}

	if err := uc.blogRepo.Create(ctx, blog); err != nil {
		return nil, err
	}

	return blog, nil
}

// UpdateBlog updates a blog
func (uc *BlogUseCase) UpdateBlog(ctx context.Context, id, title, content string, tags []string, userID string) (*entity.Blog, error) {
	blog, err := uc.blogRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if !blog.IsAuthor(userID) {
		return nil, errors.New("user is not the author of this blog")
	}

	if err := blog.Update(title, content, tags); err != nil {
		return nil, err
	}

	if err := uc.blogRepo.Update(ctx, blog); err != nil {
		return nil, err
	}

	return blog, nil
}

// PublishBlog publishes a blog
func (uc *BlogUseCase) PublishBlog(ctx context.Context, id, userID string) (*entity.Blog, error) {
	blog, err := uc.blogRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if !blog.IsAuthor(userID) {
		return nil, errors.New("user is not the author of this blog")
	}

	if err := blog.Publish(); err != nil {
		return nil, err
	}

	if err := uc.blogRepo.Update(ctx, blog); err != nil {
		return nil, err
	}

	return blog, nil
}

// DeleteBlog deletes a blog
func (uc *BlogUseCase) DeleteBlog(ctx context.Context, id, userID string) error {
	blog, err := uc.blogRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if !blog.IsAuthor(userID) {
		return errors.New("user is not the author of this blog")
	}

	return uc.blogRepo.Delete(ctx, id)
}
