package iterator

type InterleaveIterator[T any] struct {
	its   []Iterator[T]
	index int
	value T
}

// Interleave returns an iterator that interleaves the elements of the input iterators.
func Interleave[T any](its ...Iterator[T]) *InterleaveIterator[T] {
	return &InterleaveIterator[T]{
		its:   its,
		index: 0,
	}
}

// Next advances the iterator to the next element.
func (ii *InterleaveIterator[T]) Next() bool {
	if len(ii.its) == 0 {
		return false
	}
	it := ii.its[ii.index]
	if it.Next() {
		ii.value = it.Value()
		ii.index = (ii.index + 1) % len(ii.its)
		return true
	}
	return false
}

// Value returns the current element.
func (ii *InterleaveIterator[T]) Value() T {
	return ii.value
}

// Collect returns all remaining elements as a slice.
func (ii *InterleaveIterator[T]) Collect() []T {
	return CollectFromIter[T](ii)
}

// Chan returns a channel that will receive the remaining elements.
func (ii *InterleaveIterator[T]) Chan() <-chan T {
	return ChanFromIter[T](ii)
}
