package handler

import (
	"github.com/labstack/echo/v4"
)

type CommonHandler struct {
	// service UserActions
}

func NewCommonHandler( /*service UserService*/ ) *CommonHandler {
	return &CommonHandler{ /* service: service*/ }
}

func (h *CommonHandler) SayHello(c echo.Context) error {

	name := c.Param("name")
	return c.JSON(200, map[string]string{
		"message": "Hello, " + name + "",
	})
	// return c.String(200, "Hello, "+name+"!")
}
