package messages

import (
	"os"
)

// Message is a message to be sent.
type Message interface {
	Subject() string
	Body() string
}

type SimpleMessage struct {
	subject string
	body    string
}

func (m *SimpleMessage) Subject() string {
	return m.subject
}

func (m *SimpleMessage) Body() string {
	return m.body
}

// NewMessage creates a new Message.
func NewMessage(subject string, body string) Message {
	return &SimpleMessage{
		subject: subject,
		body:    body,
	}
}

// BodyFromFile reads a message from a file.
func (m *SimpleMessage) BodyFromFile(path string) error {
	contents, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	m.body = string(contents)

	return nil
}
