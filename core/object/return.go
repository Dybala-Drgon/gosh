package object

import (
	"errors"
	"gosh/compiler/token"
)

// ReturnValue represents a value that is being returned.
type ReturnValue struct {
	Value []Object
}

// TypeName returns the name of the type.
func (o *ReturnValue) TypeName() string {
	return "return-value"
}

func (o *ReturnValue) String() string {
	return "<return-value>"
}

// IsFalsy returns true if the value of the type is falsy.
func (o *ReturnValue) IsFalsy() bool {
	return false
}

// Equals returns true if the value of the type
// is equal to the value of another object.
func (o *ReturnValue) Equals(x Object) bool {
	return false
}

func (o *ReturnValue) BinaryOp(op token.Opcode, rhs Object) (Object, error) {
	return nil, errors.New("invalid operator")
}
