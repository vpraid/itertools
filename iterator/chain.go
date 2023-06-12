package iterator

type ChainIterator[T any] struct {
	head Iterator[T]
	rest []Iterator[T]
}

// Chain returns an iterator that chains the elements of the given iterators.
func Chain[T any](it1, it2 Iterator[T], its ...Iterator[T]) *ChainIterator[T] {
	if len(its) == 0 {
		return &ChainIterator[T]{
			head: it1,
			rest: []Iterator[T]{it2},
		}
	}
	return &ChainIterator[T]{
		head: it1,
		rest: append([]Iterator[T]{it2}, its...),
	}
}

// Next advances the iterator to the next element.
func (ci *ChainIterator[T]) Next() bool {
	if ci.head.Next() {
		return true
	}
	if len(ci.rest) == 0 {
		return false
	}
	ci.head = ci.rest[0]
	ci.rest = ci.rest[1:]
	return ci.Next()
}

// Value returns the current element.
func (ci *ChainIterator[T]) Value() T {
	return ci.head.Value()
}

// Collect returns all remaining elements as a slice.
func (ci *ChainIterator[T]) Collect() []T {
	return CollectFromIter[T](ci)
}

// Chan returns a channel that will receive the elements of the underlying iterator.
func (ci *ChainIterator[T]) Chan() <-chan T {
	channel := make(chan T)
	go func() {
		defer close(channel)
		for ci.Next() {
			channel <- ci.Value()
		}
	}()
	return channel
}
