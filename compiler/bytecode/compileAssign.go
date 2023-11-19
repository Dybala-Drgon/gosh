package bytecode

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/gookit/slog"
	"gosh/compiler/parser"
	"reflect"
)

func (v *GoshVisitor) VisitASSIGN(ctx *parser.ASSIGNContext) interface{} {
	slog.Trace("enter assign stmt", ctx.GetText())
	asg := ctx.Assignment()
	table := v.SymbolTables[v.CurSymTableIdx]
	for _, id := range asg.Lvalue().AllID() {
		table.Define(id.GetText())
		//slog.Trace(id.GetText())
	}
	v.visitRule(asg.Rvalue())
	return nil
}

func (v *GoshVisitor) VisitLvalue(ctx *parser.LvalueContext) interface{} {
	slog.Trace("enter lvalue =", ctx.GetText())
	return nil
}

func (v *GoshVisitor) VisitRvalue(ctx *parser.RvalueContext) interface{} {
	slog.Trace("rvalue = ", ctx.GetText())

	for _, elem := range ctx.GetChildren() {
		switch t := elem.(type) {
		case *parser.ConstvalueContext:
			t.Accept(v)
		case *parser.ExpressionContext:
			t.Accept(v)
		case *antlr.TerminalNodeImpl:
			if t.GetText() != "," {
				slog.Error("符号错误")
			}
			continue
		case *parser.RvalueContext:
			t.Accept(v)
		default:
			slog.Trace("type:", t)
			slog.Trace("type:", reflect.TypeOf(elem))
		}
	}
	return nil
}
