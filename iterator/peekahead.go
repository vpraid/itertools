package iterator

// PeekAheadIterator is an iterator that allows reading the next element without advancing the iterator.
type PeekAheadIterator[T any] struct {
	Iterator[T]
	value     T
	peekValue T
	exhausted bool
}

// PeekAhead returns a PeekAheadIterator.
func PeekAhead[T any](it Iterator[T]) *PeekAheadIterator[T] {
	exhausted := !it.Next()
	return &PeekAheadIterator[T]{
		Iterator:  it,
		peekValue: it.Value(),
		exhausted: exhausted,
	}
}

// Next advances the iterator to the next element. It returns false if the underlying iterator was exhausted.
// Under the hood it calls the underlying iterator's Next method to advance it one step further than what is
// visible to the user, while caching the value of the current and next element.
func (pai *PeekAheadIterator[T]) Next() bool {
	pai.value = pai.Iterator.Value()
	exhausted := pai.exhausted
	pai.exhausted = !pai.Iterator.Next()
	pai.peekValue = pai.Iterator.Value()
	return !exhausted
}

// Value returns the cached value of the current element.
func (pai *PeekAheadIterator[T]) Value() T {
	return pai.value
}

// Peek returns the cached value of the next element.
func (pai *PeekAheadIterator[T]) Peek() T {
	return pai.peekValue
}

// Exhausted returns true if the underlying iterator was exhausted.
func (pai *PeekAheadIterator[T]) Exhausted() bool {
	return pai.exhausted
}
