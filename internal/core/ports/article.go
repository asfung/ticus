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
