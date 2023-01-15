package loggers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestConsoleLogger tests the ConsoleLogger type.
func TestConsoleLogger(t *testing.T) {
	logger := ConsoleLogger{}

	assert.Implements(t, (*Logger)(nil), logger)
}
