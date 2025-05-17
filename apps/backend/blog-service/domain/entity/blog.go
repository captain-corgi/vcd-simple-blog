package entity

import (
	"errors"
	"time"

	"github.com/vcd-simple-blog/apps/backend/blog-service/domain/valueobject"
)

// Blog represents a blog post entity
type Blog struct {
	ID          string
	Title       string
	Content     string
	AuthorID    string
	Status      valueobject.BlogStatus
	Tags        []string
	PublishedAt *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewBlog creates a new blog entity
func NewBlog(id, title, content, authorID string, tags []string) (*Blog, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	if content == "" {
		return nil, errors.New("content cannot be empty")
	}

	if authorID == "" {
		return nil, errors.New("author ID cannot be empty")
	}

	now := time.Now()
	return &Blog{
		ID:        id,
		Title:     title,
		Content:   content,
		AuthorID:  authorID,
		Status:    valueobject.Draft,
		Tags:      tags,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

// Publish changes the blog status to published
func (b *Blog) Publish() error {
	if b.Status == valueobject.Published {
		return errors.New("blog is already published")
	}

	b.Status = valueobject.Published
	now := time.Now()
	b.PublishedAt = &now
	b.UpdatedAt = now
	return nil
}

// Update updates the blog content
func (b *Blog) Update(title, content string, tags []string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}

	if content == "" {
		return errors.New("content cannot be empty")
	}

	b.Title = title
	b.Content = content
	b.Tags = tags
	b.UpdatedAt = time.Now()
	return nil
}

// IsAuthor checks if the given user ID is the author of the blog
func (b *Blog) IsAuthor(userID string) bool {
	return b.AuthorID == userID
}
