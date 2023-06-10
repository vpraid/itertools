package iterator

// Collector is the interface that provides a way to collect the elements of an iterator into a slice.
type Collector[T any] interface {
	Iterator[T]
	// Collect returns the elements of the underlying iterator as a slice.
	Collect() []T
}

func CollectFromIter[T any](it Iterator[T]) []T {
	var result []T
	for it.Next() {
		result = append(result, it.Value())
	}
	return result
}
