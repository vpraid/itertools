package iterator

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/source"
)

func TestStepBy_Empty(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{})
	it := StepBy[int](s, 2)
	assert.False(t, it.Next())
}

func TestStepBy_Few(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3})
	it := StepBy[int](s, 2)
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())
	assert.False(t, it.Next())
}

func TestStepBy_All(t *testing.T) {
	t.Parallel()

	s := source.Slice([]int{1, 2, 3})
	it := StepBy[int](s, 1)
	assert.True(t, it.Next())
	assert.Equal(t, 1, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 2, it.Value())
	assert.True(t, it.Next())
	assert.Equal(t, 3, it.Value())
	assert.False(t, it.Next())
}
