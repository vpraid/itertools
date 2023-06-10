package iterator

// Imbuer is an interface that lets user imbue the implementer with another underlying iterator.
type Imbuer[T any] interface {
	Iterator[T]
	// Imbue replaces the underlying iterator with the given one.
	Imbue(it Iterator[T])
}
