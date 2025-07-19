package api

import (
	"net/http"

	"github.com/asfung/ticus/internal/app/adapter/handlers/api/handler"
	"github.com/labstack/echo/v4"
)

type Router struct {
	commonHandler  *handler.CommonHandler
	articleHandler *handler.ArticleHandler
	authHandler    *handler.AuthHandler
	oauthHandler   *handler.OAuthHandler
}

func NewRouter(
	commonHandler *handler.CommonHandler,
	articleHandler *handler.ArticleHandler,
	authHandler *handler.AuthHandler,
	oauthHandler *handler.OAuthHandler,
) *Router {
	return &Router{
		commonHandler:  commonHandler,
		articleHandler: articleHandler,
		authHandler:    authHandler,
		oauthHandler:   oauthHandler,
	}
}

func (r *Router) RegisterRoutes(e *echo.Echo) {
	api := e.Group("/api/v1")
	// api.GET("", echo.HandlerFunc(func(c echo.Context) error {
	// 	return c.JSON(200, map[string]string{"message": "hello"})
	// }))
	// oauth test
	e.GET("", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<p><a href='api/v1/auth/google/login'>google</a></p>")
	})
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

	auth.GET("/google/login", r.oauthHandler.GoogleLogin)
	auth.GET("/google/callback", r.oauthHandler.GoogleCallback)
	// auth.GET("/google/logout", r.oauthHandler.GoogleLogout)

	// auth.GET("/refresh", r.authHandler.Refresh, AuthMiddleware(r.authHandler.AuthService)) // alredy has header missing on authHandler
	auth.GET("/refresh", r.authHandler.Refresh, AuthMiddleware(r.authHandler.AuthService)).Name = "auth.refresh.token"
	auth.GET("/me", r.authHandler.Me, AuthMiddleware(r.authHandler.AuthService))
	auth.POST("/logout", r.authHandler.Logout, AuthMiddleware(r.authHandler.AuthService))

}

func RegisterRoutes(e *echo.Echo, r *Router) {
	r.RegisterRoutes(e)
}
