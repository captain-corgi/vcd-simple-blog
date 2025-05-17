package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// BlogHandler handles blog-related requests
type BlogHandler struct {
	blogServiceURL string
}

// NewBlogHandler creates a new blog handler
func NewBlogHandler(blogServiceURL string) *BlogHandler {
	return &BlogHandler{
		blogServiceURL: blogServiceURL,
	}
}

// GetAllBlogs retrieves all blogs
func (h *BlogHandler) GetAllBlogs(c echo.Context) error {
	// Forward query parameters
	query := c.Request().URL.Query().Encode()
	url := h.blogServiceURL + "/blogs"
	if query != "" {
		url += "?" + query
	}

	resp, err := http.Get(url)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to blog service")
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode response")
	}

	return c.JSON(resp.StatusCode, responseBody)
}

// GetBlogByID retrieves a blog by ID
func (h *BlogHandler) GetBlogByID(c echo.Context) error {
	id := c.Param("id")
	resp, err := http.Get(h.blogServiceURL + "/blogs/" + id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to blog service")
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode response")
	}

	return c.JSON(resp.StatusCode, responseBody)
}

// CreateBlog creates a new blog
func (h *BlogHandler) CreateBlog(c echo.Context) error {
	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	// Add user ID from token
	userID := c.Get("userID").(string)
	requestBody["author_id"] = userID

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to marshal request")
	}

	req, err := http.NewRequest("POST", h.blogServiceURL+"/blogs", bytes.NewBuffer(jsonBody))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create request")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to blog service")
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode response")
	}

	return c.JSON(resp.StatusCode, responseBody)
}

// UpdateBlog updates a blog
func (h *BlogHandler) UpdateBlog(c echo.Context) error {
	id := c.Param("id")
	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	// Add user ID from token for authorization check in the blog service
	userID := c.Get("userID").(string)
	requestBody["user_id"] = userID

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to marshal request")
	}

	req, err := http.NewRequest("PUT", h.blogServiceURL+"/blogs/"+id, bytes.NewBuffer(jsonBody))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create request")
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to blog service")
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode response")
	}

	return c.JSON(resp.StatusCode, responseBody)
}

// DeleteBlog deletes a blog
func (h *BlogHandler) DeleteBlog(c echo.Context) error {
	id := c.Param("id")
	userID := c.Get("userID").(string)

	req, err := http.NewRequest("DELETE", h.blogServiceURL+"/blogs/"+id+"?user_id="+userID, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create request")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to blog service")
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode response")
	}

	return c.JSON(resp.StatusCode, responseBody)
}
