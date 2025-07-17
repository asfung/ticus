package ports

import (
	"github.com/asfung/ticus/internal/app/adapter/handlers/api/mapper"
	"github.com/asfung/ticus/internal/core/models"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ArticleService interface {
	CreateArticle(ctx echo.Context, request mapper.ArticleRequest) (*mapper.ArticleResponse, error)
	UpdateArticle(id string, request mapper.ArticleRequest) (*mapper.ArticleResponse, error)
	GetArticleByID(id string) (*mapper.ArticleResponse, error)
	DeleteArticle(id string) error
	ListArticles() ([]mapper.ArticleResponse, error)
}

type ArticleRepository struct {
	DB *gorm.DB
	Repository[models.Article]
	Log *logrus.Logger
}

func NewArticleRepository(log *logrus.Logger, db *gorm.DB) *ArticleRepository {
	return &ArticleRepository{
		Log: log,
		DB:  db,
	}
}

// rest functions for model to databbase
func (r *ArticleRepository) FindByID(id string) (*models.Article, error) {
	var article models.Article
	err := r.DB.Preload("Tags").Preload("Category").First(&article, "id = ?", id).Error
	return &article, err
}

func (r *ArticleRepository) Update(article *models.Article) error {
	return r.DB.Session(&gorm.Session{FullSaveAssociations: true}).Updates(article).Error
}

func (r *ArticleRepository) Delete(id string) error {
	return r.DB.Delete(&models.Article{}, "id = ?", id).Error
}

func (r *ArticleRepository) FindAll() ([]models.Article, error) {
	var articles []models.Article
	err := r.DB.Preload("Tags").Preload("Category").Find(&articles).Error
	return articles, err
}
