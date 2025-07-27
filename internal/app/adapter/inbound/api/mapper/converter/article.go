package converter

import (
	"time"

	"github.com/asfung/ticus/internal/app/adapter/inbound/api/mapper"
	"github.com/asfung/ticus/internal/core/models"
	"gorm.io/gorm"
)

// // func ArticleToResponse(article *models.Article) *mapper.ArticleResponse { // must be like this
// func ArticleToResponse(article *mapper.ArticleRequest) *mapper.ArticleResponse { // make sure it work
// 	return &mapper.ArticleResponse{
// 		Author: article.Author,
// 	}
// }

// func ArticleToResponse(a *models.Article) *mapper.ArticleResponse {
// 	// var db *gorm.DB

// 	var publishedAt *string
// 	if a.PublishedAt != nil {
// 		t := a.PublishedAt.Format(time.RFC3339)
// 		publishedAt = &t
// 	}

// 	var tagIDs []string
// 	for _, tag := range a.Tags {
// 		tagIDs = append(tagIDs, tag.ID)
// 	}

// 	return &mapper.ArticleResponse{
// 		// User:            user,
// 		ID:              a.ID,
// 		Title:           a.Title,
// 		Slug:            a.Slug,
// 		ContentMarkdown: a.ContentMarkdown,
// 		ContentHTML:     a.ContentHTML,
// 		ContentJSON:     a.ContentJSON,
// 		IsDraft:         a.IsDraft,
// 		PublishedAt:     publishedAt,
// 		// ViewCount:       a.GetViewCount(db, userID),
// 		// UpvoteCount:     a.GetUpvoteCount(db),
// 		// IsUpvoted: func() bool {
// 		// 	upvoted, _ := a.HasBeenUpvotedByUser(db, a.UserID)
// 		// 	return upvoted
// 		// }(),
// 		// IsViewed: func() bool {
// 		// 	viewed, _ := a.HasBeenViewedByUser(db, a.UserID)
// 		// 	return viewed
// 		// }(),
// 		CategoryID: a.CategoryID,
// 		TagIDs:     tagIDs,
// 	}
// }

func BuildArticleResponse(db *gorm.DB, a *models.Article) mapper.ArticleResponse {
	// ensure the article's user is loaded before building the response
	if db.Model(a).Association("User").Error == nil && a.User.ID == "" {
		db.Model(a).Association("User").Find(&a.User)
	}
	return mapper.ArticleResponse{
		ID:              a.ID,
		User:            mapper.ToUserResponse(&a.User),
		Title:           a.Title,
		Slug:            a.Slug,
		ContentMarkdown: a.ContentMarkdown,
		ContentHTML:     a.ContentHTML,
		ContentJSON:     a.ContentJSON,
		IsDraft:         a.IsDraft,
		CategoryID:      a.CategoryID,
		TagIDs: func(tags []models.Tag) []string {
			var ids []string
			for _, tag := range tags {
				ids = append(ids, tag.ID)
			}
			return ids
		}(a.Tags),
		ViewCount: func() int {
			return a.GetViewCount(db, a.UserID)
		}(),
		UpvoteCount: func() int {
			return a.GetUpvoteCount(db)
		}(),
		IsUpvoted: func() bool {
			upvoted, _ := a.HasBeenUpvotedByUser(db, a.UserID)
			return upvoted
		}(),
		IsViewed: func() bool {
			viewed, _ := a.HasBeenViewedByUser(db, a.UserID)
			return viewed
		}(),
		LatestViewedAt: func() *time.Time {
			viewed, _ := a.HasBeenViewedByUser(db, a.UserID)
			if viewed {
				articleView := models.ArticleView{}
				db.Where("article_id = ? AND user_id = ?", a.ID, a.UserID).First(&articleView)
				return &articleView.UpdatedAt
			}
			return nil
		}(),
	}
}
