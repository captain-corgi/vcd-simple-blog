package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vcd-simple-blog/apps/backend/user-service/interfaces/http/dto"
	"github.com/vcd-simple-blog/apps/backend/user-service/usecases"
)

// UserHandler handles user-related requests
type UserHandler struct {
	userUseCase *usecases.UserUseCase
}

// NewUserHandler creates a new user handler
func NewUserHandler(userUseCase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// GetCurrentUser retrieves the current user's information
func (h *UserHandler) GetCurrentUser(c echo.Context) error {
	userID := c.Get("userID").(string)

	// Find user by auth service user ID
	user, err := h.userUseCase.GetUserByUserID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, dto.NewUserResponse(user))
}

// GetUserByID retrieves a user by ID
func (h *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")

	user, err := h.userUseCase.GetUserByID(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, dto.NewUserResponse(user))
}

// GetUserByUsername retrieves a user by username
func (h *UserHandler) GetUserByUsername(c echo.Context) error {
	username := c.Param("username")

	user, err := h.userUseCase.GetUserByUsername(c.Request().Context(), username)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, dto.NewUserResponse(user))
}

// GetAllUsers retrieves all users with pagination
func (h *UserHandler) GetAllUsers(c echo.Context) error {
	limitStr := c.QueryParam("limit")
	offsetStr := c.QueryParam("offset")

	limit := 10
	if limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	offset := 0
	if offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err == nil && parsedOffset >= 0 {
			offset = parsedOffset
		}
	}

	users, err := h.userUseCase.GetAllUsers(c.Request().Context(), limit, offset)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve users"})
	}

	// Convert to response DTOs
	var response []*dto.UserResponse
	for _, user := range users {
		response = append(response, dto.NewUserResponse(user))
	}

	return c.JSON(http.StatusOK, response)
}

// CreateUser creates a new user
func (h *UserHandler) CreateUser(c echo.Context) error {
	var req dto.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	user, err := h.userUseCase.CreateUser(c.Request().Context(), req.UserID, req.Username, req.Email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, dto.NewUserResponse(user))
}

// UpdateProfile updates the current user's profile
func (h *UserHandler) UpdateProfile(c echo.Context) error {
	userID := c.Get("userID").(string)

	// Find user by auth service user ID
	user, err := h.userUseCase.GetUserByUserID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	var req dto.UpdateProfileRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updatedUser, err := h.userUseCase.UpdateUserProfile(
		c.Request().Context(),
		user.ID,
		req.DisplayName,
		req.Bio,
		req.AvatarURL,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update profile"})
	}

	return c.JSON(http.StatusOK, dto.NewUserResponse(updatedUser))
}

// UpdateProfileStatus updates the current user's profile status
func (h *UserHandler) UpdateProfileStatus(c echo.Context) error {
	userID := c.Get("userID").(string)

	// Find user by auth service user ID
	user, err := h.userUseCase.GetUserByUserID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	var req dto.UpdateProfileStatusRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	status := dto.ToProfileStatus(req.Status)
	updatedUser, err := h.userUseCase.UpdateProfileStatus(c.Request().Context(), user.ID, status)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update profile status"})
	}

	return c.JSON(http.StatusOK, dto.NewUserResponse(updatedUser))
}
