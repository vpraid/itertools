package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRange_ZeroRange(t *testing.T) {
	t.Parallel()

	s := Range(0, 0)
	assert.False(t, s.Next())
	assert.Equal(t, 0, s.Value())
}

func TesstRange_EmptyRange(t *testing.T) {
	t.Parallel()

	s := Range(1, 0)
	assert.False(t, s.Next())
	assert.Equal(t, 0, s.Value())
}

func TestRange_NonEmpty(t *testing.T) {
	t.Parallel()

	s := Range(1, 4)
	assert.True(t, s.Next())
	assert.Equal(t, 1, s.Value())
	assert.True(t, s.Next())
	assert.Equal(t, 2, s.Value())
	assert.True(t, s.Next())
	assert.Equal(t, 3, s.Value())
	assert.False(t, s.Next())
}

func TestRange_WithStep(t *testing.T) {
	t.Parallel()

	s := RangeWithStep(1, 4, 2)
	assert.True(t, s.Next())
	assert.Equal(t, 1, s.Value())
	assert.True(t, s.Next())
	assert.Equal(t, 3, s.Value())
	assert.False(t, s.Next())
}

func TestRange_WithNegativeStep(t *testing.T) {
	t.Parallel()

	s := RangeWithStep(4, 1, -2)
	assert.True(t, s.Next())
	assert.Equal(t, 4, s.Value())
	assert.True(t, s.Next())
	assert.Equal(t, 2, s.Value())
	assert.False(t, s.Next())
}

func TestRange_WithStepEmptyRange1(t *testing.T) {
	t.Parallel()

	s := RangeWithStep(1, 4, -2)
	assert.False(t, s.Next())
}

func TestRange_WithStepEmptyRange2(t *testing.T) {
	t.Parallel()

	s := RangeWithStep(5, 4, 2)
	assert.False(t, s.Next())
}

func TestRange_WithStepCollection(t *testing.T) {
	t.Parallel()

	s := RangeWithStep(1, 4, 2)
	assert.Equal(t, []int{1, 3}, s.Collect())
}
