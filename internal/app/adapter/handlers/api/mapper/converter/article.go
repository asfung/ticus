package converter

import (
	"github.com/asfung/ticus/internal/app/adapter/handlers/api/mapper"
)

// func ArticleToResponse(article *models.Article) *mapper.ArticleResponse { // must be like this
func ArticleToResponse(article *mapper.ArticleRequest) *mapper.ArticleResponse { // make sure it work
	return &mapper.ArticleResponse{
		Author: article.Author,
	}
}
