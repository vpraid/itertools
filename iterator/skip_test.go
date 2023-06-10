package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkip_ZeroEmpty(t *testing.T) {
	t.Parallel()

	s := Slice([]int{})
	it := Skip[int](s, 0)
	assert.False(t, it.Next())
}

func TestSkip_ZeroSome(t *testing.T) {
	t.Parallel()

	s := Slice([]int{1, 2})
	it := Skip[int](s, 0)
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Value())
	assert.False(t, it.Next())
}

func TestSkip_Empty(t *testing.T) {
	t.Parallel()

	s := Slice([]int{})
	it := Skip[int](s, 2)
	assert.False(t, it.Next())
}

func TestSkip_Few(t *testing.T) {
	t.Parallel()

	s := Slice([]int{1, 2, 3, 4})
	it := Skip[int](s, 2)
	assert.True(t, it.Next())
	assert.Equal(t, 3, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 4, it.Value())
	assert.False(t, it.Next())
}

func TestSkip_Many(t *testing.T) {
	t.Parallel()

	s := Slice([]int{1, 2, 3, 4})
	it := Skip[int](s, 5)
	assert.False(t, it.Next())
}

func ExampleSkip() {
	s := Slice([]int{1, 2, 3, 4, 5})
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
