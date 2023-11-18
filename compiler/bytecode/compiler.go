package bytecode

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"gosh/compiler/parser"
)

type Bytecode struct {
	Instructions []byte
	//Constants    []object.Object
}

type GoshVisitor struct {
	parser.BaseGoshVisitor
	SymbolTables   []*SymbolTable
	CurSymTableIdx int
}

// InitCompiler 初始化编译器
func (v *GoshVisitor) InitCompiler() {
	globalSymbol := NewSymbolTable()
	if len(v.SymbolTables) == 0 {
		globalSymbol.parent = nil
		v.CurSymTableIdx = 0
	} else {
		globalSymbol.parent = v.SymbolTables[v.CurSymTableIdx]
	}
	v.SymbolTables = append(v.SymbolTables, globalSymbol)
}

func (v *GoshVisitor) AddSymbolTable() *SymbolTable {
	newSymbolTable := NewSymbolTable()
	if len(v.SymbolTables) == 0 {
		newSymbolTable.parent = nil
		v.CurSymTableIdx = 0
	} else {
		newSymbolTable.parent = v.SymbolTables[v.CurSymTableIdx]
	}
	newSymbolTable.idx = len(v.SymbolTables)
	v.SymbolTables = append(v.SymbolTables, newSymbolTable)
	return newSymbolTable
}

func (v *GoshVisitor) visitRule(node antlr.RuleNode) interface{} {
	node.Accept(v)
	return nil
}

func (v *GoshVisitor) VisitStatements(ctx *parser.StatementsContext) interface{} {
	fmt.Println("visit stmt, count = ", ctx.GetChildCount())
	for _, elem := range ctx.AllStatement() {
		return v.visitRule(elem)
	}
	return nil
}

func (v *GoshVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	fmt.Println("visit program")
	return v.visitRule(ctx.Statements())
}

func (v *GoshVisitor) VisitASSIGN(ctx *parser.ASSIGNContext) interface{} {
	fmt.Println("enter assign stmt", ctx.GetText(), ctx.GetChildCount())
	//return v.visitRule(ctx.Assignment().Lvalue())
	asg := ctx.Assignment()
	table := v.SymbolTables[v.CurSymTableIdx]
	for _, id := range asg.Lvalue().AllID() {
		table.Define(id.GetText())
	}

	v.visitRule(asg.Lvalue())
	v.visitRule(asg.Rvalue())
	return nil
}

func (v *GoshVisitor) VisitLvalue(ctx *parser.LvalueContext) interface{} {
	fmt.Println("enter lvalue =", ctx.GetText())
	return nil
}
