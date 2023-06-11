package aggregate

import "github.com/vpraid/itertools/iterator"

// Fold applies a function to each element of the iterator, threading an accumulator argument through the computation.
func Fold[T, U any](it iterator.Iterator[T], initial U, fn func(U, T) U) U {
	var result = initial
	for it.Next() {
		result = fn(result, it.Value())
	}
	return result
}

// Reduce applies a function to each element of the iterator, threading an accumulator argument through the computation.
// The first element of the iterator is used as the initial value of the accumulator.
// If the iterator is empty, Reduce panics.
func Reduce[T any](it iterator.Iterator[T], fn func(T, T) T) T {
	if !it.Next() {
		panic("Reduce called on empty iterator")
	}
	initial := it.Value()
	return Fold(it, initial, fn)
}
