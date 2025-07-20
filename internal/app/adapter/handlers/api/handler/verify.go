package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/asfung/ticus/internal/app/adapter/handlers/provider/mailer"
	"github.com/labstack/echo/v4"
)

type verificationCode struct {
	Code      string
	ExpiresAt time.Time
}

var (
	verificationCodes = make(map[string]verificationCode) // email -> code
	verificationMu    sync.Mutex
)

func generateCode() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}

type RequestVerificationInput struct {
	Email string `json:"email"`
}

type ConfirmVerificationInput struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

func RequestVerificationCode(m *mailer.Mailer) echo.HandlerFunc {

	return func(c echo.Context) error {
		var input RequestVerificationInput
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
		}
		code := generateCode()
		expiry := time.Now().Add(10 * time.Minute)
		verificationMu.Lock()
		verificationCodes[input.Email] = verificationCode{Code: code, ExpiresAt: expiry}
		verificationMu.Unlock()
		subject := "Your Verification Code"
		body := "Your verification code is: <b>" + code + "</b>"
		_ = m.SendMail(input.Email, subject, body) // ignore err
		return c.JSON(http.StatusOK, map[string]string{"message": "Verification code sent"})
	}
}

func ConfirmVerificationCode() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input ConfirmVerificationInput
		if err := c.Bind(&input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
		}
		verificationMu.Lock()
		vc, ok := verificationCodes[input.Email]
		verificationMu.Unlock()
		if !ok || vc.ExpiresAt.Before(time.Now()) || vc.Code != input.Code {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid or expired code"})
		}
		// mark user as verified (to be implemented in user db)

		// remove code after successful verification
		verificationMu.Lock()
		delete(verificationCodes, input.Email)
		verificationMu.Unlock()
		return c.JSON(http.StatusOK, map[string]string{"message": "Email verified"})
	}
}
