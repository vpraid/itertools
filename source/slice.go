package source

import "github.com/vpraid/itertools/internal/common"

// SliceIterator is an iterator that iterates over a given slice.
type SliceIterator[T any] struct {
	elements []T
	value    T
}

// Slice returns a SliceIterator for the given slice.
func Slice[T any](elements []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		elements: elements,
	}
}

// Next advances the iterator to the next element of the underlying slice. It returns false if
// the underlying iterator was exhausted.
func (si *SliceIterator[T]) Next() bool {
	if len(si.elements) == 0 {
		si.value = common.Zero[T]()
		return false
	}
	si.value = si.elements[0]
	si.elements = si.elements[1:]
	return true
}

// Value returns the current element of the slice pointed by the iterator. If the iterator was exhausted, it returns
// the zero value.
func (si *SliceIterator[T]) Value() T {
	return si.value
}

// Collect returns the underlying slice as is.
func (si *SliceIterator[T]) Collect() []T {
	return si.elements
}

// Chan returns a channel that yields the elements of the underlying iterator.
func (si *SliceIterator[T]) Chan() <-chan T {
	channel := make(chan T)
	go func() {
		defer close(channel)
		for si.Next() {
			channel <- si.Value()
		}
	}()
	return channel
}
