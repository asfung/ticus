package ports

import (
	"github.com/asfung/ticus/internal/core/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AuthService interface {
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
