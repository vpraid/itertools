package source

// ChannelIterator is an iterator that iterates through the values of a channel.
type ChannelIterator[T any] struct {
	source <-chan T
	value  T
}

// Channel returns a ChannelIterator for the given channel.
func Channel[T any](ch <-chan T) *ChannelIterator[T] {
	return &ChannelIterator[T]{
		source: ch,
	}
}

// Next advances the iterator to the next value of the underlying channel. It returns false if the channel was closed.
func (ci *ChannelIterator[T]) Next() bool {
	var ok bool
	ci.value, ok = <-ci.source
	return ok
}

// Value returns the current value of the channel pointed by the iterator. If the channel was closed, it returns
// the zero value.
func (ci *ChannelIterator[T]) Value() T {
	return ci.value
}

// Collect returns the values of the channel collected into a slice.
func (ci *ChannelIterator[T]) Collect() []T {
	var result []T
	for ci.Next() {
		result = append(result, ci.Value())
	}
	return result
}
