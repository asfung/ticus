package services

import (
	"errors"
	"time"

	"github.com/asfung/ticus/internal/core/models"
	"github.com/asfung/ticus/internal/core/ports"
	"github.com/asfung/ticus/internal/infrastructure/config"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	Repository *ports.AuthRepository
	DB         *gorm.DB
	JWTSecret  []byte
}

func NewAuthService(repository *ports.AuthRepository, db *gorm.DB, cfg *config.AppConfig) ports.AuthService {
	return &AuthService{
		DB:         db,
		JWTSecret:  []byte(cfg.JWTSecret),
		Repository: repository,
	}
}

type jwtCustomClaims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func (s *AuthService) Register(username, email, password string) (*models.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashed),
	}

	if err := s.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthService) Login(username, password string) (string, string, error) {
	var user models.User
	if err := s.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return "", "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", errors.New("invalid credentials")
	}

	access, err := s.generateToken(user, 5*time.Minute)
	if err != nil {
		return "", "", err
	}
	refresh, err := s.generateToken(user, 24*time.Hour)
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (s *AuthService) Refresh(refreshToken string) (string, error) {
	claims := &jwtCustomClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return s.JWTSecret, nil
		// })
	}, jwt.WithLeeway(0), jwt.WithoutClaimsValidation())

	// if err != nil {
	// 	return "", err
	// }

	if err != nil && !errors.Is(err, jwt.ErrTokenExpired) {
		return "", err // reject for reasons other than expiration
	}
	if !token.Valid && !errors.Is(err, jwt.ErrTokenExpired) {
		return "", errors.New("invalid token")
	}

	var user models.User
	if err := s.DB.First(&user, "id = ?", claims.UserID).Error; err != nil {
		return "", err
	}

	return s.generateToken(user, 5*time.Minute)
}

func (s *AuthService) Verify(tokenStr string) (*models.User, error) {
	claims := &jwtCustomClaims{}
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return s.JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := s.DB.First(&user, "id = ?", claims.UserID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *AuthService) generateToken(user models.User, duration time.Duration) (string, error) {
	claims := &jwtCustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		Email:    user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.JWTSecret)
}

func (s *AuthService) FindById(id string) (*models.User, error) {
	var user models.User
	if err := s.Repository.FindById(s.DB, &user, id); err != nil {
		return nil, err
	}
	return &user, nil
}
