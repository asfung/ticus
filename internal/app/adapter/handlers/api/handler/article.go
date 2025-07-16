package handler

import (
	"github.com/asfung/ticus/internal/app/adapter/handlers/api/mapper"
	"github.com/asfung/ticus/internal/core/ports"
	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	ArticleService ports.ArticleService
}

func NewArticleHandler(articleService ports.ArticleService) *ArticleHandler {
	return &ArticleHandler{
		ArticleService: articleService,
	}
}

func (h *ArticleHandler) CreateArticle(ctx echo.Context) error {
	request := new(mapper.ArticleRequest)
	err := ctx.Bind(request)
	if err != nil {
		// logrus.Warnf("Invalid request body : %+v", err)
		return echo.NewHTTPError(400, "Invalid request")
	}
	// error handling so verboseeeee

	response, err := h.ArticleService.CreateArticle(ctx, *request)
	if err != nil {
		// logrus.Warnf("Error creating article: %+v", err)
		return echo.NewHTTPError(500, "Failed to create article")
	}

	return ctx.JSON(201, response)
}
