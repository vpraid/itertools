package iterator

import "github.com/vpraid/itertools/internal/common"

// PeekAheadIterator is an iterator that allows reading the next element without advancing the iterator.
type PeekAheadIterator[T any] struct {
	it        Iterator[T]
	value     T
	peekValue T
	exhausted bool
}

// PeekAhead returns a PeekAheadIterator.
func PeekAhead[T any](it Iterator[T]) *PeekAheadIterator[T] {
	if it == nil {
		return &PeekAheadIterator[T]{
			it:        nil,
			value:     common.Zero[T](),
			peekValue: common.Zero[T](),
			exhausted: true,
		}
	}
	exhausted := !it.Next()
	return &PeekAheadIterator[T]{
		it:        it,
		peekValue: it.Value(),
		exhausted: exhausted,
	}
}

// Next advances the iterator to the next element. It returns false if the underlying iterator was exhausted.
// Under the hood it calls the underlying iterator's Next method to advance it one step further than what is
// visible to the user, while caching the value of the current and next element.
func (pai *PeekAheadIterator[T]) Next() bool {
	if pai.it == nil {
		return false
	}
	pai.value = pai.it.Value()
	exhausted := pai.exhausted
	pai.exhausted = !pai.it.Next()
	pai.peekValue = pai.it.Value()
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

// Collect returns the elements of the underlying iterator as a slice.
func (pai *PeekAheadIterator[T]) Collect() []T {
	return CollectFromIter[T](pai)
}

// Chan returns a channel that will receive the elements of the underlying iterator.
func (pai *PeekAheadIterator[T]) Chan() <-chan T {
	return ChanFromIter[T](pai)
}

// Bind replaces the underlying iterator with the given one. Elements of the underlying iterator
// will not be skipped anew if Next was already called.
func (pai *PeekAheadIterator[T]) Bind(it Iterator[T]) {
	pai.exhausted = !it.Next()
	pai.it = it
	pai.peekValue = it.Value()
}
