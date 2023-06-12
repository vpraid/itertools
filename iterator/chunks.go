package iterator

import "github.com/vpraid/itertools/source"

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
	// Skip all remaining elements in the current chunk until we find the next chunk.
	// If the chunk is detached, we don't need to skip the elements, since the iterator is
	// already pointing to the beginning of the next chunk.
	if !ci.chunk.detached {
		for size := ci.chunk.size; size != 0; size-- {
			if !ci.it.Next() {
				break
			}
		}
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
	it       Iterator[T]
	size     int
	detached bool
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

// Chan returns a channel that will receive the elements of the chunk.
func (c *Chunk[T]) Chan() <-chan T {
	return ChanFromIter[T](c)
}

// Detach separates the chunk from the underlying iterator. The elements of the chunk will be placed in a newly allocated
// slice, and the chunk iterator will no longer read from the underlying iterator but instead will start reading
// from the beginning of the newly allocated slice. Addtionally, detaching the chunk will advance the underlying iterator
// to the end of the chunk.
func (c *Chunk[T]) Detach() *Chunk[T] {
	remainingSize := c.size
	c.it = PeekAhead[T](source.Slice[T](c.Collect()))
	c.size = remainingSize
	c.detached = true
	return c
}
