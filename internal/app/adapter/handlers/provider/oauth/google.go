package oauth

import (
	"context"
	"net/url"
	"os"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "randomString"
	MaxAge = 86400 * 30
	IsProd = false
)

type GoogleOAuthService struct {
	provider goth.Provider
}

func NewGoogleOAuthService() *GoogleOAuthService {
	clientID := os.Getenv("OAUTH_GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_GOOGLE_CLIENT_SECRET")
	callbackURL := "http://localhost:8080/api/v1/auth/google/callback"

	provider := google.New(clientID, clientSecret, callbackURL, "email", "profile")
	goth.UseProviders(provider)

	return &GoogleOAuthService{
		provider: provider,
	}
}

func (s *GoogleOAuthService) GetAuthURL(state string) (string, error) {
	sess, err := s.provider.BeginAuth(state)
	if err != nil {
		return "", err
	}
	url, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s *GoogleOAuthService) CompleteUserAuth(ctx context.Context, state, code string) (*goth.User, error) {
	sess, err := s.provider.BeginAuth(state)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Set("state", state)
	params.Set("code", code)

	if _, err := sess.Authorize(s.provider, params); err != nil {
		return nil, err
	}

	user, err := s.provider.FetchUser(sess)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
