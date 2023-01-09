package senders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSMTPSender tests the SMTPSender type.
func TestSMTPSender(t *testing.T) {
	sender := SMTPSender{}

	assert.IsType(t, SMTPSender{}, sender)
}

// TestNewSMTPSender tests the NewSMTPSender function.
func TestNewSMTPSender(t *testing.T) {
	sender := NewSMTPSender("localhost", 25, "user", "pass", nil)

	assert.IsType(t, &SMTPSender{}, sender)
	assert.Equal(t, "localhost", sender.Host)
	assert.Equal(t, 25, sender.Port)
	assert.Equal(t, "user", sender.Username)
	assert.Equal(t, "pass", sender.Password)
	assert.Nil(t, sender.Message)
}
