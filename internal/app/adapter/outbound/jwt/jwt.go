package jwt

import (
	"fmt"
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secret []byte
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwtlib.RegisteredClaims
}

func NewJWTService(secret string) *JWTService {
	return &JWTService{secret: []byte(secret)}
}

func (s *JWTService) GenerateAccessToken(username string, admin bool) (string, error) {
	return s.generateToken(username, admin, 5*time.Minute)
}

func (s *JWTService) GenerateRefreshToken(username string, admin bool) (string, error) {
	return s.generateToken(username, admin, 24*time.Hour)
}

func (s *JWTService) RefreshAccessToken(refreshToken string) (string, error) {
	token, err := jwtlib.ParseWithClaims(refreshToken, &jwtCustomClaims{}, func(t *jwtlib.Token) (interface{}, error) {
		return s.secret, nil
	})
	if err != nil || !token.Valid {
		return "", err
	}

	claims, ok := token.Claims.(*jwtCustomClaims)
	if !ok || claims.ExpiresAt.Time.Before(time.Now()) {
		return "", fmt.Errorf("invalid refresh token")
	}

	return s.GenerateAccessToken(claims.Name, claims.Admin)
}

func (s *JWTService) generateToken(username string, admin bool, duration time.Duration) (string, error) {
	claims := &jwtCustomClaims{
		Name:  username,
		Admin: admin,
		RegisteredClaims: jwtlib.RegisteredClaims{
			ExpiresAt: jwtlib.NewNumericDate(time.Now().Add(duration)),
		},
	}
	token := jwtlib.NewWithClaims(jwtlib.SigningMethodHS256, claims)
	return token.SignedString(s.secret)
}
