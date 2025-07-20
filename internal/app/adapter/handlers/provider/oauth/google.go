package oauth

import (
	"fmt"
	"net/http"
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

	if clientID == "" || clientSecret == "" {
		panic("OAUTH_GOOGLE_CLIENT_ID and OAUTH_GOOGLE_CLIENT_SECRET environment variables are required")
	}

	provider := google.New(clientID, clientSecret, callbackURL, "email", "profile")
	goth.UseProviders(provider)

	return &GoogleOAuthService{
		provider: provider,
	}
}

func (s *GoogleOAuthService) GetAuthURL(w http.ResponseWriter, r *http.Request) (string, error) {
	state := r.URL.Query().Get("state")
	if state == "" {
		state = "state"
	}

	sess, err := s.provider.BeginAuth(state)
	if err != nil {
		return "", fmt.Errorf("failed to begin auth: %w", err)
	}

	url, err := sess.GetAuthURL()
	if err != nil {
		return "", fmt.Errorf("failed to get auth URL: %w", err)
	}

	return url, nil
}

func (s *GoogleOAuthService) CompleteUserAuth(w http.ResponseWriter, r *http.Request) (*goth.User, error) {
	state := r.URL.Query().Get("state")
	if state == "" {
		state = "state"
	}

	sess, err := s.provider.BeginAuth(state)
	if err != nil {
		return nil, fmt.Errorf("failed to begin auth: %w", err)
	}

	_, err = sess.Authorize(s.provider, r.URL.Query())
	if err != nil {
		return nil, fmt.Errorf("failed to authorize: %w", err)
	}

	user, err := s.provider.FetchUser(sess)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user: %w", err)
	}

	return &user, nil
}

func (s *GoogleOAuthService) GetProvider() goth.Provider {
	return s.provider
}
