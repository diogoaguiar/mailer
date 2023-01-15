package strategies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSequential tests the Sequential type.
func TestSequential(t *testing.T) {
	strategy := Sequential{}

	assert.Implements(t, (*Strategy)(nil), strategy)
}
