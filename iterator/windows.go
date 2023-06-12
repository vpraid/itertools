package iterator

import "github.com/vpraid/itertools/internal/common"

type WindowIterator[T any] struct {
	it     Iterator[T]
	n      int
	window *Window[T]
}

func Windows[T any](it Iterator[T], n int) *WindowIterator[T] {
	if n <= 0 {
		panic("n must be greater than 0")
	}
	return &WindowIterator[T]{
		it: it,
		n:  n,
		window: &Window[T]{
			elements: make([]T, 0, n),
			value:    common.Zero[T](),
		},
	}
}

// Next advances the iterator and returns true if there unless there are no more elements to read from the underlying iterator.
func (wi *WindowIterator[T]) Next() bool {
	if wi.it == nil {
		return false
	}
	if len(wi.window.elements) == wi.n {
		wi.window = makeWindow(wi.window)
	}
	for len(wi.window.elements) < wi.n && wi.it.Next() {
		wi.window.pushBack(wi.it.Value())
	}
	if len(wi.window.elements) < wi.n {
		wi.window = nil
		return false
	}
	return true
}

// Value returns the current element's value.
func (wi *WindowIterator[T]) Value() *Window[T] {
	return wi.window
}

// Collect returns all elements as a slice.
func (wi *WindowIterator[T]) Collect() []*Window[T] {
	return CollectFromIter[*Window[T]](wi)
}

type Window[T any] struct {
	elements []T
	value    T
}

func (w *Window[T]) pushBack(e T) {
	w.elements = append(w.elements, e)
}

func makeWindow[T any](other *Window[T]) *Window[T] {
	window := &Window[T]{
		elements: make([]T, len(other.elements)-1, len(other.elements)),
		value:    common.Zero[T](),
	}
	copy(window.elements, other.elements[1:])
	return window
}

// Next advances the iterator and returns true if there unless there are no more elements to read from the underlying slice.
func (w *Window[T]) Next() bool {
	if len(w.elements) == 0 {
		w.value = common.Zero[T]()
		return false
	}
	w.value = w.elements[0]
	w.elements = w.elements[1:]
	return true
}

// Value returns the current element's value.
func (w *Window[T]) Value() T {
	return w.value
}

// Collect returns all elements as a slice.
func (w *Window[T]) Collect() []T {
	return w.elements
}
