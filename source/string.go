package source

import (
	"io"
	"strings"
)

// StringIterator is an iterator that iterates through the runes of a string.
type StringIterator struct {
	source io.RuneReader
	value  rune
}

// String returns a StringIterator for the given string.
func String(s string) *StringIterator {
	return &StringIterator{
		source: strings.NewReader(s),
	}
}

// Next advances the iterator to the next rune of the underlying string. It returns false if there was an error reading
// the next rune.
func (si *StringIterator) Next() bool {
	r, _, err := si.source.ReadRune()
	if err != nil {
		si.value = 0
		return false
	}
	si.value = r
	return true
}

// Value returns the current rune of the string pointed by the iterator. If the iterator was exhausted, it returns
// the zero value.
func (si *StringIterator) Value() rune {
	return si.value
}

// Collect returns the rune of the string collected into a slice.
func (si *StringIterator) Collect() []rune {
	var result []rune
	for si.Next() {
		result = append(result, si.Value())
	}
	return result
}
