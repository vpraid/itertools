package functional

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/iterator"
	"github.com/vpraid/itertools/partial"
	"github.com/vpraid/itertools/source"
)

func TestCompose(t *testing.T) {
	it := Compose4[int, int, int, int](
		source.Slice([]int{1, 2, 3, 4, 5}),
		partial.Take[int](3),
		partial.Skip[int](1),
		partial.Map(func(i int) int { return i * 2 }),
	)
	assert.Equal(t, []int{4, 6, 8}, it.Collect())
}

func ExampleCompose4() {
	it := Compose4[string, string, int, *iterator.Group[int, bool]](
		source.Slice([]string{"a", "bb", "cc", "ddd", "eee"}),
		partial.Filter(func(s string) bool { return len(s) > 1 }),
		partial.Map(func(s string) int { return len(s) }),
		partial.GroupBy(func(i int) bool { return i%2 == 0 }),
	)
	for it.Next() {
		fmt.Println(it.Value().Collect())
	}
	// Output:
	// [2 2]
	// [3 3]
}
