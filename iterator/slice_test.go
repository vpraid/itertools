package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIterator_Empty(t *testing.T) {
	t.Parallel()

	s := Slice([]int{})
	assert.False(t, s.Next())
	assert.Equal(t, 0, s.Value())
}

func TestSliceIterator_NonEmpty(t *testing.T) {
	t.Parallel()

	s := Slice([]int{1, 2, 3})
	assert.True(t, s.Next())
	assert.Equal(t, 1, s.Value())
	assert.True(t, s.Next())
	assert.Equal(t, 2, s.Value())
	assert.True(t, s.Next())
	assert.Equal(t, 3, s.Value())
	assert.False(t, s.Next())
	assert.Equal(t, 0, s.Value())
}
