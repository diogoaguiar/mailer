package loggers

import (
	"testing"

	"github.com/diogoaguiar/mailer/internal/mailer"
	"github.com/stretchr/testify/assert"
)

// TestConsoleLogger tests the ConsoleLogger type.
func TestConsoleLogger(t *testing.T) {
	logger := ConsoleLogger{}

	assert.Implements(t, (*mailer.Logger)(nil), logger)
}

// TestNewConsoleLogger tests the NewConsoleLogger function.
func TestNewConsoleLogger(t *testing.T) {
	logger := NewConsoleLogger()

	assert.NotNil(t, logger)
	assert.IsType(t, &ConsoleLogger{}, logger)
}
