package functional

import "github.com/vpraid/itertools/iterator"

func Any[T any](source iterator.Iterator[T], predicate func(T) bool) bool {
	for source.Next() {
		if predicate(source.Value()) {
			return true
		}
	}
	return false
}

func All[T any](source iterator.Iterator[T], predicate func(T) bool) bool {
	for source.Next() {
		if !predicate(source.Value()) {
			return false
		}
	}
	return true
}
