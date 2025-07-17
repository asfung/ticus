package api

import (
	"github.com/asfung/ticus/internal/app/adapter/handlers/api/handler"
	"github.com/labstack/echo/v4"
)

type Router struct {
	commonHandler  *handler.CommonHandler
	articleHandler *handler.ArticleHandler
	authHandler    *handler.AuthHandler
}

func NewRouter(
	commonHandler *handler.CommonHandler,
	articleHandler *handler.ArticleHandler,
	authHandler *handler.AuthHandler,
) *Router {
	return &Router{
		commonHandler:  commonHandler,
		articleHandler: articleHandler,
		authHandler:    authHandler,
	}
}

func (r *Router) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")
	api.GET("", echo.HandlerFunc(func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "hello"})
	}))
	api.GET("/hello/:name", r.commonHandler.SayHello)

	// Article
	article := api.Group("/article")
	article.Use(AuthMiddleware(r.authHandler.AuthService))
	// article.POST("", r.articleHandler.CreateArticle)
	article.POST("", r.articleHandler.CreateArticle)
	article.GET("", r.articleHandler.ListArticles)
	article.GET("/:id", r.articleHandler.GetArticle)
	article.PUT("/:id", r.articleHandler.UpdateArticle)
	article.DELETE("/:id", r.articleHandler.DeleteArticle)

	// Auth
	auth := api.Group("/auth")
	// auth.Use(AuthMiddleware(r.authHandler.AuthService))
	auth.POST("/register", r.authHandler.Register)
	auth.POST("/login", r.authHandler.Login)

	// auth.GET("/refresh", r.authHandler.Refresh, AuthMiddleware(r.authHandler.AuthService)) // alredy has header missing on authHandler
	auth.GET("/refresh", r.authHandler.Refresh).Name = "auth-refresh-token"
	auth.GET("/me", r.authHandler.Me, AuthMiddleware(r.authHandler.AuthService))

}

func RegisterRoutes(e *echo.Echo, r *Router) {
	r.RegisterRoutes(e)
}
