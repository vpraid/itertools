package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/source"
)

func TestSkip_ZeroEmpty(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{})
	it := Skip[int](s, 0)
	assert.False(t, it.Next())
}

func TestSkip_ZeroSome(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2})
	it := Skip[int](s, 0)
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Value())
	assert.False(t, it.Next())
}

func TestSkip_Empty(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{})
	it := Skip[int](s, 2)
	assert.False(t, it.Next())
}

func TestSkip_Few(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3, 4})
	it := Skip[int](s, 2)
	assert.True(t, it.Next())
	assert.Equal(t, 3, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 4, it.Value())
	assert.False(t, it.Next())
}

func TestSkip_Many(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3, 4})
	it := Skip[int](s, 5)
	assert.False(t, it.Next())
}

func TestSkip_Collect(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3, 4})
	it := Skip[int](s, 2)
	assert.Equal(t, []int{3, 4}, it.Collect())
}

func TestSkip_Chan(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3, 4})
	it := Skip[int](s, 2)
	c := it.Chan()
	assert.Equal(t, 3, <-c)
	assert.Equal(t, 4, <-c)
}

func TestSkip_Bind(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3, 4})
	it := Skip[int](s, 2)
	assert.True(t, it.Next())
	assert.Equal(t, 3, it.Value())
	it.Bind(source.Slice([]int{5, 6, 7}))
	assert.True(t, it.Next())
	assert.Equal(t, 5, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 6, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 7, it.Value())
	assert.False(t, it.Next())
}

func ExampleSkip() {
	s := source.Slice([]int{1, 2, 3, 4, 5})
	// We need to specify the type of the iterator explicitly because the compiler cannot infer it yet. This is a known
	// limitation of Go.
	it := Skip[int](s, 2)
	for it.Next() {
		fmt.Println(it.Value())
	}
	// Output:
	// 3
	// 4
	// 5
}
