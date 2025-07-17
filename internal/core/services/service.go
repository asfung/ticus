// DEPRECATED
package services

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type ServiceCommon struct {
}

func NewServiceCommon() *ServiceCommon {
	return &ServiceCommon{}
}

func (s *ServiceCommon) GetTokenFromHeader(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid token")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	return token, nil
}
