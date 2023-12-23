// Code generated from .//Gosh.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Gosh

import "github.com/antlr4-go/antlr/v4"

type BaseGoshVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGoshVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitStatements(ctx *StatementsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitASSIGN(ctx *ASSIGNContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitFUNCCALL(ctx *FUNCCALLContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitEXP(ctx *EXPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitFORLOOP(ctx *FORLOOPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitBLOCK(ctx *BLOCKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitFUNCDEF(ctx *FUNCDEFContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitRUN(ctx *RUNContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitRETURN(ctx *RETURNContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitIF(ctx *IFContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitBREAK(ctx *BREAKContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitCONTINUE(ctx *CONTINUEContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitAssignment(ctx *AssignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitLvalue(ctx *LvalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitRvalue(ctx *RvalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitConstvalue(ctx *ConstvalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitArguments(ctx *ArgumentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitLoop(ctx *LoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitOuterAccess(ctx *OuterAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitParameters(ctx *ParametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitRunStatement(ctx *RunStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitIfStmt(ctx *IfStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitSimpleStmt(ctx *SimpleStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitForClause(ctx *ForClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitMulDivOP(ctx *MulDivOPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitBinOP(ctx *BinOPContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoshVisitor) VisitUnOP(ctx *UnOPContext) interface{} {
	return v.VisitChildren(ctx)
}
