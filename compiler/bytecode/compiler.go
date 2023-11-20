package bytecode

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/gookit/slog"
	"gosh/compiler/parser"
	"gosh/compiler/token"
	"gosh/core/object"
)

type Bytecode struct {
	Instructions []*CompilationScope // 可能有多个指令域
	Constants    []object.Object     // 常量也需要保存下来
	SymbolTables []*SymbolTable
}

type GoshVisitor struct {
	*parser.BaseGoshVisitor
	SymbolTables   []*SymbolTable
	Ins            []*CompilationScope // 可能有多个指令域
	Constants      []object.Object
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
	v.Ins = append(v.Ins, &CompilationScope{})
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
	return node.Accept(v)
}

func (v *GoshVisitor) VisitStatements(ctx *parser.StatementsContext) interface{} {
	slog.Trace("visit stmt, count = ", ctx.GetChildCount())
	for _, elem := range ctx.AllStatement() {
		return v.visitRule(elem)
	}
	return nil
}

func (v *GoshVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	slog.Trace("visit program")
	v.visitRule(ctx.Statements())
	return &Bytecode{Instructions: v.Ins,
		Constants: v.Constants, SymbolTables: v.SymbolTables}
}
func (v *GoshVisitor) addConstant(o object.Object) int {
	v.Constants = append(v.Constants, o)
	return len(v.Constants) - 1
}

func (v *GoshVisitor) currentInstructions() (b []byte) {
	return v.Ins[v.CurSymTableIdx].Instruction
}

func (v *GoshVisitor) addInstruction(b []byte) int {
	posNewIns := len(v.currentInstructions())
	v.Ins[v.CurSymTableIdx].Instruction = append(v.Ins[v.CurSymTableIdx].Instruction, b...)
	return posNewIns
}

func (v *GoshVisitor) emit(opcode token.Opcode, operands ...int) int {
	inst := MakeInstruction(opcode, operands...)
	pos := v.addInstruction(inst)
	return pos
}
