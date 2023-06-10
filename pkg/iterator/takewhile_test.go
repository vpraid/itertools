package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/pkg/source"
)

func TestTakeWhile_Empty(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{})
	it := TakeWhile[int](s, func(i int) bool { return true })
	assert.False(t, it.Next())
}

func TestTakeWhile_Some(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 3, 5, 4, 6})
	it := TakeWhile[int](s, func(i int) bool { return i%2 == 1 })
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 3, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 5, it.Value())
	assert.False(t, it.Next())
}

func TestTakeWhile_All(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 3, 6})
	it := TakeWhile[int](s, func(i int) bool { return i < 10 })
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 3, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 6, it.Value())
	assert.False(t, it.Next())
}

func TestTakeWhile_None(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 3, 5, 4, 6})
	it := TakeWhile[int](s, func(i int) bool { return i > 10 })
	assert.False(t, it.Next())
}

func ExampleTakeWhile() {
	s := source.Slice([]int{1, 3, 5, 4, 6})
	// We need to specify the type of the iterator explicitly because the compiler cannot infer it yet. This is a known
	// limitation of the Go compiler which will be fixed in Go 1.21.
	it := TakeWhile[int](s, func(i int) bool { return i%2 == 1 })
	for it.Next() {
		fmt.Println(it.Value())
	}
	// Output:
	// 1
	// 3
	// 5
}
