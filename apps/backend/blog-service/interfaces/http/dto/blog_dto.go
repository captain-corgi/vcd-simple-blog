package dto

import (
	"time"
	"github.com/vcd-simple-blog/apps/backend/blog-service/domain/valueobject"
)

// CreateBlogRequest represents the request for creating a blog
type CreateBlogRequest struct {
	Title    string   `json:"title" validate:"required"`
	Content  string   `json:"content" validate:"required"`
	Tags     []string `json:"tags"`
	AuthorID string   `json:"author_id" validate:"required"`
}

// UpdateBlogRequest represents the request for updating a blog
type UpdateBlogRequest struct {
	Title   string   `json:"title" validate:"required"`
	Content string   `json:"content" validate:"required"`
	Tags    []string `json:"tags"`
}

// BlogResponse represents the response with blog information
type BlogResponse struct {
	ID          string             `json:"id"`
	Title       string             `json:"title"`
	Content     string             `json:"content"`
	AuthorID    string             `json:"author_id"`
	Status      valueobject.BlogStatus `json:"status"`
	Tags        []string           `json:"tags"`
	PublishedAt *time.Time         `json:"published_at,omitempty"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

// BlogListResponse represents the response with a list of blogs
type BlogListResponse struct {
	Blogs []BlogResponse `json:"blogs"`
	Total int64          `json:"total"`
}
