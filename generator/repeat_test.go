package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	t.Parallel()

	s := Repeat(1)
	assert.True(t, s.Next())
	assert.Equal(t, 1, s.Value())
	assert.True(t, s.Next())
	assert.Equal(t, 1, s.Value())
}
