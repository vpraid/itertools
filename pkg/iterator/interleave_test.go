package iterator_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/pkg/generator"
	"github.com/vpraid/itertools/pkg/iterator"
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
