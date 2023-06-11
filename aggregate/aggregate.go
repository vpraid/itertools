package aggregate

import (
	"github.com/vpraid/itertools/iterator"
	"golang.org/x/exp/constraints"
)

// Number represents any number type that can be aggregated.
type Number interface {
	constraints.Integer | constraints.Float
}

// Sum returns the sum of all elements in the iterator.
func Sum[T Number](it iterator.Iterator[T]) T {
	var sum T
	for it.Next() {
		sum += it.Value()
	}
	return sum
}

// Product returns the product of all elements in the iterator.
func Product[T Number](it iterator.Iterator[T]) T {
	var product T = 1
	for it.Next() {
		product *= it.Value()
	}
	return product
}
