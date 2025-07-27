package handler

import (
	"fmt"

	"github.com/asfung/ticus/internal/app/adapter/inbound/api/mapper"
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

	if request.ID != "" {
		response, err := h.ArticleService.UpdateArticle(request.ID, *request)
		if err != nil {
			return echo.NewHTTPError(500, "Failed to update existing article")
		}
		return ctx.JSON(200, response)
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
	page := 1
	size := 10
	if p := ctx.QueryParam("page"); p != "" {
		fmt.Sscanf(p, "%d", &page)
	}
	if s := ctx.QueryParam("size"); s != "" {
		fmt.Sscanf(s, "%d", &size)
	}
	res, page, total, totalPage, err := h.ArticleService.ListArticles(page, size)
	if err != nil {
		return echo.NewHTTPError(500, "Failed to list articles")
	}
	return ctx.JSON(200, mapper.PageResponse[mapper.ArticleResponse]{
		Data: res,
		PageMetadata: mapper.PageMetadata{
			Page:      page,
			Size:      size,
			TotalItem: total,
			TotalPage: totalPage,
		},
	})
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

func (h *ArticleHandler) ToggleUpvote(ctx echo.Context) error {
	id := ctx.Param("id")
	res, err := h.ArticleService.ToggleUpvote(ctx, id)
	if err != nil {
		return echo.NewHTTPError(500, "Failed to toggle upvote")
	}
	// return ctx.JSON(200, res)
	return ctx.JSON(200, mapper.WebResponse[mapper.ArticleResponse]{
		Data:   *res,
		Errors: "",
	})
}

func (h *ArticleHandler) ToggleView(ctx echo.Context) error {
	id := ctx.Param("id")
	res, err := h.ArticleService.ToggleView(ctx, id)
	if err != nil {
		return echo.NewHTTPError(500, "Failed to toggle view")
	}
	// return ctx.JSON(200, res)
	return ctx.JSON(200, mapper.WebResponse[mapper.ArticleResponse]{
		Data:   *res,
		Errors: "",
	})
}
