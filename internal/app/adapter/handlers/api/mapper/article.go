package mapper

type ArticleRequest struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

type ArticleResponse struct {
	Author string `json:"author"`
}
