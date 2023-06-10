// Package iterator provides a generic iterator interface and methods to manipulate it.
package iterator

// Iterator is the interface that provides a way to iterate over a collection of elements of type T.
type Iterator[T any] interface {
	// Next advances the iterator to the next element. It returns false if the underlying iterator was exhausted.
	Next() bool
	// Value returns the current element of the underlying iterator. If the underlying iterator was exhausted, it returns the zero value.
	Value() T
}

// Collector is the interface that provides a way to collect the elements of an iterator into a slice.
type Collector[T any] interface {
	Iterator[T]
	// Collect returns the elements of the underlying iterator as a slice.
	Collect() []T
}
