package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiteralIterator_Empty(t *testing.T) {
	t.Parallel()

	it := Literal[int]()
	assert.False(t, it.Next())
	assert.Equal(t, 0, it.Value())
}

func TestLiteralIterator_NonEmpty(t *testing.T) {
	t.Parallel()

	it := Literal(1, 2, 3)
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 3, it.Value())
	assert.False(t, it.Next())
	assert.Equal(t, 0, it.Value())
}

func TestLiteral_Collect(t *testing.T) {
	t.Parallel()

	it := Literal(1, 2, 3)
	assert.Equal(t, []int{1, 2, 3}, it.Collect())
}

func TestLiteral_Chan(t *testing.T) {
	t.Parallel()

	ch := Literal(1, 2, 3).Chan()
	assert.Equal(t, 1, <-ch)
	assert.Equal(t, 2, <-ch)
	assert.Equal(t, 3, <-ch)
}
