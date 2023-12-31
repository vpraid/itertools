package iterator

// FilterIterator is an iterator that filters the elements of af the input iterator.
type FilterIterator[T any] struct {
	it    Iterator[T]
	p     func(T) bool
	value T
}

// Filter returns an iterator that filters the elements of the input iterator with the given predicate.
func Filter[T any](it Iterator[T], p func(T) bool) *FilterIterator[T] {
	return &FilterIterator[T]{
		it: it,
		p:  p,
	}
}

// Next advances the iterator to the next element.
func (fi *FilterIterator[T]) Next() bool {
	if fi.it == nil {
		return false
	}
	for fi.it.Next() {
		value := fi.it.Value()
		if fi.p(value) {
			fi.value = value
			return true
		}
	}
	return false
}

// Value returns the current element.
func (fi *FilterIterator[T]) Value() T {
	return fi.value
}

// Collect returns the elements of the underlying iterator as a slice.
func (fi *FilterIterator[T]) Collect() []T {
	return CollectFromIter[T](fi)
}

// Chan returns a channel that will receive the elements of the underlying iterator.
func (fi *FilterIterator[T]) Chan() <-chan T {
	return ChanFromIter[T](fi)
}

// Bind replaces the underlying iterator with the given one.
func (fi *FilterIterator[T]) Bind(it Iterator[T]) {
	fi.it = it
}
