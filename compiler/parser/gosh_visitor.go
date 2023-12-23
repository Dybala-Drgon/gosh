// Code generated from .//Gosh.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Gosh

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GoshParser.
type GoshVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GoshParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by GoshParser#statements.
	VisitStatements(ctx *StatementsContext) interface{}

	// Visit a parse tree produced by GoshParser#ASSIGN.
	VisitASSIGN(ctx *ASSIGNContext) interface{}

	// Visit a parse tree produced by GoshParser#FUNCCALL.
	VisitFUNCCALL(ctx *FUNCCALLContext) interface{}

	// Visit a parse tree produced by GoshParser#EXP.
	VisitEXP(ctx *EXPContext) interface{}

	// Visit a parse tree produced by GoshParser#FORLOOP.
	VisitFORLOOP(ctx *FORLOOPContext) interface{}

	// Visit a parse tree produced by GoshParser#BLOCK.
	VisitBLOCK(ctx *BLOCKContext) interface{}

	// Visit a parse tree produced by GoshParser#FUNCDEF.
	VisitFUNCDEF(ctx *FUNCDEFContext) interface{}

	// Visit a parse tree produced by GoshParser#RUN.
	VisitRUN(ctx *RUNContext) interface{}

	// Visit a parse tree produced by GoshParser#RETURN.
	VisitRETURN(ctx *RETURNContext) interface{}

	// Visit a parse tree produced by GoshParser#IF.
	VisitIF(ctx *IFContext) interface{}

	// Visit a parse tree produced by GoshParser#BREAK.
	VisitBREAK(ctx *BREAKContext) interface{}

	// Visit a parse tree produced by GoshParser#CONTINUE.
	VisitCONTINUE(ctx *CONTINUEContext) interface{}

	// Visit a parse tree produced by GoshParser#assignment.
	VisitAssignment(ctx *AssignmentContext) interface{}

	// Visit a parse tree produced by GoshParser#lvalue.
	VisitLvalue(ctx *LvalueContext) interface{}

	// Visit a parse tree produced by GoshParser#rvalue.
	VisitRvalue(ctx *RvalueContext) interface{}

	// Visit a parse tree produced by GoshParser#constvalue.
	VisitConstvalue(ctx *ConstvalueContext) interface{}

	// Visit a parse tree produced by GoshParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by GoshParser#arguments.
	VisitArguments(ctx *ArgumentsContext) interface{}

	// Visit a parse tree produced by GoshParser#loop.
	VisitLoop(ctx *LoopContext) interface{}

	// Visit a parse tree produced by GoshParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by GoshParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by GoshParser#outerAccess.
	VisitOuterAccess(ctx *OuterAccessContext) interface{}

	// Visit a parse tree produced by GoshParser#functionDefinition.
	VisitFunctionDefinition(ctx *FunctionDefinitionContext) interface{}

	// Visit a parse tree produced by GoshParser#parameters.
	VisitParameters(ctx *ParametersContext) interface{}

	// Visit a parse tree produced by GoshParser#runStatement.
	VisitRunStatement(ctx *RunStatementContext) interface{}

	// Visit a parse tree produced by GoshParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by GoshParser#ifStmt.
	VisitIfStmt(ctx *IfStmtContext) interface{}

	// Visit a parse tree produced by GoshParser#breakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by GoshParser#continueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by GoshParser#simpleStmt.
	VisitSimpleStmt(ctx *SimpleStmtContext) interface{}

	// Visit a parse tree produced by GoshParser#forClause.
	VisitForClause(ctx *ForClauseContext) interface{}

	// Visit a parse tree produced by GoshParser#mulDivOP.
	VisitMulDivOP(ctx *MulDivOPContext) interface{}

	// Visit a parse tree produced by GoshParser#binOP.
	VisitBinOP(ctx *BinOPContext) interface{}

	// Visit a parse tree produced by GoshParser#unOP.
	VisitUnOP(ctx *UnOPContext) interface{}
}
