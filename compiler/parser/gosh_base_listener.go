// Code generated from .//Gosh.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Gosh

import "github.com/antlr4-go/antlr/v4"

// BaseGoshListener is a complete listener for a parse tree produced by GoshParser.
type BaseGoshListener struct{}

var _ GoshListener = &BaseGoshListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoshListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoshListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoshListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoshListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseGoshListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseGoshListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatements is called when production statements is entered.
func (s *BaseGoshListener) EnterStatements(ctx *StatementsContext) {}

// ExitStatements is called when production statements is exited.
func (s *BaseGoshListener) ExitStatements(ctx *StatementsContext) {}

// EnterASSIGN is called when production ASSIGN is entered.
func (s *BaseGoshListener) EnterASSIGN(ctx *ASSIGNContext) {}

// ExitASSIGN is called when production ASSIGN is exited.
func (s *BaseGoshListener) ExitASSIGN(ctx *ASSIGNContext) {}

// EnterFUNCCALL is called when production FUNCCALL is entered.
func (s *BaseGoshListener) EnterFUNCCALL(ctx *FUNCCALLContext) {}

// ExitFUNCCALL is called when production FUNCCALL is exited.
func (s *BaseGoshListener) ExitFUNCCALL(ctx *FUNCCALLContext) {}

// EnterEXP is called when production EXP is entered.
func (s *BaseGoshListener) EnterEXP(ctx *EXPContext) {}

// ExitEXP is called when production EXP is exited.
func (s *BaseGoshListener) ExitEXP(ctx *EXPContext) {}

// EnterFORLOOP is called when production FORLOOP is entered.
func (s *BaseGoshListener) EnterFORLOOP(ctx *FORLOOPContext) {}

// ExitFORLOOP is called when production FORLOOP is exited.
func (s *BaseGoshListener) ExitFORLOOP(ctx *FORLOOPContext) {}

// EnterBLOCK is called when production BLOCK is entered.
func (s *BaseGoshListener) EnterBLOCK(ctx *BLOCKContext) {}

// ExitBLOCK is called when production BLOCK is exited.
func (s *BaseGoshListener) ExitBLOCK(ctx *BLOCKContext) {}

// EnterFUNCDEF is called when production FUNCDEF is entered.
func (s *BaseGoshListener) EnterFUNCDEF(ctx *FUNCDEFContext) {}

// ExitFUNCDEF is called when production FUNCDEF is exited.
func (s *BaseGoshListener) ExitFUNCDEF(ctx *FUNCDEFContext) {}

// EnterRUN is called when production RUN is entered.
func (s *BaseGoshListener) EnterRUN(ctx *RUNContext) {}

// ExitRUN is called when production RUN is exited.
func (s *BaseGoshListener) ExitRUN(ctx *RUNContext) {}

// EnterRETURN is called when production RETURN is entered.
func (s *BaseGoshListener) EnterRETURN(ctx *RETURNContext) {}

// ExitRETURN is called when production RETURN is exited.
func (s *BaseGoshListener) ExitRETURN(ctx *RETURNContext) {}

// EnterIF is called when production IF is entered.
func (s *BaseGoshListener) EnterIF(ctx *IFContext) {}

// ExitIF is called when production IF is exited.
func (s *BaseGoshListener) ExitIF(ctx *IFContext) {}

// EnterBREAK is called when production BREAK is entered.
func (s *BaseGoshListener) EnterBREAK(ctx *BREAKContext) {}

// ExitBREAK is called when production BREAK is exited.
func (s *BaseGoshListener) ExitBREAK(ctx *BREAKContext) {}

// EnterCONTINUE is called when production CONTINUE is entered.
func (s *BaseGoshListener) EnterCONTINUE(ctx *CONTINUEContext) {}

