package generator

import (
	"github.com/vpraid/itertools/iterator"
	"golang.org/x/exp/constraints"
)

// RangeIterator is an iterator that produces a sequence of integers.
type RangeIterator[T constraints.Integer] struct {
	start T
	end   T
	step  T
	value T
}

// Range returns a RangeIterator for the given start and end values and a step of 1.
func Range[T constraints.Integer](start, end T) *RangeIterator[T] {
	return &RangeIterator[T]{
		start: start,
		end:   end,
		step:  1,
	}
}

// RangeBy returns a RangeIterator for the given start, end and step values.
func RangeBy[T constraints.Integer](start, end, step T) *RangeIterator[T] {
	return &RangeIterator[T]{
		start: start,
		end:   end,
		step:  step,
	}
}

// Next advances the iterator to the next value of the underlying slice. It returns false if the range was exhausted.
func (ri *RangeIterator[T]) Next() bool {
	if ri.step > 0 {
		if ri.start >= ri.end {
			ri.value = 0
			return false
		}
		ri.value = ri.start
		ri.start += ri.step
		return true
	}
	if ri.start <= ri.end {
		ri.value = 0
		return false
	}
	ri.value = ri.start
	ri.start += ri.step
	return true
}

// Value returns the current value of the range pointed by the iterator. If the range was exhausted, returned value is unspecified.
func (ri *RangeIterator[T]) Value() T {
	return ri.value
}

// Collect returns all the values of the range as a slice.
func (ri *RangeIterator[T]) Collect() []T {
	return iterator.CollectFromIter[T](ri)
}
