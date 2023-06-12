package iterator

type StepByIterator[T any] struct {
	it      Iterator[T]
	step    int
	atStart bool
}

// StepBy returns an iterator that steps through the elements of the input iterator by the given step.
func StepBy[T any](it Iterator[T], step int) *StepByIterator[T] {
	return &StepByIterator[T]{
		it:      it,
		step:    step,
		atStart: true,
	}
}

// Next advances the iterator to the next element.
func (sbi *StepByIterator[T]) Next() bool {
	if sbi.it == nil {
		return false
	}
	if sbi.atStart {
		sbi.atStart = false
		return sbi.it.Next()
	}
	for i := 0; i < sbi.step; i++ {
		if !sbi.it.Next() {
			return false
		}
	}
	return true
}

// Value returns the current element.
func (sbi *StepByIterator[T]) Value() T {
	return sbi.it.Value()
}

// Collect returns all remaining elements in the iterator.
func (sbi *StepByIterator[T]) Collect() []T {
	return CollectFromIter[T](sbi)
}

// Chan returns a channel that will receive the remaining elements.
func (sbi *StepByIterator[T]) Chan() <-chan T {
	return ChanFromIter[T](sbi)
}
