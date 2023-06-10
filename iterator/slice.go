package iterator

// SliceIterator is an iterator that iterates over a given slice.
type SliceIterator[T any] struct {
	elements []T
	value    T
}

// Slice returns a SliceIterator for the given slice.
func Slice[T any](elements []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		elements: elements,
	}
}

// Next advances the iterator to the next element of the underlying slice. It returns false if
// the underlying iterator was exhausted.
func (si *SliceIterator[T]) Next() bool {
	if len(si.elements) == 0 {
		si.value = zero[T]()
		return false
	}
	si.value = si.elements[0]
	si.elements = si.elements[1:]
	return true
}

// Value returns the current element of the slice pointed by the iterator. If the iterator was exhausted, it returns
// the zero value.
func (si *SliceIterator[T]) Value() T {
	return si.value
}
