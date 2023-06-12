package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/source"
)

func TestWindows_Empty(t *testing.T) {
	t.Parallel()

	it := Windows[int](source.Slice([]int{}), 5)
	assert.False(t, it.Next())
}

func TestWindows_NotEnoughElements(t *testing.T) {
	t.Parallel()

	it := Windows[int](source.Slice([]int{1, 2, 3}), 5)
	assert.False(t, it.Next())
}

func TestWindows(t *testing.T) {
	t.Parallel()

	it := Windows[int](source.Slice([]int{1, 2, 3, 4, 5}), 3)
	assert.True(t, it.Next())
	assert.Equal(t, []int{1, 2, 3}, it.Value().Collect())
	assert.True(t, it.Next())
	assert.Equal(t, []int{2, 3, 4}, it.Value().Collect())
	assert.True(t, it.Next())
	assert.Equal(t, []int{3, 4, 5}, it.Value().Collect())
	assert.False(t, it.Next())
}

func TestWindows_Chan(t *testing.T) {
	t.Parallel()

	it := Windows[int](source.Slice([]int{1, 2, 3, 4}), 2)
	assert.True(t, it.Next())

	c1 := it.Value().Chan()
	assert.True(t, it.Next())
	c2 := it.Value().Chan()
	assert.True(t, it.Next())
	c3 := it.Value().Chan()
	assert.False(t, it.Next())

	assert.Equal(t, 1, <-c1)
	assert.Equal(t, 2, <-c1)
	assert.Equal(t, 2, <-c2)
	assert.Equal(t, 3, <-c2)
	assert.Equal(t, 3, <-c3)
	assert.Equal(t, 4, <-c3)
}
