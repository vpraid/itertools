package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupBy_Empty(t *testing.T) {
	t.Parallel()

	it := GroupBy[int, bool](
		Slice([]int{}),
		func(i int) bool { return i%2 == 0 },
	)

	assert.False(t, it.Next())
}

func TestGroupBy_OneGroup(t *testing.T) {
	t.Parallel()

	it := GroupBy[int, bool](
		Slice([]int{1, 3, 5}),
		func(i int) bool { return i%2 == 0 },
	)

	assert.True(t, it.Next())
	group := it.Value()
	assert.False(t, group.Key)
	assert.True(t, group.Next())
	assert.Equal(t, 1, group.Value())
	assert.True(t, group.Next())
	assert.Equal(t, 3, group.Value())
	assert.True(t, group.Next())
	assert.Equal(t, 5, group.Value())
	assert.False(t, group.Next())

	assert.False(t, it.Next())
}

func TestGroupBy_Some(t *testing.T) {
	t.Parallel()

	it := GroupBy[int, bool](
		Slice([]int{1, 3, 5, 2, 4, 7}),
		func(i int) bool { return i%2 == 0 },
	)

	assert.True(t, it.Next())
	group := it.Value()
	assert.False(t, group.Key)
	assert.True(t, group.Next())
	assert.Equal(t, 1, group.Value())
	assert.True(t, group.Next())
	assert.Equal(t, 3, group.Value())

	it.Next()
	group = it.Value()
	assert.True(t, group.Key)
	assert.True(t, group.Next())
	assert.Equal(t, 2, group.Value())
	assert.True(t, group.Next())
	assert.Equal(t, 4, group.Value())
	assert.False(t, group.Next())

	assert.True(t, it.Next())
	group = it.Value()
	assert.False(t, group.Key)
	assert.True(t, group.Next())
	assert.Equal(t, 7, group.Value())

	assert.False(t, it.Next())
}

func ExampleGroupBy() {
	// GroupBy groups consecutive elements of the input iterator into iterable groups based on the provided mapping
	// function.
	it := GroupBy[int, bool](
		Slice([]int{1, 3, 5, 2, 4, 7}),
		func(i int) bool { return i%2 == 0 },
	)

	// Iterate over the groups.
	for it.Next() {
		group := it.Value()
		// Iterate over the elements in the group.
		for group.Next() {
			fmt.Print(group.Value())
		}
		fmt.Println()
	}

	// Output:
	// 135
	// 24
	// 7
}
