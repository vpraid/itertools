package common

// Zero returns the zero value of the given type.
func Zero[T any]() T {
	var zero T
	return zero
}
