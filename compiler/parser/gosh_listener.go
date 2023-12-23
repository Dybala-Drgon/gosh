// Code generated from .//Gosh.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Gosh

import "github.com/antlr4-go/antlr/v4"

// GoshListener is a complete listener for a parse tree produced by GoshParser.
type GoshListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatements is called when entering the statements production.
	EnterStatements(c *StatementsContext)

	// EnterASSIGN is called when entering the ASSIGN production.
	EnterASSIGN(c *ASSIGNContext)

	// EnterFUNCCALL is called when entering the FUNCCALL production.
	EnterFUNCCALL(c *FUNCCALLContext)

	// EnterEXP is called when entering the EXP production.
	EnterEXP(c *EXPContext)

	// EnterFORLOOP is called when entering the FORLOOP production.
	EnterFORLOOP(c *FORLOOPContext)

	// EnterBLOCK is called when entering the BLOCK production.
	EnterBLOCK(c *BLOCKContext)

	// EnterFUNCDEF is called when entering the FUNCDEF production.
	EnterFUNCDEF(c *FUNCDEFContext)

	// EnterRUN is called when entering the RUN production.
	EnterRUN(c *RUNContext)

	// EnterRETURN is called when entering the RETURN production.
	EnterRETURN(c *RETURNContext)

	// EnterIF is called when entering the IF production.
	EnterIF(c *IFContext)

	// EnterBREAK is called when entering the BREAK production.
	EnterBREAK(c *BREAKContext)

	// EnterCONTINUE is called when entering the CONTINUE production.
	EnterCONTINUE(c *CONTINUEContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterLvalue is called when entering the lvalue production.
	EnterLvalue(c *LvalueContext)

	// EnterRvalue is called when entering the rvalue production.
	EnterRvalue(c *RvalueContext)

	// EnterConstvalue is called when entering the constvalue production.
	EnterConstvalue(c *ConstvalueContext)

	// EnterFunctionCall is called when entering the functionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterArguments is called when entering the arguments production.
	EnterArguments(c *ArgumentsContext)

	// EnterLoop is called when entering the loop production.
	EnterLoop(c *LoopContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterOuterAccess is called when entering the outerAccess production.
	EnterOuterAccess(c *OuterAccessContext)

	// EnterFunctionDefinition is called when entering the functionDefinition production.
	EnterFunctionDefinition(c *FunctionDefinitionContext)

	// EnterParameters is called when entering the parameters production.
	EnterParameters(c *ParametersContext)

	// EnterRunStatement is called when entering the runStatement production.
	EnterRunStatement(c *RunStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterIfStmt is called when entering the ifStmt production.
	EnterIfStmt(c *IfStmtContext)

	// EnterBreakStmt is called when entering the breakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the continueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterSimpleStmt is called when entering the simpleStmt production.
	EnterSimpleStmt(c *SimpleStmtContext)

	// EnterForClause is called when entering the forClause production.
	EnterForClause(c *ForClauseContext)

	// EnterMulDivOP is called when entering the mulDivOP production.
	EnterMulDivOP(c *MulDivOPContext)

	// EnterBinOP is called when entering the binOP production.
	EnterBinOP(c *BinOPContext)

	// EnterUnOP is called when entering the unOP production.
	EnterUnOP(c *UnOPContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatements is called when exiting the statements production.
	ExitStatements(c *StatementsContext)

	// ExitASSIGN is called when exiting the ASSIGN production.
	ExitASSIGN(c *ASSIGNContext)

	// ExitFUNCCALL is called when exiting the FUNCCALL production.
	ExitFUNCCALL(c *FUNCCALLContext)

	// ExitEXP is called when exiting the EXP production.
	ExitEXP(c *EXPContext)

	// ExitFORLOOP is called when exiting the FORLOOP production.
	ExitFORLOOP(c *FORLOOPContext)

	// ExitBLOCK is called when exiting the BLOCK production.
	ExitBLOCK(c *BLOCKContext)

	// ExitFUNCDEF is called when exiting the FUNCDEF production.
	ExitFUNCDEF(c *FUNCDEFContext)

	// ExitRUN is called when exiting the RUN production.
	ExitRUN(c *RUNContext)

	// ExitRETURN is called when exiting the RETURN production.
	ExitRETURN(c *RETURNContext)

	// ExitIF is called when exiting the IF production.
	ExitIF(c *IFContext)

	// ExitBREAK is called when exiting the BREAK production.
	ExitBREAK(c *BREAKContext)

	// ExitCONTINUE is called when exiting the CONTINUE production.
	ExitCONTINUE(c *CONTINUEContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitLvalue is called when exiting the lvalue production.
	ExitLvalue(c *LvalueContext)

	// ExitRvalue is called when exiting the rvalue production.
	ExitRvalue(c *RvalueContext)

	// ExitConstvalue is called when exiting the constvalue production.
	ExitConstvalue(c *ConstvalueContext)

	// ExitFunctionCall is called when exiting the functionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitArguments is called when exiting the arguments production.
	ExitArguments(c *ArgumentsContext)

	// ExitLoop is called when exiting the loop production.
	ExitLoop(c *LoopContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitOuterAccess is called when exiting the outerAccess production.
	ExitOuterAccess(c *OuterAccessContext)

	// ExitFunctionDefinition is called when exiting the functionDefinition production.
	ExitFunctionDefinition(c *FunctionDefinitionContext)

	// ExitParameters is called when exiting the parameters production.
	ExitParameters(c *ParametersContext)

	// ExitRunStatement is called when exiting the runStatement production.
	ExitRunStatement(c *RunStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitIfStmt is called when exiting the ifStmt production.
	ExitIfStmt(c *IfStmtContext)

	// ExitBreakStmt is called when exiting the breakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the continueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitSimpleStmt is called when exiting the simpleStmt production.
	ExitSimpleStmt(c *SimpleStmtContext)

	// ExitForClause is called when exiting the forClause production.
	ExitForClause(c *ForClauseContext)

	// ExitMulDivOP is called when exiting the mulDivOP production.
	ExitMulDivOP(c *MulDivOPContext)

	// ExitBinOP is called when exiting the binOP production.
	ExitBinOP(c *BinOPContext)

	// ExitUnOP is called when exiting the unOP production.
	ExitUnOP(c *UnOPContext)
}
