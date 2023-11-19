package object

import (
	"gosh/compiler/token"
)

type Object interface {
	// TypeName should return the name of the type.
	TypeName() string

	BinaryOp(op token.Opcode, rhs Object) (Object, error)

	// String should return a string representation of the type's value.
	String() string

	// IsFalsy should return true if the value of the type
	// should be considered as falsy.
	IsFalsy() bool

	// Equals should return true if the value of the type
	// should be considered as equal to the value of another object.
	Equals(another Object) bool
}
