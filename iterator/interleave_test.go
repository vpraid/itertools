package iterator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/generator"
	"github.com/vpraid/itertools/iterator"
	"github.com/vpraid/itertools/source"
)

func TestInterleave_Empty(t *testing.T) {
	t.Parallel()

	ii := iterator.Interleave[int]()
	assert.False(t, ii.Next())
	assert.Equal(t, 0, ii.Value())
}

func TestInterleave_NonEmpty(t *testing.T) {
	t.Parallel()

	ii := iterator.Interleave[int](generator.Count(0), generator.Count(1), generator.Count(2))
	assert.True(t, ii.Next())
	assert.Equal(t, 0, ii.Value())
	assert.True(t, ii.Next())
	assert.Equal(t, 1, ii.Value())
	assert.True(t, ii.Next())
	assert.Equal(t, 2, ii.Value())
	assert.True(t, ii.Next())
	assert.Equal(t, 1, ii.Value())
	assert.True(t, ii.Next())
	assert.Equal(t, 2, ii.Value())
	assert.True(t, ii.Next())
	assert.Equal(t, 3, ii.Value())
}

func TestInterleave_Collect(t *testing.T) {
	t.Parallel()

	s1 := source.Literal(1, 2)
	s2 := source.Literal(3, 4)
	ii := iterator.Interleave[int](s1, s2)
	assert.Equal(t, []int{1, 3, 2, 4}, ii.Collect())
}

func TestInterleave_Chan(t *testing.T) {
	t.Parallel()

	s1 := source.Literal(1, 2)
	s2 := source.Literal(3, 4)
	ii := iterator.Interleave[int](s1, s2)
	c := ii.Chan()
	assert.Equal(t, 1, <-c)
	assert.Equal(t, 3, <-c)
	assert.Equal(t, 2, <-c)
	assert.Equal(t, 4, <-c)
}
