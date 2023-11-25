package bytecode

import (
	"github.com/gookit/slog"
	"gosh/compiler/parser"
	"gosh/compiler/token"
)

func (v *GoshVisitor) VisitFORLOOP(ctx *parser.FORLOOPContext) interface{} {
	slog.Trace("enter for loop: ", ctx.GetText())
	return v.visitRule(ctx.Loop())
}

/*
*
1. init
2. judge
3. blk
4. post
5. judge
*/
func (v *GoshVisitor) VisitLoop(ctx *parser.LoopContext) interface{} {
	slog.Info("has ", ctx.GetChildCount(), " kids")
	forClause := ctx.ForClause()
	blk := ctx.Block()
	slog.Info(forClause.GetText())
	initStmt := forClause.GetInitStmt()
	exp := forClause.Expression()
	postStmt := forClause.GetPostStmt()
	if initStmt != nil {
		v.visitRule(initStmt)
	}
	if exp == nil {
		v.emit(token.OpConstant, 1)
	} else {
		v.visitRule(exp)
	}

	v.emit(token.OpForIfPre)
	pos1 := v.emit(token.OpJump, 0)
	pos := len(v.currentInstructions())
	// 首轮校验
	v.visitRule(blk)
	if postStmt != nil {
		v.visitRule(postStmt)
	}
	if exp == nil {
		v.emit(token.OpConstant, 1)
	} else {
		v.visitRule(exp)
	}
	v.emit(token.OpForIfPost)
	v.emit(token.OpJump, pos)
	curPos := len(v.currentInstructions())
	v.changeOperand(pos1, curPos)
	slog.Info("for loop出口位置: ", curPos)
	return nil
}

func (v *GoshVisitor) VisitSimpleStmt(ctx *parser.SimpleStmtContext) interface{} {
	return v.visitRule(ctx.Assignment())
}

func (v *GoshVisitor) VisitBREAK(ctx *parser.BREAKContext) interface{} {
	return v.visitRule(ctx.BreakStmt())
}

func (v *GoshVisitor) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {
	v.emit(token.OpBreak)
	return nil
}
