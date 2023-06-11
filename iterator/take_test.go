package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/source"
)

func TestTake_ZeroEmpty(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{})
	take := Take[int](s, 0)
	assert.False(t, take.Next())
}

func TestTake_ZeroSome(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2})
	take := Take[int](s, 0)
	assert.False(t, take.Next())
}

func TestTake_Empty(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{})
	take := Take[int](s, 2)
	assert.False(t, take.Next())
}

func TestTake_Few(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3})
	take := Take[int](s, 1)
	assert.True(t, take.Next())
	assert.Equal(t, 1, take.Value())
	assert.False(t, take.Next())
}

func TestTake_All(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3})
	take := Take[int](s, 2)
	assert.True(t, take.Next())
	assert.Equal(t, 1, take.Value())
	assert.True(t, take.Next())
	assert.Equal(t, 2, take.Value())
	assert.False(t, take.Next())
}

func TestTake_Many(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3})
	take := Take[int](s, 5)
	assert.True(t, take.Next())
	assert.Equal(t, 1, take.Value())
	assert.True(t, take.Next())
	assert.Equal(t, 2, take.Value())
	assert.True(t, take.Next())
	assert.Equal(t, 3, take.Value())
	assert.False(t, take.Next())
}

func TestTake_Collect(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3})
	take := Take[int](s, 2)
	assert.Equal(t, []int{1, 2}, take.Collect())
}

func TestTake_Bind(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3})
	take := Take[int](s, 2)
	assert.True(t, take.Next())
	assert.Equal(t, 1, take.Value())
	take.Bind(source.Slice([]int{4, 5, 6}))
	assert.True(t, take.Next())
	assert.Equal(t, 4, take.Value())
	assert.False(t, take.Next())
}

func ExampleTake() {
	s := source.Slice([]int{1, 2, 3, 4, 5})
	// We need to specify the type of the iterator explicitly because the compiler cannot infer it yet. This is a known
	// limitation of the Go compiler which will be fixed in Go 1.21.
	ti := Take[int](s, 2)
	for ti.Next() {
		fmt.Println(ti.Value())
	}
	// Output:
	// 1
	// 2
}
