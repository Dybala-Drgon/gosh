package bytecode

import (
	"github.com/gookit/slog"
	"gosh/compiler/parser"
	"gosh/compiler/token"
)

/**
1. 插入标记:进入某张符号表
2. 编译stmt
3. 退出该层符号表
*/

func (v *GoshVisitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	slog.Info("enter blk: ", ctx.GetText())
	origin := v.CurSymTableIdx
	// 新增符号表
	newSymbol := v.AddSymbolTable()
	v.emit(token.OpJumpSymTable, newSymbol.idx)
	v.CurSymTableIdx = newSymbol.idx
	v.visitRule(ctx.Statements())
	v.CurSymTableIdx = origin
	v.emit(token.OpExitSymTable)
	return nil
}
