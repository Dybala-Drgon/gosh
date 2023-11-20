package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/gookit/slog"
	"gosh/compiler/bytecode"
	"gosh/compiler/parser"
	"gosh/core/vm"
)

type goshErrorListener struct {
	*antlr.DefaultErrorListener
}

func (e *goshErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, ex antlr.RecognitionException) {
	str := fmt.Sprintf("SyntaxError %d:%d  [%s]\n", line, column, msg)
	slog.Error(str)
	panic(str)
}

func RunGosh(input string) {
	is := antlr.NewInputStream(input)
	lexer := parser.NewGoshLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewGoshParser(stream)
	p.AddErrorListener(&goshErrorListener{})
	visitor := &bytecode.GoshVisitor{}
	visitor.InitCompiler()
	s := p.Program()

	//slog.Info(s.GetText())
	code := s.Accept(visitor)
	//slog.Trace(reflect.TypeOf(code))
	virtualMachine := vm.NewVM(code.(*bytecode.Bytecode))
	virtualMachine.Run()
}

func main() {
	slog.SetLogLevel(slog.TraceLevel)
	slog.Configure(func(logger *slog.SugaredLogger) {
		f := logger.Formatter.(*slog.TextFormatter)
		f.EnableColor = true

	})
	RunGosh(`
t4=999
c,b,a,t,w2=2.2,"str",1+2*5,t4,-8
`)
	//RunGosh(`c,b,a,t,w=2.2,"str",1+2*5,8-(5+1),-8`)
	//RunGosh(`c,b,a,t,w=2.2,"str",1+2*5,-8,fun()fun`)
	//RunGosh(`7+8*9+(8-4)*2`)
	//RunGosh(`-8`)
}

//
//func main() {
//	is := antlr.NewInputStream("1 +     2 * 3+1+1")
//	// Create the Lexer
//	lexer := parser.NewCalcLexer(is)
//	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
//
//	// Create the Parser
//	p := parser.NewCalcParser(stream)
//
//	visitor := CalVisitor{}
//	// Finally parse the expression
//	p.Start_().Accept(&visitor)
//
//	slog.Trace(visitor.Pop())
//}
//
//type CalVisitor struct {
//	parser.BaseCalcVisitor
//	Stack []int
//}
//
//func (l *CalVisitor) push(i int) {
//	l.Stack = append(l.Stack, i)
//}
//
//func (l *CalVisitor) Pop() int {
//	if len(l.Stack) < 1 {
//		panic("Stack is empty unable to Pop")
//	}
//	// Get the last value from the Stack.
//	result := l.Stack[len(l.Stack)-1]
//
//	// Remove the last element from the Stack.
//	l.Stack = l.Stack[:len(l.Stack)-1]
//
//	return result
//}
//
//func (v *CalVisitor) visitRule(node antlr.RuleNode) interface{} {
//	node.Accept(v)
//	return nil
//}
//
//func (v *CalVisitor) VisitStart(ctx *parser.StartContext) interface{} {
//	slog.Trace("VisitStart")
//	return v.visitRule(ctx.Expr())
//}
//
//func (v *CalVisitor) VisitNumber(ctx *parser.NumberContext) interface{} {
//	slog.Trace("VisitNumber")
//	i, err := strconv.Atoi(ctx.NUMBER().GetText())
//	if err != nil {
//		panic(err.Error())
//	}
//
//	v.push(i)
//	return nil
//}
//
//func (v *CalVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
//	slog.Trace("VisitMulDiv")
//	//push expression result to Stack
//	v.visitRule(ctx.Expr(0))
//	v.visitRule(ctx.Expr(1))
//	slog.Trace("kids size = ", ctx.GetChildCount(), ctx.Expr(2))
//	//push result to Stack
//	var t antlr.Token = ctx.GetOp()
//	slog.Trace(t.GetText())
//	right := v.Pop()
//	left := v.Pop()
//	switch t.GetTokenType() {
//	case parser.CalcParserMUL:
//		v.push(left * right)
//	case parser.CalcParserDIV:
//		v.push(left / right)
//	default:
//		panic("should not happen")
//
//	}
//
//	return nil
//}
//
//func (v *CalVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
//	slog.Trace("VisitAddSub======")
//	//push expression result to Stack
//	v.visitRule(ctx.Expr(0))
//	v.visitRule(ctx.Expr(1))
//
//	//push result to Stack
//	var t antlr.Token = ctx.GetOp()
//	right := v.Pop()
//	left := v.Pop()
//	switch t.GetTokenType() {
//	case parser.CalcParserADD:
//		v.push(left + right)
//	case parser.CalcParserSUB:
//		v.push(left - right)
//	default:
//		panic("should not happen")
//	}
//	slog.Trace("VisitAddSub======exit")
//	return nil
//
//}
