package object

import (
	"errors"
	"github.com/gookit/slog"
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
			// TODO:
		case token.OpLess:
			lhs := o.Value
			rhs := rhs.Value
			if lhs < rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OpLessEqual:
			lhs := o.Value
			rhs := rhs.Value
			if lhs <= rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OPGreater:
			lhs := o.Value
			rhs := rhs.Value
			if lhs > rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OPGreaterEqual:
			lhs := o.Value
			rhs := rhs.Value
			if lhs >= rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OpEqual:
			lhs := o.Value
			rhs := rhs.Value
			if lhs == rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OpNotEqual:
			lhs := o.Value
			rhs := rhs.Value
			if lhs != rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		default:
			slog.Debug("尚未支持")

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
			// TODO:
		case token.OpLess:
			lhs := float64(o.Value)
			rhs := rhs.Value
			if lhs < rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OpLessEqual:
			lhs := float64(o.Value)
			rhs := rhs.Value
			if lhs <= rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OPGreater:
			lhs := float64(o.Value)
			rhs := rhs.Value
			if lhs > rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OPGreaterEqual:
			lhs := float64(o.Value)
			rhs := rhs.Value
			if lhs >= rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OpEqual:
			lhs := float64(o.Value)
			rhs := rhs.Value
			if lhs == rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OpNotEqual:
			lhs := float64(o.Value)
			rhs := rhs.Value
			if lhs != rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		default:
			slog.Debug("尚未支持")

		}
	}

	return nil, errors.New("二元操作符类型错误")
}
