package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// UserHandler handles user-related requests
type UserHandler struct {
	userServiceURL string
}

// NewUserHandler creates a new user handler
func NewUserHandler(userServiceURL string) *UserHandler {
	return &UserHandler{
		userServiceURL: userServiceURL,
	}
}

// GetCurrentUser retrieves the current user's information
func (h *UserHandler) GetCurrentUser(c echo.Context) error {
	userID := c.Get("userID").(string)
	resp, err := http.Get(h.userServiceURL + "/users/" + userID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to user service")
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode response")
	}

	return c.JSON(resp.StatusCode, responseBody)
}

// UpdateCurrentUser updates the current user's information
func (h *UserHandler) UpdateCurrentUser(c echo.Context) error {
	userID := c.Get("userID").(string)
	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to marshal request")
	}

	req, err := http.NewRequest("PUT", h.userServiceURL+"/users/"+userID, bytes.NewBuffer(jsonBody))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create request")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to user service")
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode response")
	}

	return c.JSON(resp.StatusCode, responseBody)
}

// GetUserByID retrieves a user by ID
func (h *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("id")
	resp, err := http.Get(h.userServiceURL + "/users/" + id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to user service")
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode response")
	}

	return c.JSON(resp.StatusCode, responseBody)
}
