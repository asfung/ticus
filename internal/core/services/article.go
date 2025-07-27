package services

import (
	"time"

	"github.com/asfung/ticus/internal/app/adapter/inbound/api/mapper"
	"github.com/asfung/ticus/internal/app/adapter/inbound/api/mapper/converter"
	"github.com/asfung/ticus/internal/core/models"
	"github.com/asfung/ticus/internal/core/ports"
	"github.com/asfung/ticus/internal/pkg/utils"
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
	user := ctx.Get("user").(*models.User)
	logrus.Info(user)

	slug := utils.GenerateSlug(req.Title)

	article := &models.Article{
		Title: req.Title,
		// Slug:            req.Slug,
		Slug:            slug,
		ContentMarkdown: req.ContentMarkdown,
		ContentHTML:     req.ContentHTML,
		ContentJSON:     req.ContentJSON,
		IsDraft:         req.IsDraft,
		CategoryID:      req.CategoryID,
		UserID:          user.ID,
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

	// return converter.ArticleToResponse(article), nil
	ar := converter.BuildArticleResponse(s.DB, article)
	return &ar, nil
}

func (s *ArticleService) UpdateArticle(id string, req mapper.ArticleRequest) (*mapper.ArticleResponse, error) {
	article, err := s.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	slug := utils.GenerateSlug(req.Title)

	article.Title = req.Title
	// article.Slug = req.Slug
	article.Slug = slug
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
	// return converter.ArticleToResponse(article), nil
	ar := converter.BuildArticleResponse(s.DB, article)
	return &ar, nil
}

func (s *ArticleService) GetArticleByID(id string) (*mapper.ArticleResponse, error) {
	article, err := s.Repository.FindByID(id)
	if err != nil {
		return nil, err
	}
	// return converter.ArticleToResponse(article), nil
	ar := converter.BuildArticleResponse(s.DB, article)
	return &ar, nil
}

func (s *ArticleService) DeleteArticle(id string) error {
	return s.Repository.Delete(id)
}

func (s *ArticleService) ListArticles(page, size int) ([]mapper.ArticleResponse, int, int64, int64, error) {
	if size <= 0 {
		size = 10
	}
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * size
	articles, total, err := s.Repository.FindAllPaginated(size, offset)
	if err != nil {
		return nil, page, 0, 0, err
	}
	var res []mapper.ArticleResponse
	for _, a := range articles {
		ar := converter.BuildArticleResponse(s.DB, &a)
		res = append(res, ar)
	}
	totalPage := (total + int64(size) - 1) / int64(size)
	return res, page, total, totalPage, nil
}

func (s *ArticleService) ToggleUpvote(ctx echo.Context, articleID string) (*mapper.ArticleResponse, error) {
	user := ctx.Get("user").(*models.User)

	article, err := s.Repository.FindByID(articleID)
	if err != nil {
		return nil, err
	}

	var upvote models.ArticleUpvote
	tx := s.DB.Where("article_id = ? AND user_id = ?", articleID, user.ID).First(&upvote)
	if tx.Error == nil {
		if err := s.DB.Delete(&upvote).Error; err != nil {
			return nil, err
		}
	} else {
		newUpvote := models.ArticleUpvote{
			ArticleID: articleID,
			UserID:    user.ID,
		}
		if err := s.DB.Create(&newUpvote).Error; err != nil {
			return nil, err
		}
	}

	// return converter.ArticleToResponse(article), nil
	ar := converter.BuildArticleResponse(s.DB, article)
	return &ar, nil
}

func (s *ArticleService) ToggleView(ctx echo.Context, articleID string) (*mapper.ArticleResponse, error) {
	user := ctx.Get("user").(*models.User)

	article, err := s.Repository.FindByID(articleID)
	if err != nil {
		return nil, err
	}

	var view models.ArticleView
	tx := s.DB.Where("article_id = ? AND user_id = ?", articleID, user.ID).First(&view)
	if tx.Error == nil {
		if err := s.DB.Model(&view).Update("updated_at", time.Now()).Error; err != nil {
			return nil, err
		}
	} else {
		newView := models.ArticleView{
			ArticleID: articleID,
			UserID:    user.ID,
		}
		if err := s.DB.Create(&newView).Error; err != nil {
			return nil, err
		}
	}

	// return converter.ArticleToResponse(article), nil
	// ensure the article's user is loaded before building the response
	// if s.DB.Model(article).Association("User").Error == nil && article.User.ID == "" {
	// 	s.DB.Model(article).Association("User").Find(&article.User)
	// }
	ar := converter.BuildArticleResponse(s.DB, article)
	return &ar, nil
}
