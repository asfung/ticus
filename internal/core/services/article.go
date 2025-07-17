package services

import (
	"time"

	"github.com/asfung/ticus/internal/app/adapter/handlers/api/mapper"
	"github.com/asfung/ticus/internal/app/adapter/handlers/api/mapper/converter"
	"github.com/asfung/ticus/internal/core/models"
	"github.com/asfung/ticus/internal/core/ports"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ArticleService struct {
	Repository *ports.ArticleRepository
	DB         *gorm.DB
	Log        *logrus.Logger
}

func NewArticleService(repository *ports.ArticleRepository, db *gorm.DB, log *logrus.Logger) *ArticleService {
	return &ArticleService{
		Repository: repository,
		DB:         db,
		Log:        log,
	}
}

func (s *ArticleService) CreateArticle(ctx echo.Context, req mapper.ArticleRequest) (*mapper.ArticleResponse, error) {
	article := &models.Article{
		Title:           req.Title,
		Slug:            req.Slug,
		ContentMarkdown: req.ContentMarkdown,
		ContentHTML:     req.ContentHTML,
		ContentJSON:     req.ContentJSON,
		IsDraft:         req.IsDraft,
		CategoryID:      req.CategoryID,
		UserID:          1, // TODO: inject real user
	}
	if !req.IsDraft {
		now := time.Now()
		article.PublishedAt = &now
	}

	if len(req.TagIDs) > 0 {
		var tags []models.Tag
		if err := s.DB.Where("id IN ?", req.TagIDs).Find(&tags).Error; err != nil {
			return nil, err
		}
		article.Tags = tags
	}

	if err := s.DB.Create(article).Error; err != nil {
		return nil, err
	}

	return converter.ArticleToResponse(article), nil
}

func (s *ArticleService) UpdateArticle(id string, req mapper.ArticleRequest) (*mapper.ArticleResponse, error) {
	article, err := s.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	article.Title = req.Title
	article.Slug = req.Slug
	article.ContentMarkdown = req.ContentMarkdown
	article.ContentHTML = req.ContentHTML
	article.ContentJSON = req.ContentJSON
	article.IsDraft = req.IsDraft
	article.CategoryID = req.CategoryID
	if !req.IsDraft && article.PublishedAt == nil {
		now := time.Now()
		article.PublishedAt = &now
	}

	if len(req.TagIDs) > 0 {
		var tags []models.Tag
		if err := s.DB.Where("id IN ?", req.TagIDs).Find(&tags).Error; err != nil {
			return nil, err
		}
		article.Tags = tags
	}

	if err := s.Repository.Update(article); err != nil {
		return nil, err
	}
	return converter.ArticleToResponse(article), nil
}

func (s *ArticleService) GetArticleByID(id string) (*mapper.ArticleResponse, error) {
	article, err := s.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return converter.ArticleToResponse(article), nil
}

func (s *ArticleService) DeleteArticle(id string) error {
	return s.Repository.Delete(id)
}

func (s *ArticleService) ListArticles() ([]mapper.ArticleResponse, error) {
	articles, err := s.Repository.FindAll()
	if err != nil {
		return nil, err
	}
	var res []mapper.ArticleResponse
	for _, a := range articles {
		res = append(res, *converter.ArticleToResponse(&a))
	}
	return res, nil
}
