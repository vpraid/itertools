package functional

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/source"
)

func TestAny(t *testing.T) {
	t.Parallel()

	assert.True(t, Any[int](source.Literal(1, 2, 3), func(i int) bool { return i == 2 }))
	assert.False(t, Any[int](source.Literal(1, 2, 3), func(i int) bool { return i == 4 }))
}

func TestAll(t *testing.T) {
	t.Parallel()

	assert.True(t, All[int](source.Literal(1, 2, 3), func(i int) bool { return i > 0 }))
	assert.False(t, All[int](source.Literal(1, 2, 3), func(i int) bool { return i < 3 }))
}