// ExitCONTINUE is called when production CONTINUE is exited.
func (s *BaseGoshListener) ExitCONTINUE(ctx *CONTINUEContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseGoshListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseGoshListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterLvalue is called when production lvalue is entered.
func (s *BaseGoshListener) EnterLvalue(ctx *LvalueContext) {}

// ExitLvalue is called when production lvalue is exited.
func (s *BaseGoshListener) ExitLvalue(ctx *LvalueContext) {}

// EnterRvalue is called when production rvalue is entered.
func (s *BaseGoshListener) EnterRvalue(ctx *RvalueContext) {}

// ExitRvalue is called when production rvalue is exited.
func (s *BaseGoshListener) ExitRvalue(ctx *RvalueContext) {}

// EnterConstvalue is called when production constvalue is entered.
func (s *BaseGoshListener) EnterConstvalue(ctx *ConstvalueContext) {}

// ExitConstvalue is called when production constvalue is exited.
func (s *BaseGoshListener) ExitConstvalue(ctx *ConstvalueContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseGoshListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseGoshListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArguments is called when production arguments is entered.
func (s *BaseGoshListener) EnterArguments(ctx *ArgumentsContext) {}

// ExitArguments is called when production arguments is exited.
func (s *BaseGoshListener) ExitArguments(ctx *ArgumentsContext) {}

// EnterLoop is called when production loop is entered.
func (s *BaseGoshListener) EnterLoop(ctx *LoopContext) {}

// ExitLoop is called when production loop is exited.
func (s *BaseGoshListener) ExitLoop(ctx *LoopContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGoshListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGoshListener) ExitExpression(ctx *ExpressionContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseGoshListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseGoshListener) ExitBlock(ctx *BlockContext) {}

// EnterOuterAccess is called when production outerAccess is entered.
func (s *BaseGoshListener) EnterOuterAccess(ctx *OuterAccessContext) {}

// ExitOuterAccess is called when production outerAccess is exited.
func (s *BaseGoshListener) ExitOuterAccess(ctx *OuterAccessContext) {}

// EnterFunctionDefinition is called when production functionDefinition is entered.
func (s *BaseGoshListener) EnterFunctionDefinition(ctx *FunctionDefinitionContext) {}

// ExitFunctionDefinition is called when production functionDefinition is exited.
func (s *BaseGoshListener) ExitFunctionDefinition(ctx *FunctionDefinitionContext) {}

// EnterParameters is called when production parameters is entered.
func (s *BaseGoshListener) EnterParameters(ctx *ParametersContext) {}

// ExitParameters is called when production parameters is exited.
func (s *BaseGoshListener) ExitParameters(ctx *ParametersContext) {}

// EnterRunStatement is called when production runStatement is entered.
func (s *BaseGoshListener) EnterRunStatement(ctx *RunStatementContext) {}

// ExitRunStatement is called when production runStatement is exited.
func (s *BaseGoshListener) ExitRunStatement(ctx *RunStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseGoshListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseGoshListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterIfStmt is called when production ifStmt is entered.
func (s *BaseGoshListener) EnterIfStmt(ctx *IfStmtContext) {}

// ExitIfStmt is called when production ifStmt is exited.
func (s *BaseGoshListener) ExitIfStmt(ctx *IfStmtContext) {}

// EnterBreakStmt is called when production breakStmt is entered.
func (s *BaseGoshListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production breakStmt is exited.
func (s *BaseGoshListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production continueStmt is entered.
func (s *BaseGoshListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production continueStmt is exited.
func (s *BaseGoshListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterSimpleStmt is called when production simpleStmt is entered.
func (s *BaseGoshListener) EnterSimpleStmt(ctx *SimpleStmtContext) {}

// ExitSimpleStmt is called when production simpleStmt is exited.
func (s *BaseGoshListener) ExitSimpleStmt(ctx *SimpleStmtContext) {}

// EnterForClause is called when production forClause is entered.
func (s *BaseGoshListener) EnterForClause(ctx *ForClauseContext) {}

// ExitForClause is called when production forClause is exited.
func (s *BaseGoshListener) ExitForClause(ctx *ForClauseContext) {}

// EnterMulDivOP is called when production mulDivOP is entered.
func (s *BaseGoshListener) EnterMulDivOP(ctx *MulDivOPContext) {}

// ExitMulDivOP is called when production mulDivOP is exited.
func (s *BaseGoshListener) ExitMulDivOP(ctx *MulDivOPContext) {}

// EnterBinOP is called when production binOP is entered.
func (s *BaseGoshListener) EnterBinOP(ctx *BinOPContext) {}

// ExitBinOP is called when production binOP is exited.
func (s *BaseGoshListener) ExitBinOP(ctx *BinOPContext) {}

// EnterUnOP is called when production unOP is entered.
func (s *BaseGoshListener) EnterUnOP(ctx *UnOPContext) {}

// ExitUnOP is called when production unOP is exited.
func (s *BaseGoshListener) ExitUnOP(ctx *UnOPContext) {}
