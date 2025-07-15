package api

import (
	"github.com/asfung/ticus/internal/app/adapter/handlers/api/handler"
	"github.com/labstack/echo/v4"
)

type Router struct {
	commonHandler *handler.CommonHandler
}

func NewRouter(
	commonHandler *handler.CommonHandler,
) *Router {
	return &Router{
		commonHandler: commonHandler,
	}
}


func (r *Router) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")
	api.GET("", echo.HandlerFunc(func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "hello"})
	}))
}


func RegisterRoutes(e *echo.Echo, r *Router){
	r.RegisterRoutes(e)
}

