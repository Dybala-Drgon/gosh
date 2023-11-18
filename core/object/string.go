package object

import (
	"strconv"
)

// String represents a string value.
type String struct {
	Value   string
	runeStr []rune
}

// TypeName returns the name of the type.
func (o *String) TypeName() string {
	return "string"
}

func (o *String) String() string {
	return strconv.Quote(o.Value)
}

// IsFalsy returns true if the value of the type is falsy.
func (o *String) IsFalsy() bool {
	return len(o.Value) == 0
}

// Equals returns true if the value of the type
// is equal to the value of another object.
func (o *String) Equals(x Object) bool {
	t, ok := x.(*String)
	if !ok {
		return false
	}

	return o.Value == t.Value
}
