package senders

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSMTPSender tests the SMTPSender type.
func TestSMTPSender(t *testing.T) {
	sender := SMTPSender{}

	assert.Implements(t, (*Sender)(nil), sender)
}
