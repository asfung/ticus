package mapper

import (
	"time"
)

type ArticleRequest struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	// Slug            string   `json:"slug"`
	ContentMarkdown string   `json:"content_markdown"`
	ContentHTML     string   `json:"content_html"`
	ContentJSON     string   `json:"content_json"`
	IsDraft         bool     `json:"is_draft"`
	CategoryID      *uint64  `json:"category_id"`
	TagIDs          []string `json:"tag_ids"` // ids of selected tags
}

type ArticleResponse struct {
	User            *UserResponse `json:"user"`
	ID              string        `json:"id"`
	Title           string        `json:"title"`
	Slug            string        `json:"slug"`
	ContentMarkdown string        `json:"content_markdown"`
	ContentHTML     string        `json:"content_html"`
	ContentJSON     string        `json:"content_json"`
	IsDraft         bool          `json:"is_draft"`
	PublishedAt     *string       `json:"published_at,omitempty"`
	UpvoteCount     int           `json:"upvote_count"`
	IsUpvoted       bool          `json:"is_upvoted"`
	ViewCount       int           `json:"view_count"`
	IsViewed        bool          `json:"is_viewed"`
	LatestViewedAt  *time.Time    `json:"latest_viewed_at,omitempty"`
	CategoryID      *uint64       `json:"category_id"`
	TagIDs          []string      `json:"tag_ids"`
}
