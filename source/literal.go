package source

// LiteralIterator is an iterator that iterates through the values of a slice given as a variadic argument list.
type LiteralIterator[T any] struct {
	*SliceIterator[T]
}

// Literal returns a LiteralIterator for the given argument list.
func Literal[T any](args ...T) *LiteralIterator[T] {
	return &LiteralIterator[T]{
		SliceIterator: Slice(args),
	}
}

// Next advances the iterator to the next value of the underlying slice. It returns false if the slice was exhausted.
func (li *LiteralIterator[T]) Next() bool {
	return li.SliceIterator.Next()
}

// Value returns the current value of the slice pointed by the iterator. If the slice was exhausted, it returns
// the zero value.
func (li *LiteralIterator[T]) Value() T {
	return li.SliceIterator.Value()
}

// Collect returns the values of the slice collected into a slice.
func (li *LiteralIterator[T]) Collect() []T {
	return li.SliceIterator.Collect()
}

// Chan returns a channel that yields the elements of the underlying iterator.
func (li *LiteralIterator[T]) Chan() <-chan T {
	return li.SliceIterator.Chan()
}
