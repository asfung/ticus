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
	if err := ctx.Bind(request); err != nil {
		return echo.NewHTTPError(400, "Invalid request body")
	}

	response, err := h.ArticleService.CreateArticle(ctx, *request)
	if err != nil {
		return echo.NewHTTPError(500, "Failed to create article")
	}

	return ctx.JSON(201, response)
}

func (h *ArticleHandler) GetArticle(ctx echo.Context) error {
	id := ctx.Param("id")
	res, err := h.ArticleService.GetArticleByID(id)
	if err != nil {
		return echo.NewHTTPError(404, "Article not found")
	}
	return ctx.JSON(200, res)
}

func (h *ArticleHandler) ListArticles(ctx echo.Context) error {
	res, err := h.ArticleService.ListArticles()
	if err != nil {
		return echo.NewHTTPError(500, "Failed to list articles")
	}
	return ctx.JSON(200, res)
}

func (h *ArticleHandler) UpdateArticle(ctx echo.Context) error {
	id := ctx.Param("id")
	req := new(mapper.ArticleRequest)
	if err := ctx.Bind(req); err != nil {
		return echo.NewHTTPError(400, "Invalid request body")
	}
	res, err := h.ArticleService.UpdateArticle(id, *req)
	if err != nil {
		return echo.NewHTTPError(500, "Failed to update article")
	}
	return ctx.JSON(200, res)
}

func (h *ArticleHandler) DeleteArticle(ctx echo.Context) error {
	id := ctx.Param("id")
	if err := h.ArticleService.DeleteArticle(id); err != nil {
		return echo.NewHTTPError(500, "Failed to delete article")
	}
	return ctx.NoContent(204)
}
