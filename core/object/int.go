package object

import (
	"errors"
	"gosh/compiler/token"

	"strconv"
)

// Int represents an integer value.
type Int struct {
	Value int64
}

func (o *Int) String() string {
	return strconv.FormatInt(o.Value, 10)
}

// TypeName returns the name of the type.
func (o *Int) TypeName() string {
	return "int"
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Int) IsFalsy() bool {
	return o.Value == 0
}

// Equals returns true if the value of the type
// is equal to the value of another object.
func (o *Int) Equals(x Object) bool {
	t, ok := x.(*Int)
	if !ok {
		return false
	}

	return o.Value == t.Value
}

func (o *Int) BinaryOp(op token.Opcode, rhs Object) (Object, error) {
	switch rhs := rhs.(type) {
	case *Int:
		switch op {
		case token.OpAdd:
			r := o.Value + rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.OpSub:
			r := o.Value - rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.OpMul:
			r := o.Value * rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil
		case token.OpDiv:
			r := o.Value / rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Int{Value: r}, nil

		}
	case *Float:
		switch op {
		case token.OpAdd:
			return &Float{float64(o.Value) + rhs.Value}, nil
		case token.OpSub:
			return &Float{float64(o.Value) - rhs.Value}, nil
		case token.OpMul:
			return &Float{float64(o.Value) * rhs.Value}, nil
		case token.OpDiv:
			return &Float{float64(o.Value) / rhs.Value}, nil
		}
	}

	return nil, errors.New("二元操作符类型错误")
}
