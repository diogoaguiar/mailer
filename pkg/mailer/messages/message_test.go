package messages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMessage tests the Message type.
func TestMessage(t *testing.T) {
	message := Message{
		Subject: "Test",
		Body:    "This is a test.",
	}

	assert.IsType(t, Message{}, message)
}

// TestNewMessage tests the NewMessage function.
func TestNewMessage(t *testing.T) {
	message := NewMessage("Test", "This is a test.")

	assert.IsType(t, &Message{}, message)
	assert.Equal(t, "Test", message.Subject)
	assert.Equal(t, "This is a test.", message.Body)
}

// TestMessageFromFile tests the Message type.
func TestMessageFromFile(t *testing.T) {
	message := Message{}

	err := message.BodyFromFile("test/message.txt")
	assert.Nil(t, err)
	assert.Equal(t, "This is a simple test message.", message.Body)
}
