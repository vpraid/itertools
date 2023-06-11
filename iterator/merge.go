package iterator

type MergeIterator[U any] struct {
	it Iterator[U]
}

// Merge returns an iterator that merges the elements of the given iterators using the given function.
func Merge[T, S, U any](fn func(T, S) U, it1 Iterator[T], it2 Iterator[S]) *MergeIterator[U] {
	zi := Zip[T, S](it1, it2)
	mi := Map[Pair[T, S], U](zi, func(t Pair[T, S]) U { return fn(t.First, t.Second) })
	return &MergeIterator[U]{it: mi}
}

// Next returns true if the iterator has more elements.
func (mi *MergeIterator[U]) Next() bool {
	return mi.it.Next()
}

// Value returns the current element in the iterator.
func (mi *MergeIterator[U]) Value() U {
	return mi.it.Value()
}

// Collect return a slice containing all elements in the iterator.
func (mi *MergeIterator[U]) Collect() []U {
	return CollectFromIter[U](mi)
}
