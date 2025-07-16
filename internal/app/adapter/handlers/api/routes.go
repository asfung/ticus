package api

import (
	"github.com/asfung/ticus/internal/app/adapter/handlers/api/handler"
	"github.com/labstack/echo/v4"
)

type Router struct {
	commonHandler  *handler.CommonHandler
	articleHandler *handler.ArticleHandler
}

func NewRouter(
	commonHandler *handler.CommonHandler,
	articleHandler *handler.ArticleHandler,
) *Router {
	return &Router{
		commonHandler:  commonHandler,
		articleHandler: articleHandler,
	}
}

func (r *Router) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")
	api.GET("", echo.HandlerFunc(func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "hello"})
	}))
	api.GET("/hello/:name", r.commonHandler.SayHello)

	// Urticle
	article := api.Group("/article")
	article.POST("", r.articleHandler.CreateArticle)

}

func RegisterRoutes(e *echo.Echo, r *Router) {
	r.RegisterRoutes(e)
}
