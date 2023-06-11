package aggregate

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vpraid/itertools/source"
)

func TestSum(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 0, Sum[int](source.Literal[int]()))
	assert.Equal(t, 10, Sum[int](source.Literal[int](1, 2, 3, 4)))
}

func TestProduct(t *testing.T) {
	t.Parallel()

	assert.Equal(t, 1, Product[int](source.Literal[int]()))
	assert.Equal(t, 24, Product[int](source.Literal[int](1, 2, 3, 4)))
}
