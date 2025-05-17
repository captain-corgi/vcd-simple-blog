package handlers

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
	"github.com/vcd-simple-blog/apps/backend/blog-service/interfaces/http/dto"
	"github.com/vcd-simple-blog/apps/backend/blog-service/usecases"
)

// BlogHandler handles blog-related HTTP requests
type BlogHandler struct {
	blogUseCase *usecases.BlogUseCase
}

// NewBlogHandler creates a new blog handler
func NewBlogHandler(blogUseCase *usecases.BlogUseCase) *BlogHandler {
	return &BlogHandler{
		blogUseCase: blogUseCase,
	}
}

// GetBlogs handles getting all blogs
func (h *BlogHandler) GetBlogs(c echo.Context) error {
	// Parse pagination parameters
	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil || offset < 0 {
		offset = 0
	}

	// Get blogs
	blogs, err := h.blogUseCase.GetAllBlogs(c.Request().Context(), limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// Convert to response
	response := make([]dto.BlogResponse, len(blogs))
	for i, blog := range blogs {
		response[i] = dto.BlogResponse{
			ID:          blog.ID,
			Title:       blog.Title,
			Content:     blog.Content,
			AuthorID:    blog.AuthorID,
			Status:      blog.Status,
			Tags:        blog.Tags,
			PublishedAt: blog.PublishedAt,
			CreatedAt:   blog.CreatedAt,
			UpdatedAt:   blog.UpdatedAt,
		}
	}

	return c.JSON(http.StatusOK, dto.BlogListResponse{
		Blogs: response,
		Total: int64(len(blogs)),
	})
}

// GetBlog handles getting a blog by ID
func (h *BlogHandler) GetBlog(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	blog, err := h.blogUseCase.GetBlogByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.BlogResponse{
		ID:          blog.ID,
		Title:       blog.Title,
		Content:     blog.Content,
		AuthorID:    blog.AuthorID,
		Status:      blog.Status,
		Tags:        blog.Tags,
		PublishedAt: blog.PublishedAt,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	})
}

// CreateBlog handles creating a new blog
func (h *BlogHandler) CreateBlog(c echo.Context) error {
	var req dto.CreateBlogRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	blog, err := h.blogUseCase.CreateBlog(c.Request().Context(), req.Title, req.Content, req.AuthorID, req.Tags)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.BlogResponse{
		ID:          blog.ID,
		Title:       blog.Title,
		Content:     blog.Content,
		AuthorID:    blog.AuthorID,
		Status:      blog.Status,
		Tags:        blog.Tags,
		PublishedAt: blog.PublishedAt,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	})
}

// UpdateBlog handles updating a blog
func (h *BlogHandler) UpdateBlog(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	var req dto.UpdateBlogRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Get user ID from token
	userID := c.Get("user_id").(string)

	blog, err := h.blogUseCase.UpdateBlog(c.Request().Context(), id, req.Title, req.Content, req.Tags, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.BlogResponse{
		ID:          blog.ID,
		Title:       blog.Title,
		Content:     blog.Content,
		AuthorID:    blog.AuthorID,
		Status:      blog.Status,
		Tags:        blog.Tags,
		PublishedAt: blog.PublishedAt,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	})
}

// PublishBlog handles publishing a blog
func (h *BlogHandler) PublishBlog(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	// Get user ID from token
	userID := c.Get("user_id").(string)

	blog, err := h.blogUseCase.PublishBlog(c.Request().Context(), id, userID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, dto.BlogResponse{
		ID:          blog.ID,
		Title:       blog.Title,
		Content:     blog.Content,
		AuthorID:    blog.AuthorID,
		Status:      blog.Status,
		Tags:        blog.Tags,
		PublishedAt: blog.PublishedAt,
		CreatedAt:   blog.CreatedAt,
		UpdatedAt:   blog.UpdatedAt,
	})
}

// DeleteBlog handles deleting a blog
func (h *BlogHandler) DeleteBlog(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID is required"})
	}

	// Get user ID from token
	userID := c.Get("user_id").(string)

	if err := h.blogUseCase.DeleteBlog(c.Request().Context(), id, userID); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Blog deleted successfully"})
}
