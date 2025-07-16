package services

import (
	"github.com/asfung/ticus/internal/app/adapter/handlers/api/mapper"
	"github.com/asfung/ticus/internal/app/adapter/handlers/api/mapper/converter"
	"github.com/asfung/ticus/internal/core/models"
	"github.com/asfung/ticus/internal/core/ports"
	"github.com/bwmarrin/snowflake"
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

func (s *ArticleService) CreateArticle( /*ctx context.Context,*/ ectx echo.Context, request mapper.ArticleRequest) (*mapper.ArticleResponse, error) {
	// transaction
	// tx := s.DB.WithContext(ctx).Begin()
	// defer tx.Rollback()
	var articleRequest mapper.ArticleRequest
	// err := ectx.Bind(&articleRequest)
	// if err != nil {
	// 	s.Log.Warnf("Invalid request body : %+v", err)
	// 	return nil, err
	// }

	// migrained by error handling

	node, err := snowflake.NewNode(1)
	if err != nil {
		s.Log.Errorf("Failed to create snowflake node: %v", err)
		return nil, err
	}
	id := node.Generate().Base58()
	s.Repository.Create(s.DB, &models.Article{
		ID:              id,
		UserID:          1,
		ContentMarkdown: request.Content,
	})
	logrus.Info("author request " + request.Author + "")

	return converter.ArticleToResponse(&articleRequest), nil

}
