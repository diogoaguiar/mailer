package senders

import (
	"fmt"
	"net/smtp"

	"github.com/diogoaguiar/mailer/internal/mailer/messages"
)

// SMTPSender is a sender that sends emails using SMTP.
type SMTPSender struct {
	Host     string
	Port     int
	Username string
	Password string
	Message  *messages.Message
}

// NewSMTPSender creates a new SMTPSender.
func NewSMTPSender(host string, port int, username string, password string, message *messages.Message) *SMTPSender {
	return &SMTPSender{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Message:  message,
	}
}

// SendTo sends an email to the specified recipient.
func (s *SMTPSender) SendTo(recipient string) error {
	message := "To: " + recipient + "\r\n" +
		"Subject: " + s.Message.Subject + "\r\n" +
		"\r\n" +
		s.Message.Body + "\r\n"

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)

	auth := smtp.PlainAuth("", s.Username, s.Password, s.Host)
	err := smtp.SendMail(addr, auth, s.Username, []string{recipient}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
