package api

import (
	"errors"
	"net/http"
	"strings"

	"github.com/asfung/ticus/internal/core/ports"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func BasicMiddleware(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	// e.Use(middleware.CORS())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderXRequestedWith, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

}

func AuthMiddleware(authService ports.AuthService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")

			user, err := authService.Verify(token)
			if err != nil {
				// return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
				if errors.Is(err, jwt.ErrTokenExpired) {
					return c.JSON(http.StatusUnauthorized, map[string]interface{}{
						"message": "token is expired",
						"key":     "refresh-token",
					})
				}
				if errors.Is(err, jwt.ErrTokenMalformed) {
					return echo.NewHTTPError(http.StatusUnauthorized, "token is malformed")
				}
				if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
					return echo.NewHTTPError(http.StatusUnauthorized, "invalid token signature")
				}
				if errors.Is(err, jwt.ErrTokenNotValidYet) {
					return echo.NewHTTPError(http.StatusUnauthorized, "token not valid yet")
				}
			}

			c.Set("user", user)
			return next(c)
		}
	}
}
