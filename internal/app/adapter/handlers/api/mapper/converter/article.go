package converter

import (
	"time"

	"github.com/asfung/ticus/internal/app/adapter/handlers/api/mapper"
	"github.com/asfung/ticus/internal/core/models"
)

// // func ArticleToResponse(article *models.Article) *mapper.ArticleResponse { // must be like this
// func ArticleToResponse(article *mapper.ArticleRequest) *mapper.ArticleResponse { // make sure it work
// 	return &mapper.ArticleResponse{
// 		Author: article.Author,
// 	}
// }

func ArticleToResponse(a *models.Article) *mapper.ArticleResponse {
	var publishedAt *string
	if a.PublishedAt != nil {
		t := a.PublishedAt.Format(time.RFC3339)
		publishedAt = &t
	}

	var tagIDs []string
	for _, tag := range a.Tags {
		tagIDs = append(tagIDs, tag.ID)
	}

	return &mapper.ArticleResponse{
		ID:              a.ID,
		Title:           a.Title,
		Slug:            a.Slug,
		ContentMarkdown: a.ContentMarkdown,
		ContentHTML:     a.ContentHTML,
		ContentJSON:     a.ContentJSON,
		IsDraft:         a.IsDraft,
		PublishedAt:     publishedAt,
		ViewCount:       a.ViewCount,
		LikeCount:       a.LikeCount,
		CategoryID:      a.CategoryID,
		TagIDs:          tagIDs,
	}
}
