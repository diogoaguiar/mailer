package messages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMessage tests the Message type.
func TestSimpleMessage(t *testing.T) {
	message := SimpleMessage{}

	assert.Implements(t, (*Message)(nil), message)
}

// TestNewMessage tests the NewMessage function.
func TestNewMessage(t *testing.T) {
	message := NewMessage("Test", "This is a test.")

	assert.IsType(t, SimpleMessage{}, message)
	assert.Equal(t, "Test", message.Subject)
	assert.Equal(t, "This is a test.", message.Body)
}

// TestMessageFromFile tests the Message type.
func TestMessageFromFile(t *testing.T) {
	message := SimpleMessage{}

	err := message.BodyFromFile("test/message.txt")
	assert.Nil(t, err)
	assert.Equal(t, "This is a simple test message.", message.Body)
}
