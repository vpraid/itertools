package generator

import (
	"github.com/vpraid/itertools/iterator"
	"golang.org/x/exp/constraints"
)

type CountIterator[T constraints.Integer] struct {
	start T
	step  T
	value T
}

// Count returns an iterator that counts up from the given start value.
func Count[T constraints.Integer](start T) *CountIterator[T] {
	return &CountIterator[T]{
		start: start,
		step:  1,
		value: start,
	}
}

// CountBy returns an iterator that counts up from the given start value with the given step.
func CountBy[T constraints.Integer](start, step T) *CountIterator[T] {
	return &CountIterator[T]{
		start: start,
		step:  step,
		value: start,
	}
}

// Next advances the iterator to the next element.
func (ci *CountIterator[T]) Next() bool {
	ci.value += ci.step
	return true
}

// Value returns the current element.
func (ci *CountIterator[T]) Value() T {
	return ci.value - ci.step
}

// Chan returns a channel that yields the elements of the underlying iterator.
func (ci *CountIterator[T]) Chan() <-chan T {
	return iterator.ChanFromIter[T](ci)
}
