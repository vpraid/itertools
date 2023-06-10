package iterator

// MapIterator is an iterator that applies a function to each element of the underlying iterator.
type MapIterator[T any, U any] struct {
	it Iterator[T]
	fn func(T) U
}

// Map returns a MapIterator that applies fn to each element of the underlying iterator.
func Map[T any, U any](it Iterator[T], fn func(t T) U) *MapIterator[T, U] {
	return &MapIterator[T, U]{
		it: it,
		fn: fn,
	}
}

// Next advances the iterator to the next element. It returns false if the underlying iterator was exhausted.
func (mi *MapIterator[T, U]) Next() bool {
	return mi.it.Next()
}

// Value returns the current element of the underlying iterator after applying the mapping function.
func (mi *MapIterator[T, U]) Value() U {
	return mi.fn(mi.it.Value())
}

// Collect returns a slice containing all elements of the iterator.
func (mi *MapIterator[T, U]) Collect() []U {
	return CollectFromIter[U](mi)
}
