package iterator

type Channeler[T any] interface {
	Iterator[T]
	// Chan returns a channel that will receive the elements of the underlying iterator.
	Chan() <-chan T
}
