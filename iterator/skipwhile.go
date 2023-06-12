package iterator

type SkipWhileIterator[T any] struct {
	it      Iterator[T]
	pred    func(T) bool
	skipped bool
}

// SkipWhile returns a SkipWhileIterator for the given predicate.
func SkipWhile[T any](it Iterator[T], pred func(T) bool) *SkipWhileIterator[T] {
	return &SkipWhileIterator[T]{
		it:   it,
		pred: pred,
	}
}

// Next advances the iterator to the next element. It returns false if the underlying iterator was exhausted.
func (si *SkipWhileIterator[T]) Next() bool {
	if si.it == nil {
		return false
	}
	if si.skipped {
		return si.it.Next()
	}
	for si.it.Next() {
		if !si.pred(si.it.Value()) {
			si.skipped = true
			return true
		}
	}
	return false
}

// Value returns the current element of the underlying iterator.
func (si *SkipWhileIterator[T]) Value() T {
	return si.it.Value()
}

// Collect returns the elements of the underlying iterator as a slice.
func (si *SkipWhileIterator[T]) Collect() []T {
	return CollectFromIter[T](si)
}

// Chan returns a channel that will receive the elements of the underlying iterator.
func (si *SkipWhileIterator[T]) Chan() <-chan T {
	return ChanFromIter[T](si)
}

// Bind replaces the underlying iterator with the given one.
func (si *SkipWhileIterator[T]) Bind(it Iterator[T]) {
	si.it = it
}
