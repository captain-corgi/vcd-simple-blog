package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// AuthHandler handles authentication requests
type AuthHandler struct {
	authServiceURL string
}

// NewAuthHandler creates a new auth handler
func NewAuthHandler(authServiceURL string) *AuthHandler {
	return &AuthHandler{
		authServiceURL: authServiceURL,
	}
}

// Login handles user login
func (h *AuthHandler) Login(c echo.Context) error {
	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to marshal request")
	}

	resp, err := http.Post(h.authServiceURL+"/login", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to auth service")
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode response")
	}

	return c.JSON(resp.StatusCode, responseBody)
}

// Register handles user registration
func (h *AuthHandler) Register(c echo.Context) error {
	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to marshal request")
	}

	resp, err := http.Post(h.authServiceURL+"/register", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to auth service")
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode response")
	}

	return c.JSON(resp.StatusCode, responseBody)
}

// RefreshToken handles token refresh
func (h *AuthHandler) RefreshToken(c echo.Context) error {
	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to marshal request")
	}

	resp, err := http.Post(h.authServiceURL+"/refresh", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to auth service")
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode response")
	}

	return c.JSON(resp.StatusCode, responseBody)
}

// Logout handles user logout
func (h *AuthHandler) Logout(c echo.Context) error {
	var requestBody map[string]interface{}
	if err := c.Bind(&requestBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to marshal request")
	}

	resp, err := http.Post(h.authServiceURL+"/logout", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to connect to auth service")
	}
	defer resp.Body.Close()

	var responseBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to decode response")
	}

	return c.JSON(resp.StatusCode, responseBody)
}
