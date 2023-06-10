package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPeekAhead_Empty(t *testing.T) {
	t.Parallel()

	it := PeekAhead[int](Slice([]int{}))
	assert.False(t, it.Next())
}

func TestPeekAhead_Some(t *testing.T) {
	t.Parallel()

	it := PeekAhead[int](Slice([]int{1, 2, 3}))
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())
	assert.Equal(t, 2, it.Peek())
	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Value())
	assert.Equal(t, 3, it.Peek())
	assert.True(t, it.Next())
	assert.Equal(t, 3, it.Value())
	assert.Equal(t, 0, it.Peek())
	assert.False(t, it.Next())
}

func TestPeekAhead_Collect(t *testing.T) {
	t.Parallel()

	it := PeekAhead[int](Slice([]int{1, 2, 3}))
	assert.Equal(t, []int{1, 2, 3}, it.Collect())
}

func TestPeekAhead_Imbued(t *testing.T) {
	t.Parallel()

	it := PeekAhead[int](Slice([]int{1, 2, 3}))
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())
	assert.Equal(t, 2, it.Peek())
	it.Imbue(Slice([]int{4, 5, 6}))
	assert.True(t, it.Next())
	assert.Equal(t, 4, it.Value())
	assert.Equal(t, 5, it.Peek())
}

func ExamplePeekAhead() {
	s := Slice([]int{1, 2, 3})
	// We need to specify the type of the iterator explicitly because the compiler cannot infer it yet. This is a known
	// limitation of Go.
	it := PeekAhead[int](s)
	for it.Next() {
		fmt.Println(it.Value())
		fmt.Println(it.Peek())
	}
	// Output:
	// 1
	// 2
	// 2
	// 3
	// 3
	// 0
}
