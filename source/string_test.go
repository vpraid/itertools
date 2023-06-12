package source

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString_Empty(t *testing.T) {
	t.Parallel()

	si := String("")
	assert.False(t, si.Next())
}

func TestString_NonEmpty(t *testing.T) {
	t.Parallel()

	si := String("abc")
	assert.True(t, si.Next())
	assert.Equal(t, 'a', si.Value())
	assert.True(t, si.Next())
	assert.Equal(t, 'b', si.Value())
	assert.True(t, si.Next())
	assert.Equal(t, 'c', si.Value())
	assert.False(t, si.Next())
}

func TestString_Collect(t *testing.T) {
	t.Parallel()

	si := String("abc")
	assert.Equal(t, []rune{'a', 'b', 'c'}, si.Collect())
}

func TestString_Chan(t *testing.T) {
	t.Parallel()

	ch := String("abc").Chan()
	assert.Equal(t, 'a', <-ch)
	assert.Equal(t, 'b', <-ch)
	assert.Equal(t, 'c', <-ch)
}
