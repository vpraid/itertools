package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter_Empty(t *testing.T) {
	t.Parallel()

	s := Slice([]int{})
	fi := Filter[int](s, func(i int) bool { return i%2 == 0 })
	assert.False(t, fi.Next())
	assert.Equal(t, 0, fi.Value())
}

func TestFilter_NonEmpty(t *testing.T) {
	t.Parallel()

	s := Slice([]int{1, 2, 3, 4, 5})
	fi := Filter[int](s, func(i int) bool { return i%2 == 0 })
	assert.True(t, fi.Next())
	assert.Equal(t, 2, fi.Value())
	assert.True(t, fi.Next())
	assert.Equal(t, 4, fi.Value())
	assert.False(t, fi.Next())
}

func ExampleFilter() {
	s := Slice([]int{1, 2, 3, 4, 5})
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
