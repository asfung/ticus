package handler

import (
	"fmt"
	"net/http"

	"github.com/asfung/ticus/internal/app/adapter/handlers/provider/oauth"
	"github.com/labstack/echo/v4"
)

type OAuthHandler struct {
	GoogleOAuthService *oauth.GoogleOAuthService
}

func NewOAuthHandler(googleOAuthService *oauth.GoogleOAuthService) *OAuthHandler {
	return &OAuthHandler{
		GoogleOAuthService: googleOAuthService,
	}
}

func (h *OAuthHandler) GoogleLogin(ctx echo.Context) error {
	state := ctx.QueryParam("state")
	if state == "" {
		state = "state"
	}
	authURL, err := h.GoogleOAuthService.GetAuthURL(state)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to reach out the url: %v", err))
	}
	return ctx.Redirect(http.StatusTemporaryRedirect, authURL)
}

func (h *OAuthHandler) GoogleCallback(ctx echo.Context) error {
	state := ctx.QueryParam("state")
	code := ctx.QueryParam("code")
	if state == "" || code == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "missing state")
	}

	user, err := h.GoogleOAuthService.CompleteUserAuth(ctx.Request().Context(), state, code)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, fmt.Sprintf("authenticate with google failed: %v", err))
	}
	return ctx.JSON(http.StatusOK, user)
}
