package aggregate

import (
	"github.com/vpraid/itertools/functional"
	"github.com/vpraid/itertools/iterator"
	"golang.org/x/exp/constraints"
)

// Number represents any number type that can be aggregated.
type Number interface {
	constraints.Integer | constraints.Float
}

// Sum returns the sum of all elements in the iterator.
func Sum[T Number](it iterator.Iterator[T]) T {
	return functional.Fold[T, T](it, 0, func(acc, x T) T { return acc + x })
}

// Product returns the product of all elements in the iterator.
func Product[T Number](it iterator.Iterator[T]) T {
	return functional.Fold[T, T](it, 1, func(acc, x T) T { return acc * x })
}
