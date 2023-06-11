package iterator

type ChunkIterator[T any] struct {
	it    *PeekAheadIterator[T]
	size  int
	chunk *Chunk[T]
}

// Chunk returns an iterator that groups consecutive elements of the input iterator into iterable chunks of the given size.
func Chunks[T any](it Iterator[T], size int) *ChunkIterator[T] {
	if it == nil {
		return &ChunkIterator[T]{
			it:   nil,
			size: size,
		}
	}
	return &ChunkIterator[T]{
		it:   PeekAhead[T](it),
		size: size,
	}
}

// Next skips all elements within the current chunk until the next chunk is found, then returns true if there is
// such a chunk, false if the underlying iterator was exhausted.
func (ci *ChunkIterator[T]) Next() bool {
	if ci.it == nil {
		return false
	}
	// If we don't yet have a chunk, create one.
	if ci.chunk == nil && !ci.it.Exhausted() {
		ci.chunk = &Chunk[T]{
			it:   ci.it,
			size: ci.size,
		}
		return true
	}
	// Skip all elements in the current chunk until we find the next chunk.
	for ci.chunk.Next() {
		// Skip
	}
	// We found the next chunk, so we can return true.
	if !ci.it.Exhausted() {
		ci.chunk = &Chunk[T]{
			it:   ci.it,
			size: ci.size,
		}
		return true
	}
	// We exhausted the underlying iterator, so we can return false.
	ci.chunk = nil
	return false
}

// Value returns the current chunk.
func (ci *ChunkIterator[T]) Value() *Chunk[T] {
	return ci.chunk
}

// Chunk represents a group of consecutive elements of a ChunkIterator of a given size.
type Chunk[T any] struct {
	it   Iterator[T]
	size int
}

// Next advances the iterator and returns true if there is a next element, false otherwise.
func (c *Chunk[T]) Next() bool {
	if c.size == 0 || !c.it.Next() {
		return false
	}
	c.size--
	return true
}

// Value returns the current element. If the chunk iterator is exhausted, the value is undefined.
func (c *Chunk[T]) Value() T {
	return c.it.Value()
}

// Collect returns all elements of the chunk as a slice.
func (c *Chunk[T]) Collect() []T {
	return CollectFromIter[T](c)
}
