package generator

import (
	"github.com/vpraid/itertools/internal/common"
	"github.com/vpraid/itertools/iterator"
)

type FromFnIterator[T any] struct {
	fn    func() (T, bool)
	value T
}

// FromFn returns an iterator that generates elements by calling the given function.
func FromFn[T any](fn func() (T, bool)) *FromFnIterator[T] {
	return &FromFnIterator[T]{
		fn:    fn,
		value: common.Zero[T](),
	}
}

// Next advances the iterator to the next element.
func (ffi *FromFnIterator[T]) Next() bool {
	value, ok := ffi.fn()
	if !ok {
		return false
	}
	ffi.value = value
	return true
}

// Value returns the current element.
func (ffi *FromFnIterator[T]) Value() T {
	return ffi.value
}

// Collect returns all remaining elements in the iterator.
func (ffi *FromFnIterator[T]) Collect() []T {
	return iterator.CollectFromIter[T](ffi)
}

// Chan returns a channel that yields the elements of the underlying iterator.
func (ffi *FromFnIterator[T]) Chan() <-chan T {
	return iterator.ChanFromIter[T](ffi)
}
