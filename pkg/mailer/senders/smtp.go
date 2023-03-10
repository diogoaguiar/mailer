package senders

import (
	"fmt"
	"net/mail"
	"net/smtp"
)

// SMTPSender is a sender that sends emails using SMTP.
type SMTPSender struct {
	Host     string
	Port     int
	Username string
	Password string
}

// SendTo sends an email to the specified recipient.
func (s *SMTPSender) SendTo(subject string, body string, recipient *mail.Address) error {
	message := "To: " + recipient.String() + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body +
		"\r\n"

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	err := smtp.SendMail(addr, auth, s.Username, []string{recipient.Address}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
