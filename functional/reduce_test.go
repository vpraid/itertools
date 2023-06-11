package functional

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/source"
)

func TestFold(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3})
	it := Fold[int, int](s, 1, func(acc, x int) int { return acc + x })
	assert.Equal(t, 7, it)
}

func TestFold_Empty(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{})
	it := Fold[int, int](s, 5, func(acc, x int) int { return acc + x })
	assert.Equal(t, 5, it)
}

func TestReduce(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3})
	it := Reduce[int](s, func(acc, x int) int { return acc + x })
	assert.Equal(t, 6, it)
}

func ExampleFold() {
	s := source.Slice([]int{1, 2, 3})
	it := Fold[int, []int](s, []int{}, func(acc []int, x int) []int { return prependInt(acc, x) })
	fmt.Println(it)
	// Output: [3 2 1]
}

// prependInt uses an efficient method to prepend an int to a slice by reusing slice capacity as much as possible.
func prependInt(xs []int, y int) []int {
	xs = append(xs, 0)
	copy(xs[1:], xs)
	xs[0] = y
	return xs
}
