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
	//RunGosh(`
	//t4=999
	//c,b,a,t,w2=2.2,"str",1+2*5,t4,-8
	//`)
	//RunGosh(`c,b,a,t,w=2.2,"str",1+2*5,8-(5+1),-8`)
	//RunGosh(`c,b,a,t,w=2.2,"str",1+2*5,-8,fun()fun`)
	//RunGosh(`7+8*9+(8-4)*2`)
	//RunGosh(`-8`)
	//RunGosh(`t=3
	//		a=2
	//		if (t<2 )|| (t==4 ){
	//			t = t+1
	//		a=100
	//		}
	//	else if t==3{
	//			t= t-1
	//a=897
	//		}
	//	else{
	//		t=10
	//	a = 23
	//	}
	//		`)
	//RunGosh(`t=3
	//	a=2
	//	(t<2) || (t==3)
	//`)

	//RunGosh(`
	//b=1
	//i=0
	//c="str"
	//for ;i<10;i=i+1{
	//	if i == 0 {
	//		tt = 9
	//		b = b+5
	//	}
	// b = b+1
	//}
	//c=b
	//`)
	//RunGosh(`a,=1`)
	RunGosh(`func first(a,b,c){
		   return c,b,a
		}
	
		a,b,c=first(1,2,3)
	`)
	//RunGosh(`
	//func fib(a){
	//	if a< 3{
	//		return 1
	//	}
	//	f1=1
	//	f2=1
	//	for i=2;i<a;i=i+1{
	//		f2,f1=f1+f2,f2
	//	}
	//	return f2
	//}
	//res=fib(10)
	//`)
	//	RunGosh(`for i=0;i<10;i=i+1{
	//    if i==5{
	//        break
	//    }
	//}`)

	//	RunGosh(`a=2
	//	b=3
	//	a<b
	//a!=b
	//a==b
	//a<=b
	//a>=b
	//a>b
	//	`)
}
