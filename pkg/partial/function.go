package partial

import "github.com/vpraid/itertools/pkg/iterator"

// Function is an interface that lets user imbue the implementer with another underlying iterator.
type Function[From, To any] interface {
	iterator.Iterator[To]
	// Bind replaces the underlying iterator with the given one.
	Bind(it iterator.Iterator[From])
}
