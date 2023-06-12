package generator

import (
	"github.com/vpraid/itertools/internal/common"
	"github.com/vpraid/itertools/iterator"
)

type CycleIterator[T any] struct {
	elements []T
	index    int
	value    T
}

// Cycle returns an iterator that endlessly cycles through the elements of the input slice.
func Cycle[T any](elements []T) *CycleIterator[T] {
	return &CycleIterator[T]{
		elements: elements,
		index:    -1,
		value:    common.Zero[T](),
	}
}

// CycleLiteral returns an iterator that endlessly cycles through the given elements.
func CycleLiteral[T any](elements ...T) *CycleIterator[T] {
	return Cycle[T](elements)
}

// Next advances the iterator to the next element.
func (ci *CycleIterator[T]) Next() bool {
	if len(ci.elements) == 0 {
		return false
	}
	ci.index = (ci.index + 1) % len(ci.elements)
	ci.value = ci.elements[ci.index]
	return true
}

// Value returns the current element.
func (ci *CycleIterator[T]) Value() T {
	return ci.value
}

// Chan returns a channel that yields the elements of the underlying iterator.
func (ci *CycleIterator[T]) Chan() <-chan T {
	return iterator.ChanFromIter[T](ci)
}
