package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/source"
)

func TestFilter_Empty(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{})
	fi := Filter[int](s, func(i int) bool { return i%2 == 0 })
	assert.False(t, fi.Next())
	assert.Equal(t, 0, fi.Value())
}

func TestFilter_NonEmpty(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3, 4, 5})
	fi := Filter[int](s, func(i int) bool { return i%2 == 0 })
	assert.True(t, fi.Next())
	assert.Equal(t, 2, fi.Value())
	assert.True(t, fi.Next())
	assert.Equal(t, 4, fi.Value())
	assert.False(t, fi.Next())
}

func TestFilter_Collect(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3, 4, 5})
	fi := Filter[int](s, func(i int) bool { return i%2 == 0 })
	assert.Equal(t, []int{2, 4}, fi.Collect())
}

func TestFilter_Chan(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3, 4, 5})
	fi := Filter[int](s, func(i int) bool { return i%2 == 0 })
	c := fi.Chan()
	assert.Equal(t, 2, <-c)
	assert.Equal(t, 4, <-c)
}

func TestFilter_Bind(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3, 4, 5})
	fi := Filter[int](s, func(i int) bool { return i%2 == 0 })
	assert.True(t, fi.Next())
	assert.Equal(t, 2, fi.Value())
	fi.Bind(source.Slice([]int{6, 7, 8}))
	assert.True(t, fi.Next())
	assert.Equal(t, 6, fi.Value())
}

func ExampleFilter() {
	s := source.Slice([]int{1, 2, 3, 4, 5})
	// We need to specify the type of the iterator explicitly because the compiler cannot infer it yet. This is a known
	// limitation of the Go compiler which will be fixed in Go 1.21.
	fi := Filter[int](s, func(i int) bool { return i%2 == 0 })
	for fi.Next() {
		fmt.Println(fi.Value())
	}
	// Output:
	// 2
	// 4
}
