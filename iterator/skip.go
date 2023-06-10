package iterator

// SkipIterator is an iterator that skips the first n elements of an iterator.
type SkipIterator[T any] struct {
	it Iterator[T]
	n  int
}

// Skip returns a SkipIterator for the given n.
func Skip[T any](it Iterator[T], n int) *SkipIterator[T] {
	return &SkipIterator[T]{it, n}
}

// Next advances the iterator to the next element. It returns false if the underlying iterator was exhausted.
// It will continue skipping elements until the underlying iterator is exhausted or n elements have been skipped.
func (it *SkipIterator[T]) Next() bool {
	for it.n > 0 && it.it.Next() {
		it.n--
	}
	return it.it.Next()
}

// Value returns the current element of the underlying iterator.
func (it *SkipIterator[T]) Value() T {
	return it.it.Value()
}
