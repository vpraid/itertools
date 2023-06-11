package functional

import (
	"github.com/vpraid/itertools/iterator"
	"github.com/vpraid/itertools/partial"
)

// Compose2 creates a chain of iterators by imbuing the second iterator with the first one.
// The second iterator in the chain is returned.
func Compose2[
	S, T1 any,
	I1 partial.Function[S, T1],
](source iterator.Iterator[S], it1 I1) I1 {
	it1.Bind(source)
	return it1
}

// Compose3 creates a chain of iterators by imbuing the second iterator with the first one, then the third one with the second one.
// The last iterator in the chain is returned.
func Compose3[
	S, T1, T2 any,
	I1 partial.Function[S, T1],
	I2 partial.Function[T1, T2],
](source iterator.Iterator[S], it1 I1, it2 I2) I2 {
	it1.Bind(source)
	it2.Bind(it1)
	return it2
}

// Compose4 creates a chain of iterators by imbuing the second iterator with the first, the third with the second, and so on.
// The last iterator in the chain is returned.
func Compose4[
	S, T1, T2, T3 any,
	I1 partial.Function[S, T1],
	I2 partial.Function[T1, T2],
	I3 partial.Function[T2, T3],
](source iterator.Iterator[S], it1 I1, it2 I2, it3 I3) I3 {
	it1.Bind(source)
	it2.Bind(it1)
	it3.Bind(it2)
	return it3
}

// Compose5 creates a chain of iterators by imbuing the second iterator with the first, the third with the second, and so on.
// The last iterator in the chain is returned.
func Compose5[
	S, T1, T2, T3, T4 any,
	I1 partial.Function[S, T1],
	I2 partial.Function[T1, T2],
	I3 partial.Function[T2, T3],
	I4 partial.Function[T3, T4],
](source iterator.Iterator[S], it1 I1, it2 I2, it3 I3, it4 I4) I4 {
	it1.Bind(source)
	it2.Bind(it1)
	it3.Bind(it2)
	it4.Bind(it3)
	return it4
}

// Compose6 creates a chain of iterators by imbuing the second iterator with the first, the third with the second, and so on.
// The last iterator in the chain is returned.
func Compose6[
	S, T1, T2, T3, T4, T5 any,
	I1 partial.Function[S, T1],
	I2 partial.Function[T1, T2],
	I3 partial.Function[T2, T3],
	I4 partial.Function[T3, T4],
	I5 partial.Function[T4, T5],
](source iterator.Iterator[S], it1 I1, it2 I2, it3 I3, it4 I4, it5 I5) I5 {
	it1.Bind(source)
	it2.Bind(it1)
	it3.Bind(it2)
	it4.Bind(it3)
	it5.Bind(it4)
	return it5
}

// Compose7 creates a chain of iterators by imbuing the second iterator with the first, the third with the second, and so on.
// The last iterator in the chain is returned.
func Compose7[
	S, T1, T2, T3, T4, T5, T6 any,
	I1 partial.Function[S, T1],
	I2 partial.Function[T1, T2],
	I3 partial.Function[T2, T3],
	I4 partial.Function[T3, T4],
	I5 partial.Function[T4, T5],
	I6 partial.Function[T5, T6],
](source iterator.Iterator[S], it1 I1, it2 I2, it3 I3, it4 I4, it5 I5, it6 I6) I6 {
	it1.Bind(source)
	it2.Bind(it1)
	it3.Bind(it2)
	it4.Bind(it3)
	it5.Bind(it4)
	it6.Bind(it5)
	return it6
}
