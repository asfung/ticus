package handler

import (
	"net/http"

	"github.com/asfung/ticus/internal/app/adapter/handlers/api/mapper"
	"github.com/asfung/ticus/internal/core/ports"
	"github.com/asfung/ticus/internal/pkg/utils"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthService ports.AuthService // interface
	// Service     services.AuthService // i floating in the ocean and idk where the vuk island at
}

func NewAuthHandler(authService ports.AuthService /*, service services.AuthService*/) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
		// Service:     service,
	}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user, err := h.AuthService.Register(req.Username, req.Email, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	access, refresh, err := h.AuthService.Login(req.Username, req.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func (h *AuthHandler) Refresh(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing Authorization header")
	}

	const bearerPrefix = "Bearer "
	if len(authHeader) > len(bearerPrefix) && authHeader[:len(bearerPrefix)] == bearerPrefix {
		authHeader = authHeader[len(bearerPrefix):]
	}

	access, err := h.AuthService.Refresh(authHeader)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"access_token": access,
	})
}

func (h *AuthHandler) Me(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing Authorization header")
	}

	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return echo.NewHTTPError(401, err.Error())
	}
	user, err := h.AuthService.Verify(token)
	if err != nil {
		return echo.NewHTTPError(401, err.Error())
	}
	// return c.JSON(200, mapper.UserResponse{
	// 	ID:        user.ID,
	// 	Username:  user.Username,
	// 	Email:     user.Email,
	// 	AvatarURL: user.AvatarURL,
	// 	Bio:       user.Bio,
	// 	CreatedAt: user.CreatedAt,
	// 	UpdatedAt: user.UpdatedAt,
	// })
	return c.JSON(200, mapper.ToUserResponse(user))

}

func (h *AuthHandler) Logout(c echo.Context) error {
	token, err := utils.GetTokenFromHeader(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.AuthService.Logout(token)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "you logged out pal",
	})
}
