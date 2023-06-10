package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/pkg/source"
)

func TestMap_Empty(t *testing.T) {
	t.Parallel()

	s := source.Slice([]string{})
	m := Map[string, int](s, func(t string) int { return len(t) })
	assert.False(t, m.Next())
	assert.Equal(t, 0, m.Value())
}

func TestMap_NonEmpty(t *testing.T) {
	t.Parallel()

	s := source.Slice([]string{"a", "bb", "ccc"})
	m := Map[string, int](s, func(t string) int { return len(t) })
	assert.True(t, m.Next())
	assert.Equal(t, 1, m.Value())
	assert.True(t, m.Next())
	assert.Equal(t, 2, m.Value())
	assert.True(t, m.Next())
	assert.Equal(t, 3, m.Value())
	assert.False(t, m.Next())
	assert.Equal(t, 0, m.Value())
}

func TestMap_Collect(t *testing.T) {
	t.Parallel()

	s := source.Slice([]string{"a", "bb", "ccc"})
	m := Map[string, int](s, func(t string) int { return len(t) })
	assert.Equal(t, []int{1, 2, 3}, m.Collect())
}

func TestMap_Bind(t *testing.T) {
	t.Parallel()

	s := source.Slice([]string{"a", "bb", "ccc"})
	m := Map[string, int](s, func(t string) int { return len(t) })
	assert.Equal(t, []int{1, 2, 3}, m.Collect())
	m.Bind(source.Slice([]string{"aaa", "bb", "c"}))
	assert.Equal(t, []int{3, 2, 1}, m.Collect())
}

func ExampleMap() {
	s := source.Slice([]string{"a", "bb", "ccc"})
	// We need to specify the type of the iterator explicitly because the compiler cannot infer it yet. This is a known
	// limitation of the Go compiler which will be fixed in Go 1.21.
	m := Map[string](s, func(t string) int { return len(t) })
	for m.Next() {
		fmt.Println(m.Value())
	}
	// Output:
	// 1
	// 2
	// 3
}
