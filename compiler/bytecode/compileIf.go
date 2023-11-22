package bytecode

import (
	"github.com/gookit/slog"
	"gosh/compiler/token"

	//"go/token"
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
	v.emit(token.OpIf)
	// blk有且仅有一个
	var res int
	v.visitRule(ctx.Block(0))
	if ctx.ELSE() == nil {
		return res
	}

	pos1 := v.emit(token.OpJump, 0)

	count := ctx.GetChildCount()
	// 获取最后一个子节点
	kid := ctx.GetChild(count - 1)
	//slog.Info(reflect.TypeOf(kid))
	switch t := kid.(type) {
	case *parser.IfStmtContext:
		v.visitRule(t)
	case *parser.BlockContext:
		v.visitRule(t)
	default:
		slog.Info("暂未实现")
	}

	curPos := len(v.currentInstructions())
	v.changeOperand(pos1, curPos)
	slog.Info("出口位置: ", curPos)
	return nil
}
