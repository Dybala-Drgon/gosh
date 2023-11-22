package object

import (
	"errors"
	"github.com/gookit/slog"
	"gosh/compiler/token"
	"math"
	"strconv"
)

// Float represents a floating point number value.
type Float struct {
	Value float64
}

func (o *Float) String() string {
	return strconv.FormatFloat(o.Value, 'f', -1, 64)
}

// TypeName returns the name of the type.
func (o *Float) TypeName() string {
	return "float"
}

// IsFalsy returns true if the value of the type is falsy.
func (o *Float) IsFalsy() bool {
	return math.IsNaN(o.Value)
}

// Equals returns true if the value of the type
// is equal to the value of another object.
func (o *Float) Equals(x Object) bool {
	t, ok := x.(*Float)
	if !ok {
		return false
	}

	return o.Value == t.Value
}

func (o *Float) BinaryOp(op token.Opcode, rhs Object) (Object, error) {
	switch rhs := rhs.(type) {
	case *Float:
		switch op {
		case token.OpAdd:
			r := o.Value + rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.OpSub:
			r := o.Value - rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.OpMul:
			r := o.Value * rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.OpDiv:
			r := o.Value / rhs.Value
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
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
		case token.OpBOr:
			lhs := o.Value
			rhs := rhs.Value
			if lhs != 0 || rhs != 0 {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OpBAnd:
			lhs := o.Value
			rhs := rhs.Value
			if lhs != 0 && rhs != 0 {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		default:
			slog.Debug("尚未支持")
		}
	case *Int:
		switch op {
		case token.OpAdd:
			r := o.Value + float64(rhs.Value)
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.OpSub:
			r := o.Value - float64(rhs.Value)
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.OpMul:
			r := o.Value * float64(rhs.Value)
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil
		case token.OpDiv:
			r := o.Value / float64(rhs.Value)
			if r == o.Value {
				return o, nil
			}
			return &Float{Value: r}, nil

		case token.OpLess:
			lhs := o.Value
			rhs := float64(rhs.Value)
			if lhs < rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OpLessEqual:
			lhs := o.Value
			rhs := float64(rhs.Value)
			if lhs <= rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OPGreater:
			lhs := o.Value
			rhs := float64(rhs.Value)
			if lhs > rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OPGreaterEqual:
			lhs := o.Value
			rhs := float64(rhs.Value)
			if lhs >= rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OpEqual:
			lhs := o.Value
			rhs := float64(rhs.Value)
			if lhs == rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OpNotEqual:
			lhs := o.Value
			rhs := float64(rhs.Value)
			if lhs != rhs {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OpBOr:
			lhs := o.Value
			rhs := float64(rhs.Value)
			if lhs != 0 || rhs != 0 {
				return &Int{Value: 1}, nil
			} else {
				return &Int{Value: 0}, nil
			}
		case token.OpBAnd:
			lhs := o.Value
			rhs := float64(rhs.Value)
			if lhs != 0 && rhs != 0 {
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
