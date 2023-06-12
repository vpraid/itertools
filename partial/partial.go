package partial

import "github.com/vpraid/itertools/iterator"

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
func PeekAhead[T any]() *iterator.PeekAheadIterator[T] {
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

// SkipWhile returns a SkipWhileIterator without an underlying iterator. Calling Next on it will always return false.
func SkipWhile[T any](pred func(T) bool) *iterator.SkipWhileIterator[T] {
	return iterator.SkipWhile[T](nil, pred)
}

// Chunks returns a ChunkIterator without an underlying iterator. Calling Next on it will always return false.
func Chunks[T any](n int) *iterator.ChunkIterator[T] {
	return iterator.Chunks[T](nil, n)
}

// StepBy returns a StepByIterator without an underlying iterator. Calling Next on it will always return false.
func StepBy[T any](n int) *iterator.StepByIterator[T] {
	return iterator.StepBy[T](nil, n)
}

// Windows returns a WindowIterator without an underlying iterator. Calling Next on it will always return false.
func Windows[T any](n int) *iterator.WindowIterator[T] {
	return iterator.Windows[T](nil, n)
}
