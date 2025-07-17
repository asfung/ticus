package utils

import (
	"strings"

	"github.com/labstack/echo/v4"
)

func GetTokenFromHeader(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", echo.NewHTTPError(401, "Missing or invalid token")
	}
	return strings.TrimPrefix(authHeader, "Bearer "), nil
}
