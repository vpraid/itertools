package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/source"
)

func TestSkipWhile_Empty(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{})
	it := SkipWhile[int](s, func(i int) bool { return true })
	assert.False(t, it.Next())
}

func TestSkipWhile_Some(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 3, 5, 4, 6})
	it := SkipWhile[int](s, func(i int) bool { return i%2 == 1 })
	assert.True(t, it.Next())
	assert.Equal(t, 4, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 6, it.Value())
	assert.False(t, it.Next())
}

func TestSkipWhile_All(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 3, 5, 4, 6})
	it := SkipWhile[int](s, func(i int) bool { return i < 10 })
	assert.False(t, it.Next())
}

func TestSkipWhile_None(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 3, 5})
	it := SkipWhile[int](s, func(i int) bool { return i > 10 })
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 3, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 5, it.Value())
	assert.False(t, it.Next())
}

func TestSkipWhile_Collect(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 3, 5, 4, 6})
	it := SkipWhile[int](s, func(i int) bool { return i%2 == 1 })
	assert.Equal(t, []int{4, 6}, it.Collect())
}

func TestSkipWhile_Chan(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 3, 5, 4, 6})
	it := SkipWhile[int](s, func(i int) bool { return i%2 == 1 })
	c := it.Chan()
	assert.Equal(t, 4, <-c)
	assert.Equal(t, 6, <-c)
}

func ExampleSkipWhile() {
	s := source.Slice([]int{1, 3, 5, 4, 6})
	// We need to specify the type of the iterator explicitly because the compiler cannot infer it yet. This is a known
	// limitation of the Go compiler which will be fixed in Go 1.21.
	it := SkipWhile[int](s, func(i int) bool { return i%2 == 1 })
	for it.Next() {
		fmt.Println(it.Value())
	}
	// Output:
	// 4
	// 6
}
