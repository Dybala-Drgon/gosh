package bytecode

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/gookit/slog"
	"gosh/compiler/parser"
	"gosh/compiler/token"
	"reflect"
)

func (v *GoshVisitor) VisitASSIGN(ctx *parser.ASSIGNContext) interface{} {
	slog.Trace("enter assign stmt", ctx.GetText())
	asg := ctx.Assignment()
	var idt []string
	for _, id := range asg.Lvalue().AllID() {
		v.SymbolTables[v.CurSymTableIdx].Define(id.GetText())
		//slog.Trace(id.GetText())
		idt = append(idt, id.GetText())
	}
	lSize := len(asg.Lvalue().AllID())
	resTmp := v.visitRule(asg.Rvalue())
	rSize, ok := resTmp.(int)
	if !ok {
		slog.Error("visit Rvalue return error type")
	}
	if lSize != rSize {
		slog.Error("assign 左右参数不匹配, left size ", lSize, " right size ", rSize)
		panic("assign 左右参数不匹配")
	}
	// 符号表id、在该符号表下的idx
	for i := len(idt) - 1; i >= 0; i-- {
		v.emit(token.OpSetVar, v.CurSymTableIdx, v.SymbolTables[v.CurSymTableIdx].store[idt[i]])
	}
	return nil
}

func (v *GoshVisitor) VisitLvalue(ctx *parser.LvalueContext) interface{} {
	slog.Trace("enter lvalue =", ctx.GetText())
	return nil
}

func (v *GoshVisitor) VisitRvalue(ctx *parser.RvalueContext) interface{} {
	slog.Trace("rvalue = ", ctx.GetText())
	var resSize int
	resSize = 0
	for _, elem := range ctx.GetChildren() {
		var tmpReturn interface{}
		switch t := elem.(type) {
		case *parser.ConstvalueContext:
			t.Accept(v)
			resSize++
		case *parser.ExpressionContext:
			t.Accept(v)
			resSize++
		case *antlr.TerminalNodeImpl:
			if t.GetText() != "," {
				slog.Error("符号错误")
			}
			continue
		case *parser.RvalueContext:
			tmpReturn = t.Accept(v)
		case *parser.FunctionCallContext:
			// TODO: func调用
			//tmpReturn = t.Accept(v)
		default:
			slog.Trace("type:", t)
			slog.Trace("type:", reflect.TypeOf(elem))
		}
		if length, ok := tmpReturn.(int); ok {
			resSize += length
		}
	}
	return resSize
}
