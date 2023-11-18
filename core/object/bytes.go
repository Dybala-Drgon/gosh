package object

import (
	"bytes"
)

// Bytes represents a byte array.
type Bytes struct {
	Value []byte
}

func (o *Bytes) String() string {
	return string(o.Value)
}

// TypeName returns the name of the type.
func (o *Bytes) TypeName() string {
	return "bytes"
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Bytes) IsFalsy() bool {
	return len(o.Value) == 0
}

// Equals returns true if the value of the type
// is equal to the value of another object.
func (o *Bytes) Equals(x Object) bool {
	t, ok := x.(*Bytes)
	if !ok {
		return false
	}

	return bytes.Compare(o.Value, t.Value) == 0
}
