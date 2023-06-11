package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCycle_Empty(t *testing.T) {
	t.Parallel()

	ci := Cycle[int]([]int{})
	assert.False(t, ci.Next())
	assert.Equal(t, 0, ci.Value())
}

func TestCycle_NonEmpty(t *testing.T) {
	t.Parallel()

	ci := Cycle[int]([]int{1, 2, 3})
	assert.True(t, ci.Next())
	assert.Equal(t, 1, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 2, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 3, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 1, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 2, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 3, ci.Value())
}

func TestCycleLiteral(t *testing.T) {
	t.Parallel()

	ci := CycleLiteral[int](1, 2, 3)
	assert.True(t, ci.Next())
	assert.Equal(t, 1, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 2, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 3, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 1, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 2, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 3, ci.Value())
}
