package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	t.Parallel()

	ci := Count[int](0)
	assert.True(t, ci.Next())
	assert.Equal(t, 0, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 1, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 2, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 3, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 4, ci.Value())
}

func TestCountBy(t *testing.T) {
	t.Parallel()

	ci := CountBy[int](0, 2)
	assert.True(t, ci.Next())
	assert.Equal(t, 0, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 2, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 4, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 6, ci.Value())
	assert.True(t, ci.Next())
	assert.Equal(t, 8, ci.Value())
}

func TestCountBy_Chan(t *testing.T) {
	t.Parallel()

	ch := CountBy[int](0, 2).Chan()
	assert.Equal(t, 0, <-ch)
	assert.Equal(t, 2, <-ch)
	assert.Equal(t, 4, <-ch)
}
