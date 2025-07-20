package mailer

import (
	"fmt"
	"net/smtp"
)

type Mailer struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

func NewMailer(host string, port int, username, password, from string) *Mailer {
	return &Mailer{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		From:     from,
	}
}

func (m *Mailer) SendMail(to, subject, body string) error {
	auth := smtp.PlainAuth("", m.Username, m.Password, m.Host)
	addr := fmt.Sprintf("%s:%d", m.Host, m.Port)
	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		"\r\n" + body)
	return smtp.SendMail(addr, auth, m.From, []string{to}, msg)
}
