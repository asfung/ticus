package grpc

import (
	"context"
	"time"

	"github.com/asfung/ticus/internal/core/ports"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ArticleServer struct {
	UnimplementedArticleServiceServer
	articleService ports.ArticleService
	log            *logrus.Logger
}

func NewArticleServer(articleService ports.ArticleService, log *logrus.Logger) *ArticleServer {
	return &ArticleServer{
		articleService: articleService,
		log:            log,
	}
}

func (s *ArticleServer) GetAllArticle(ctx context.Context, req *GetAllArticleRequest) (*GetAllArticleResponse, error) {
	s.log.Infof("GetAllArticle called with page: %d, size: %d", req.GetPage(), req.GetSize())

	articles, currentPage, totalPages, totalCount, err := s.articleService.ListArticles(int(req.GetPage()), int(req.GetSize()))
	if err != nil {
		s.log.Errorf("Error getting articles: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to get articles: %v", err)
	}

	var grpcArticles []*ArticleResponse
	for _, article := range articles {
		grpcArticle := &ArticleResponse{
			Id:              article.ID,
			Title:           article.Title,
			Slug:            article.Slug,
			ContentMarkdown: article.ContentMarkdown,
			ContentHtml:     article.ContentHTML,
			ContentJson:     article.ContentJSON,
			IsDraft:         article.IsDraft,
			UpvoteCount:     int32(article.UpvoteCount),
			IsUpvoted:       article.IsUpvoted,
			ViewCount:       int32(article.ViewCount),
			IsViewed:        article.IsViewed,
			TagIds:          article.TagIDs,
		}

		if article.PublishedAt != nil {
			grpcArticle.PublishedAt = *article.PublishedAt
		}
		if article.LatestViewedAt != nil {
			grpcArticle.LatestViewedAt = article.LatestViewedAt.Format(time.RFC3339)
		}
		if article.CategoryID != nil {
			grpcArticle.CategoryId = *article.CategoryID
		}

		if article.User != nil {
			grpcArticle.User = &UserResponse{
				Id:        article.User.ID,
				Username:  article.User.Username,
				Email:     article.User.Email,
				AvatarUrl: article.User.AvatarURL,
			}
		}

		grpcArticles = append(grpcArticles, grpcArticle)
	}

	return &GetAllArticleResponse{
		Articles:    grpcArticles,
		CurrentPage: int32(currentPage),
		TotalPages:  int32(totalPages),
		TotalCount:  totalCount,
	}, nil
}

func (s *ArticleServer) GetArticleById(ctx context.Context, req *GetArticleByIdRequest) (*GetArticleByIdResponse, error) {
	s.log.Infof("GetArticleById called with id: %s", req.GetId())

	article, err := s.articleService.GetArticleByID(req.GetId())
	if err != nil {
		s.log.Errorf("Error getting article by ID %s: %v", req.GetId(), err)
		return nil, status.Errorf(codes.NotFound, "Article not found: %v", err)
	}

	grpcArticle := &ArticleResponse{
		Id:              article.ID,
		Title:           article.Title,
		Slug:            article.Slug,
		ContentMarkdown: article.ContentMarkdown,
		ContentHtml:     article.ContentHTML,
		ContentJson:     article.ContentJSON,
		IsDraft:         article.IsDraft,
		UpvoteCount:     int32(article.UpvoteCount),
		IsUpvoted:       article.IsUpvoted,
		ViewCount:       int32(article.ViewCount),
		IsViewed:        article.IsViewed,
		TagIds:          article.TagIDs,
	}

	if article.PublishedAt != nil {
		grpcArticle.PublishedAt = *article.PublishedAt
	}
	if article.LatestViewedAt != nil {
		grpcArticle.LatestViewedAt = article.LatestViewedAt.Format(time.RFC3339)
	}
	if article.CategoryID != nil {
		grpcArticle.CategoryId = *article.CategoryID
	}

	if article.User != nil {
		grpcArticle.User = &UserResponse{
			Id:        article.User.ID,
			Username:  article.User.Username,
			Email:     article.User.Email,
			AvatarUrl: article.User.AvatarURL,
		}
	}

	return &GetArticleByIdResponse{
		Article: grpcArticle,
	}, nil
}
