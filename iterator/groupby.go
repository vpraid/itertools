package iterator

// GroupByIterator is an iterator that groups consecutive elements of the input iterator into iterable groups.
type GroupByIterator[T any, K comparable] struct {
	it    *PeekAheadIterator[T]
	fn    func(T) K
	group *Group[T, K]
}

// Group is an iterable group of consecutive elements having the same key.
type Group[T any, K comparable] struct {
	Key   K
	it    *PeekAheadIterator[T]
	fn    func(T) K
	fused bool
}

// GroupBy returns an GroupByIterator that groups consecutive elements of the input iterator into iterable groups
// based on the provided mapping function.
func GroupBy[T any, K comparable](it Iterator[T], fn func(t T) K) *GroupByIterator[T, K] {
	return &GroupByIterator[T, K]{
		it: PeekAhead(it),
		fn: fn,
	}
}

// Next skips all elements within the current group until the next group is found, then returns true if there is
// such a group, false if the underlying iterator was exhausted.
func (gi *GroupByIterator[T, K]) Next() bool {
	// Skip all elements in the current group until we find the next group.
	for !gi.it.Exhausted() {
		// Peek ahead to see if there is a next element or if it has a different key.
		// If yes, start a new group.
		value := gi.it.Peek()
		key := gi.fn(value)
		if gi.group == nil || key != gi.group.Key {
			gi.group = &Group[T, K]{
				Key:   key,
				it:    gi.it,
				fn:    gi.fn,
				fused: false,
			}
			return true
		}
		if !gi.it.Next() {
			break
		}
	}
	gi.group = nil
	return false
}

// Value returns the current group.
func (gi *GroupByIterator[T, K]) Value() *Group[T, K] {
	return gi.group
}

// Next advances the iterator within the group to the next element. When the group is exhausted, it fuses the
// group iterator. This means that the group iterator will no longer be able to read from the underlying iterator.
func (g *Group[T, K]) Next() bool {
	if g.fused {
		return false
	}
	if g.it.Exhausted() || g.fn(g.it.Peek()) != g.Key {
		g.fused = true
		return false
	}
	return g.it.Next()
}

// Value returns the current element.
func (g *Group[T, K]) Value() T {
	return g.it.Value()
}
