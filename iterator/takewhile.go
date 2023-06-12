package iterator

type TakeWhileIterator[T any] struct {
	it   Iterator[T]
	pred func(T) bool
}

// TakeWhile returns a TakeWhileIterator for the given predicate.
func TakeWhile[T any](it Iterator[T], pred func(T) bool) *TakeWhileIterator[T] {
	return &TakeWhileIterator[T]{
		it:   it,
		pred: pred,
	}
}

// Next advances the iterator to the next element. It returns false if the underlying iterator was exhausted.
func (ti *TakeWhileIterator[T]) Next() bool {
	if ti.it == nil {
		return false
	}
	return ti.it.Next() && ti.pred(ti.it.Value())
}

// Value returns the current element of the underlying iterator.
func (ti *TakeWhileIterator[T]) Value() T {
	return ti.it.Value()
}

// Collect returns the elements of the underlying iterator as a slice.
func (ti *TakeWhileIterator[T]) Collect() []T {
	return CollectFromIter[T](ti)
}

// Chan returns a channel that will receive the elements of the underlying iterator.
func (ti *TakeWhileIterator[T]) Chan() <-chan T {
	return ChanFromIter[T](ti)
}

// Bind replaces the underlying iterator with the given one.
func (ti *TakeWhileIterator[T]) Bind(it Iterator[T]) {
	ti.it = it
}
