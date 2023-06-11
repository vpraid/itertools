package iterator_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/functional"
	"github.com/vpraid/itertools/iterator"
	"github.com/vpraid/itertools/partial"
	"github.com/vpraid/itertools/source"
)

func TestGroupBy_Empty(t *testing.T) {
	t.Parallel()

	it := iterator.GroupBy[int, bool](
		source.Slice([]int{}),
		func(i int) bool { return i%2 == 0 },
	)

	assert.False(t, it.Next())
}

func TestGroupBy_OneGroup(t *testing.T) {
	t.Parallel()

	it := iterator.GroupBy[int, bool](
		source.Slice([]int{1, 3, 5}),
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

	it := iterator.GroupBy[int, bool](
		source.Slice([]int{1, 3, 5, 2, 4, 7}),
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

func TestGroupBy_Collect(t *testing.T) {
	t.Parallel()

	it := iterator.GroupBy[int, bool](
		source.Slice([]int{1, 3, 5, 2, 4, 7}),
		func(i int) bool { return i%2 == 0 },
	)

	assert.True(t, it.Next())
	assert.Equal(t, []int{1, 3, 5}, it.Value().Collect())
	assert.True(t, it.Next())
	assert.Equal(t, []int{2, 4}, it.Value().Collect())
	assert.True(t, it.Next())
	assert.Equal(t, []int{7}, it.Value().Collect())
	assert.False(t, it.Next())
}

func TestGroupBy_Detach(t *testing.T) {
	t.Parallel()

	it := iterator.GroupBy[int, bool](
		source.Slice([]int{1, 3, 5, 2, 4, 7}),
		func(i int) bool { return i%2 == 0 },
	)

	assert.True(t, it.Next())
	assert.Equal(t, []int{1, 3, 5}, it.Value().Detach().Collect())
	assert.True(t, it.Next())
	assert.Equal(t, []int{2, 4}, it.Value().Detach().Collect())
	assert.True(t, it.Next())
	assert.Equal(t, []int{7}, it.Value().Detach().Collect())
	assert.False(t, it.Next())
}

func TestGroupBy_Bind(t *testing.T) {
	t.Parallel()

	it := iterator.GroupBy[int, bool](
		source.Slice([]int{1, 3, 5, 2, 4, 7}),
		func(i int) bool { return i%2 == 0 },
	)

	assert.True(t, it.Next())
	assert.Equal(t, []int{1, 3, 5}, it.Value().Collect())
	it.Bind(source.Slice([]int{6, 8}))
	assert.True(t, it.Next())
	assert.Equal(t, []int{6, 8}, it.Value().Collect())
}

func ExampleGroupBy() {
	// GroupBy groups consecutive elements of the input iterator into iterable groups based on the provided mapping
	// function.
	it := iterator.GroupBy[int, bool](
		source.Slice([]int{1, 3, 5, 2, 4, 7}),
		func(i int) bool { return i%2 == 0 },
	)

	// Iterate over the groups.
	for it.Next() {
		group := it.Value()
		fmt.Println(group.Collect())
	}

	// Output:
	// [1 3 5]
	// [2 4]
	// [7]
}

func ExampleGroup() {
	// Transform initial iterator into groups of consecutive elements based on the provided mapping function, then
	// detach the groups and collect them into slices.
	it := functional.Compose3[int, *iterator.Group[int, bool], []int](
		source.Slice[int]([]int{1, 3, 5, 2, 4, 7}),
		partial.GroupBy[int, bool](func(i int) bool { return i%2 == 0 }),
		partial.Map(func(group *iterator.Group[int, bool]) []int { return group.Detach().Collect() }),
	)

	// Iterate over the groups.
	for it.Next() {
		fmt.Println(it.Value())
	}

	// Output:
	// [1 3 5]
	// [2 4]
	// [7]
}
