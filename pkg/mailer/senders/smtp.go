package senders

import (
	"fmt"
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
func (s *SMTPSender) SendTo(subject string, body string, recipient string) error {
	message := "To: " + recipient + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n"

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	err := smtp.SendMail(addr, auth, s.Username, []string{recipient}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
