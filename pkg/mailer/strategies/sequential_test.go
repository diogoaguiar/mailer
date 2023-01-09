package strategies

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSequential tests the Sequential type.
func TestSequential(t *testing.T) {
	strategy := Sequential{}

	assert.IsType(t, Sequential{}, strategy)
}

// TestNewSequential tests the NewSequential function.
func TestNewSequential(t *testing.T) {
	strategy := NewSequential(5, nil)

	assert.IsType(t, &Sequential{}, strategy)
	assert.Equal(t, 5, strategy.interval)
	assert.Nil(t, strategy.logger)
}
