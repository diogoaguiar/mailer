package messages

import (
	"os"
)

// Message is a message to be sent.
type Message struct {
	Subject string
	Body    string
}

// NewMessage creates a new Message.
func NewMessage(subject string, body string) *Message {
	return &Message{
		Subject: subject,
		Body:    body,
	}
}

// BodyFromFile reads a message from a file.
func (m *Message) BodyFromFile(path string) error {
	contents, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	m.Body = string(contents)

	return nil
}
