package test
//
//import (
//	"github.com/antlr4-go/antlr/v4"
//	"gosh/compiler/bytecode"
//	"gosh/compiler/parser"
//	"gosh/core/vm"
//	"testing"
//)
//
//type goshErrorListener struct {
//	*antlr.DefaultErrorListener
//}
//
//func RunGosh(input string) *vm.VM {
//	is := antlr.NewInputStream(input)
//	lexer := parser.NewGoshLexer(is)
//	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
//	p := parser.NewGoshParser(stream)
//	p.AddErrorListener(&goshErrorListener{})
//	visitor := &bytecode.GoshVisitor{}
//	visitor.InitCompiler()
//	s := p.Program()
//
//	//slog.Info(s.GetText())
//	code := s.Accept(visitor)
//	//slog.Trace(reflect.TypeOf(code))
//	virtualMachine := vm.NewVM(code.(*bytecode.Bytecode))
//	virtualMachine.Run()
//	return virtualMachine
//}
//
//func TestOperands(t *testing.T) {
//	RunGosh(`a=2
//	b=3
//	a<b
//a!=b
//a==b
//a<=b
//a>=b
//a>b
//	`)
//}
