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
