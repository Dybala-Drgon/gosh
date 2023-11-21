package bytecode

import (
	"github.com/gookit/slog"
	"gosh/compiler/parser"
)

/*
*
1. 编译中间的表达式
2. 插入一个桩代表else执行的语句
3. 编译Block语句(考虑符号表)
4. block之后的语句，编译之后将指令地址传给第二部中的桩地址
*/

func (v *GoshVisitor) VisitIF(ctx *parser.IFContext) interface{} {
	return v.visitRule(ctx.IfStmt())

}

func (v *GoshVisitor) VisitIfStmt(ctx *parser.IfStmtContext) interface{} {
	slog.Info("enter if stmt", ctx.GetText())
	v.visitRule(ctx.Expression())
	// 首先打一个桩
	// 第符号表idx

	//jumpPos1 := v.emit(token.OpJumpFalsy, 0)
	v.visitRule(ctx.Block(0))

	return nil
}
