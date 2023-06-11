package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromFn_Empty(t *testing.T) {
	t.Parallel()

	it := FromFn[int](func() (int, bool) { return 0, false })
	assert.False(t, it.Next())
}

func TestFromFn_NonEmpty(t *testing.T) {
	t.Parallel()

	count := 0
	it := FromFn[int](func() (int, bool) {
		count++
		return count, count < 3
	})
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Value())
	assert.False(t, it.Next())
}
