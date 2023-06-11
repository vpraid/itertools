package partial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	it := Filter[int](func(i int) bool { return i%2 == 0 })
	assert.False(t, it.Next())
}

func TestMap(t *testing.T) {
	it := Map[int, int](func(i int) int { return i * 2 })
	assert.False(t, it.Next())
}

func TestSkip(t *testing.T) {
	it := Skip[int](5)
	assert.False(t, it.Next())
}

func TestTake(t *testing.T) {
	it := Take[int](5)
	assert.False(t, it.Next())
}

func TestPeekAhead(t *testing.T) {
	it := PeekAhead[int]()
	assert.False(t, it.Next())
}

func TestGroupBy(t *testing.T) {
	it := GroupBy[int, int](func(i int) int { return i % 2 })
	assert.False(t, it.Next())
}

func TestTakeWhile(t *testing.T) {
	it := TakeWhile[int](func(i int) bool { return i%2 == 0 })
	assert.False(t, it.Next())
}

func TestSkipWhile(t *testing.T) {
	it := SkipWhile[int](func(i int) bool { return i%2 == 0 })
	assert.False(t, it.Next())
}
