package iterator_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/iterator"
	"github.com/vpraid/itertools/source"
)

func TestChunks_Empty(t *testing.T) {
	t.Parallel()

	it := iterator.Chunks[int](nil, 2)
	assert.False(t, it.Next())
}

func TestChunks_Exact(t *testing.T) {
	t.Parallel()

	it := iterator.Chunks[int](source.Slice([]int{1, 2, 3, 4}), 2)
	assert.True(t, it.Next())
	assert.Equal(t, []int{1, 2}, it.Value().Collect())
	assert.True(t, it.Next())
	assert.Equal(t, []int{3, 4}, it.Value().Collect())
	assert.False(t, it.Next())
}

func TestChunks_Underflow(t *testing.T) {
	t.Parallel()

	it := iterator.Chunks[int](source.Slice([]int{1, 2, 3}), 2)
	assert.True(t, it.Next())
	assert.Equal(t, []int{1, 2}, it.Value().Collect())
	assert.True(t, it.Next())
	assert.Equal(t, []int{3}, it.Value().Collect())
	assert.False(t, it.Next())
}

func TestChunks_One(t *testing.T) {
	t.Parallel()

	it := iterator.Chunks[int](source.Slice([]int{1, 2}), 1)
	assert.True(t, it.Next())
	assert.Equal(t, []int{1}, it.Value().Collect())
	assert.True(t, it.Next())
	assert.Equal(t, []int{2}, it.Value().Collect())
	assert.False(t, it.Next())
}

func ExampleChunks() {
	// This example demonstrates how to use Chunks to iterate over a slice in
	// chunks of a given size.
	it := iterator.Chunks[int](source.Slice([]int{1, 2, 3, 4, 5}), 2)
	for it.Next() {
		chunk := it.Value()
		fmt.Println(chunk.Collect())
	}

	// Output:
	// [1 2]
	// [3 4]
	// [5]
}
