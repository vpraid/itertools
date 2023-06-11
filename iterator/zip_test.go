package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/source"
)

func TestZip_Empty(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{})
	it := Zip[int, int](s, s)
	assert.False(t, it.Next())
}

func TestZip_Both(t *testing.T) {
	t.Parallel()

	s1 := source.Slice([]int{1, 2, 3})
	s2 := source.Slice([]int{4, 5, 6})
	it := Zip[int, int](s1, s2)
	assert.True(t, it.Next())
	assert.Equal(t, Pair[int, int]{1, 4}, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, Pair[int, int]{2, 5}, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, Pair[int, int]{3, 6}, it.Value())
	assert.False(t, it.Next())
}

func TestZip_OneShorter(t *testing.T) {
	t.Parallel()

	s1 := source.Slice([]int{1, 2, 3})
	s2 := source.Slice([]int{4, 5})
	it := Zip[int, int](s1, s2)
	assert.True(t, it.Next())
	assert.Equal(t, Pair[int, int]{1, 4}, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, Pair[int, int]{2, 5}, it.Value())
	assert.False(t, it.Next())
}

func ExampleZip() {
	s1 := source.Slice([]int{1, 2, 3})
	s2 := source.Slice([]int{4, 5, 6})
	it := Zip[int, int](s1, s2)
	for it.Next() {
		pair := it.Value()
		fmt.Println(pair.First, pair.Second)
	}
	// Output:
	// 1 4
	// 2 5
	// 3 6
}
