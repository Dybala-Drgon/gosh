package bytecode

import (
	"github.com/gookit/slog"
	"gosh/compiler/parser"
	"gosh/compiler/token"
)

// FunInfo 表示函数信息
type FunInfo struct {
	ParamNum  int
	ReturnNum int
	Name      string
	SymIdx    int
}

func (v *GoshVisitor) VisitFUNCDEF(ctx *parser.FUNCDEFContext) interface{} {
	return v.visitRule(ctx.FunctionDefinition())
}

func (v *GoshVisitor) VisitFunctionDefinition(ctx *parser.FunctionDefinitionContext) interface{} {
	funcName := ctx.ID().GetText()
	origin := v.CurSymTableIdx
	newSymbol := v.AddSymbolTable()
	v.CurSymTableIdx = newSymbol.idx
	for _, name := range ctx.Parameters().AllID() {
		newSymbol.Define(name.GetText(), true)
	}

	var info FunInfo
	paramNum := len(ctx.Parameters().AllID())
	info.Name = funcName
	info.SymIdx = newSymbol.idx
	info.ParamNum = paramNum
	v.FuncTable[funcName] = info

	returnNum := v.visitRule(ctx.Block())
	info.ReturnNum = returnNum.(int)
	slog.Info(info)
	v.CurSymTableIdx = origin
	v.FuncTable[funcName] = info
	return nil
}

func (v *GoshVisitor) VisitRETURN(ctx *parser.RETURNContext) interface{} {
	return v.visitRule(ctx.ReturnStatement())
}

func (v *GoshVisitor) VisitReturnStatement(ctx *parser.ReturnStatementContext) interface{} {
	res := len(ctx.AllExpression())
	for _, elem := range ctx.AllExpression() {
		v.visitRule(elem)
	}
	// 直接jmp出去
	v.emit(token.OpReturn)
	return res
}

func (v *GoshVisitor) VisitFUNCCALL(ctx *parser.FUNCCALLContext) interface{} {
	return v.visitRule(ctx.FunctionCall())
}

func (v *GoshVisitor) VisitFunctionCall(ctx *parser.FunctionCallContext) interface{} {
	// 1. 校验参数个数对不对的上
	// 2. 压参数
	// 3. call
	args := ctx.Arguments()
	funcName := ctx.ID().GetText()
	tmp := v.FuncTable[funcName]
	if len(args.AllExpression()) != tmp.ParamNum {
		slog.Error(funcName, " 参数不匹配")
		panic("参数个数不匹配")
	}
	// push stack
	for _, exp := range args.AllExpression() {
		v.visitRule(exp)
	}
	// 设置Frame中参数
	sidx := tmp.SymIdx
	// 拿到该函数的符号表
	symbol := v.SymbolTables[sidx]
	for i := len(symbol.symbols) - 1; i >= 0; i-- {
		v.emit(token.OpSetVar, sidx, i)
	}

	v.emit(token.OpCall, tmp.SymIdx)
	return nil
}
