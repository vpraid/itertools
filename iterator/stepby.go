package iterator

import "github.com/vpraid/itertools/internal/common"

type StepByIterator[T any] struct {
	it    Iterator[T]
	step  int
	value T
}

// StepBy returns an iterator that steps through the elements of the input iterator by the given step.
func StepBy[T any](it Iterator[T], step int) *StepByIterator[T] {
	return &StepByIterator[T]{
		it:    it,
		step:  step,
		value: common.Zero[T](),
	}
}

// Next advances the iterator to the next element.
func (sbi *StepByIterator[T]) Next() bool {
	if !sbi.it.Next() {
		return false
	}
	sbi.value = sbi.it.Value()
	for i := 0; i < sbi.step-1; i++ {
		if !sbi.it.Next() {
			return false
		}
	}
	return true
}

// Value returns the current element.
func (sbi *StepByIterator[T]) Value() T {
	return sbi.value
}

// Collect returns all remaining elements in the iterator.
func (sbi *StepByIterator[T]) Collect() []T {
	return CollectFromIter[T](sbi)
}
