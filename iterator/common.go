package iterator

// Zero returns the zero value of the given type.
func zero[T any]() T {
	var zero T
	return zero
}
