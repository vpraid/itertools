package iterator

// ZipIterator is an iterator that iterates over two iterators in parallel.
type ZipIterator[T, U any] struct {
	it1 Iterator[T]
	it2 Iterator[U]
}

// Zip returns a ZipIterator.
func Zip[T, U any](it1 Iterator[T], it2 Iterator[U]) *ZipIterator[T, U] {
	return &ZipIterator[T, U]{
		it1: it1,
		it2: it2,
	}
}

// Next advances the iterator to the next element. It returns false if one of the underlying iterators was exhausted.
func (zi *ZipIterator[T, U]) Next() bool {
	return zi.it1.Next() && zi.it2.Next()
}

// Pair represents a pair of elements from two iterators.
type Pair[T, U any] struct {
	First  T
	Second U
}

// Value returns the current element.
func (zi *ZipIterator[T, U]) Value() Pair[T, U] {
	return Pair[T, U]{First: zi.it1.Value(), Second: zi.it2.Value()}
}

// Collect returns the elements of the underlying iterator as a slice.
func (zi *ZipIterator[T, U]) Collect() []Pair[T, U] {
	return CollectFromIter[Pair[T, U]](zi)
}

// Chan returns a channel that yields the elements of the underlying iterator.
func (zi *ZipIterator[T, U]) Chan() <-chan Pair[T, U] {
	return ChanFromIter[Pair[T, U]](zi)
}
