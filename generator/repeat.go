package generator

import "github.com/vpraid/itertools/iterator"

type RepeatIterator[T any] struct {
	value T
}

// Repeat returns a RepeatIterator for the given value.
func Repeat[T any](value T) *RepeatIterator[T] {
	return &RepeatIterator[T]{
		value: value,
	}
}

// Next advances the iterator to the next value of the underlying slice. It returns false if the slice was exhausted.
func (ri *RepeatIterator[T]) Next() bool {
	return true
}

// Value returns the same value over and over again.
func (ri *RepeatIterator[T]) Value() T {
	return ri.value
}

// Chan returns a channel that yields the elements of the underlying iterator.
func (ri *RepeatIterator[T]) Chan() <-chan T {
	return iterator.ChanFromIter[T](ri)
}
