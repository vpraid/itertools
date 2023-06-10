package iterator

// TakeIterator is an iterator that takes the first n elements of an iterator.
type TakeIterator[T any] struct {
	it   Iterator[T]
	size int
}

// Take returns a TakeIterator for the given n.
func Take[T any](it Iterator[T], size int) *TakeIterator[T] {
	return &TakeIterator[T]{
		it:   it,
		size: size,
	}
}

// EmptyTake returns a TakeIterator without an underlying iterator. Calling Next on it will always return false.
func EmptyTake[T any](size int) *TakeIterator[T] {
	return &TakeIterator[T]{
		it:   nil,
		size: size,
	}
}

// Next advances the iterator to the next element. It returns false if the underlying iterator was exhausted.
func (ti *TakeIterator[T]) Next() bool {
	if ti.it == nil {
		return false
	}
	return ti.size > 0 && ti.it.Next()
}

// Value returns the current element of the underlying iterator.
func (ti *TakeIterator[T]) Value() T {
	ti.size--
	return ti.it.Value()
}

// Collect returns the elements of the underlying iterator as a slice.
func (ti *TakeIterator[T]) Collect() []T {
	return CollectFromIter[T](ti)
}

// Imbue replaces the underlying iterator with the given one. If the iterator was partially or fully exhausted,
// the new iterator will continue where the old one left off. The counter will not change, but the elements returned
// by Value will be from the new iterator.
func (ti *TakeIterator[T]) Imbue(it Iterator[T]) {
	ti.it = it
}
