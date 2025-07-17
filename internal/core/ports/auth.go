package ports

import (
	"github.com/asfung/ticus/internal/core/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(username, email, password string) (*models.User, error)
	Login(username, password string) (accessToken string, refreshToken string, err error)
	Refresh(refreshToken string) (accessToken string, err error)
	Verify(token string) (*models.User, error)
	FindById(id string) (*models.User, error)
}

type AuthRepository struct {
	DB *gorm.DB
	Repository[models.User]
	Log *logrus.Logger
}

func NewAuthRepository(db *gorm.DB, log *logrus.Logger) *AuthRepository {
	return &AuthRepository{
		DB:  db,
		Log: log,
	}
}
