package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/source"
)

func TestChain_Two(t *testing.T) {
	t.Parallel()

	it := Chain[int](source.Literal(1), source.Literal(2))
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Value())
	assert.False(t, it.Next())
}

func TestChain_Several(t *testing.T) {
	t.Parallel()

	it := Chain[int](
		source.Literal(1, 2, 3),
		source.Literal(4, 5, 6),
		source.Literal(7, 8, 9),
	)
	for i := 1; i < 10; i++ {
		assert.True(t, it.Next())
		assert.Equal(t, i, it.Value())
	}
	assert.False(t, it.Next())
}

func TestChain_Collect(t *testing.T) {
	t.Parallel()

	it := Chain[int](
		source.Literal(1, 2, 3),
		source.Literal(4, 5, 6),
		source.Literal(7, 8, 9),
	)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, it.Collect())
}

func TestChain_Chan(t *testing.T) {
	t.Parallel()

	it := Chain[int](
		source.Literal(1, 2, 3),
		source.Literal(4, 5, 6),
		source.Literal(7, 8, 9),
	)

	var items []int
	for item := range it.Chan() {
		items = append(items, item)
	}
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, items)
}
