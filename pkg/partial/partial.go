package partial

import "github.com/vpraid/itertools/pkg/iterator"

// Filter returns a MapIterator without an underlying iterator. Calling Next on it will always return false.
func Filter[T any](predicate func(T) bool) *iterator.FilterIterator[T] {
	return iterator.Filter[T](nil, predicate)
}

// Map returns a MapIterator without an underlying iterator. Calling Next on it will always return false.
func Map[T, U any](fn func(T) U) *iterator.MapIterator[T, U] {
	return iterator.Map[T, U](nil, fn)
}

// Skip returns a SkipIterator without an underlying iterator. Calling Next on it will always return false.
func Skip[T any](n int) *iterator.SkipIterator[T] {
	return iterator.Skip[T](nil, n)
}

// Take returns a TakeIterator without an underlying iterator. Calling Next on it will always return false.
func Take[T any](n int) *iterator.TakeIterator[T] {
	return iterator.Take[T](nil, n)
}

// PeekAhead returns a PeekAheadIterator without an underlying iterator. Calling Next on it will always return false.
func PeekAhead[T any](it iterator.Iterator[T]) *iterator.PeekAheadIterator[T] {
	return iterator.PeekAhead[T](nil)
}

// GroupBy returns a GroupByIterator without an underlying iterator. Calling Next on it will always return false.
func GroupBy[T any, U comparable](fn func(T) U) *iterator.GroupByIterator[T, U] {
	return iterator.GroupBy[T, U](nil, fn)
}

// TakeWhile returns a TakeWhileIterator without an underlying iterator. Calling Next on it will always return false.
func TakeWhile[T any](pred func(T) bool) *iterator.TakeWhileIterator[T] {
	return iterator.TakeWhile[T](nil, pred)
}
