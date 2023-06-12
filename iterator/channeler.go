package iterator

type Channeler[T any] interface {
	Iterator[T]
	// Chan returns a channel that will receive the elements of the underlying iterator.
	Chan() <-chan T
}

func ChanFromIter[T any](it Iterator[T]) <-chan T {
	channel := make(chan T)
	go func() {
		defer close(channel)
		for it.Next() {
			channel <- it.Value()
		}
	}()
	return channel
}
