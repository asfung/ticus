package mapper

type ArticleRequest struct {
	Title           string   `json:"title"`
	Slug            string   `json:"slug"`
	ContentMarkdown string   `json:"content_markdown"`
	ContentHTML     string   `json:"content_html"`
	ContentJSON     string   `json:"content_json"`
	IsDraft         bool     `json:"is_draft"`
	CategoryID      *uint64  `json:"category_id"`
	TagIDs          []string `json:"tag_ids"` // ids of selected tags
}

type ArticleResponse struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Slug            string   `json:"slug"`
	ContentMarkdown string   `json:"content_markdown"`
	ContentHTML     string   `json:"content_html"`
	ContentJSON     string   `json:"content_json"`
	IsDraft         bool     `json:"is_draft"`
	PublishedAt     *string  `json:"published_at,omitempty"`
	ViewCount       int      `json:"view_count"`
	LikeCount       int      `json:"like_count"`
	CategoryID      *uint64  `json:"category_id"`
	TagIDs          []string `json:"tag_ids"`
}
