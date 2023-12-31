package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/source"
)

func TestMerge_Two(t *testing.T) {
	t.Parallel()

	it := Merge[int, int](func(a, b int) int { return a + b },
		source.Literal(1, 2),
		source.Literal(3, 4),
	)
	assert.True(t, it.Next())
	assert.Equal(t, 4, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 6, it.Value())
	assert.False(t, it.Next())
}

func TestMerge_UnequalLength(t *testing.T) {
	t.Parallel()

	it := Merge[int, int](func(a, b int) int { return a + b },
		source.Literal(1, 2),
		source.Literal(3),
	)
	assert.True(t, it.Next())
	assert.Equal(t, 4, it.Value())
	assert.False(t, it.Next())
}

func TestMerge_Collect(t *testing.T) {
	t.Parallel()

	it := Merge[int, int](func(a, b int) int { return a + b },
		source.Literal(1, 2),
		source.Literal(3, 4),
	)
	assert.Equal(t, []int{4, 6}, it.Collect())
}

func TestMerge_Chan(t *testing.T) {
	t.Parallel()

	it := Merge[int, int](func(a, b int) int { return a + b },
		source.Literal(1, 2),
		source.Literal(3, 4),
	)
	c := it.Chan()
	assert.Equal(t, 4, <-c)
	assert.Equal(t, 6, <-c)
}
