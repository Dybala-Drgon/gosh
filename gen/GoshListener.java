// Generated from E:/project/go/gosh/compiler/parser/Gosh.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link GoshParser}.
 */
public interface GoshListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link GoshParser#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(GoshParser.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(GoshParser.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#statements}.
	 * @param ctx the parse tree
	 */
	void enterStatements(GoshParser.StatementsContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#statements}.
	 * @param ctx the parse tree
	 */
	void exitStatements(GoshParser.StatementsContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ASSIGN}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterASSIGN(GoshParser.ASSIGNContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ASSIGN}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitASSIGN(GoshParser.ASSIGNContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FUNCCALL}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterFUNCCALL(GoshParser.FUNCCALLContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FUNCCALL}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitFUNCCALL(GoshParser.FUNCCALLContext ctx);
	/**
	 * Enter a parse tree produced by the {@code EXP}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterEXP(GoshParser.EXPContext ctx);
	/**
	 * Exit a parse tree produced by the {@code EXP}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitEXP(GoshParser.EXPContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FORLOOP}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterFORLOOP(GoshParser.FORLOOPContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FORLOOP}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitFORLOOP(GoshParser.FORLOOPContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BLOCK}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterBLOCK(GoshParser.BLOCKContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BLOCK}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitBLOCK(GoshParser.BLOCKContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FUNCDEF}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterFUNCDEF(GoshParser.FUNCDEFContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FUNCDEF}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitFUNCDEF(GoshParser.FUNCDEFContext ctx);
	/**
	 * Enter a parse tree produced by the {@code RUN}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterRUN(GoshParser.RUNContext ctx);
	/**
	 * Exit a parse tree produced by the {@code RUN}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitRUN(GoshParser.RUNContext ctx);
	/**
	 * Enter a parse tree produced by the {@code RETURN}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterRETURN(GoshParser.RETURNContext ctx);
	/**
	 * Exit a parse tree produced by the {@code RETURN}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitRETURN(GoshParser.RETURNContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IF}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterIF(GoshParser.IFContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IF}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitIF(GoshParser.IFContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BREAK}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterBREAK(GoshParser.BREAKContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BREAK}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitBREAK(GoshParser.BREAKContext ctx);
	/**
	 * Enter a parse tree produced by the {@code CONTINUE}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterCONTINUE(GoshParser.CONTINUEContext ctx);
	/**
	 * Exit a parse tree produced by the {@code CONTINUE}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitCONTINUE(GoshParser.CONTINUEContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#assignment}.
	 * @param ctx the parse tree
	 */
	void enterAssignment(GoshParser.AssignmentContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#assignment}.
	 * @param ctx the parse tree
	 */
	void exitAssignment(GoshParser.AssignmentContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#lvalue}.
	 * @param ctx the parse tree
	 */
	void enterLvalue(GoshParser.LvalueContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#lvalue}.
	 * @param ctx the parse tree
	 */
	void exitLvalue(GoshParser.LvalueContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#rvalue}.
	 * @param ctx the parse tree
	 */
	void enterRvalue(GoshParser.RvalueContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#rvalue}.
	 * @param ctx the parse tree
	 */
	void exitRvalue(GoshParser.RvalueContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#constvalue}.
	 * @param ctx the parse tree
	 */
	void enterConstvalue(GoshParser.ConstvalueContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#constvalue}.
	 * @param ctx the parse tree
	 */
	void exitConstvalue(GoshParser.ConstvalueContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#functionCall}.
	 * @param ctx the parse tree
	 */
	void enterFunctionCall(GoshParser.FunctionCallContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#functionCall}.
	 * @param ctx the parse tree
	 */
	void exitFunctionCall(GoshParser.FunctionCallContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#arguments}.
	 * @param ctx the parse tree
	 */
	void enterArguments(GoshParser.ArgumentsContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#arguments}.
	 * @param ctx the parse tree
	 */
	void exitArguments(GoshParser.ArgumentsContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#loop}.
	 * @param ctx the parse tree
	 */
	void enterLoop(GoshParser.LoopContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#loop}.
	 * @param ctx the parse tree
	 */
	void exitLoop(GoshParser.LoopContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(GoshParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(GoshParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#block}.
	 * @param ctx the parse tree
	 */
	void enterBlock(GoshParser.BlockContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#block}.
	 * @param ctx the parse tree
	 */
	void exitBlock(GoshParser.BlockContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#outerAccess}.
	 * @param ctx the parse tree
	 */
	void enterOuterAccess(GoshParser.OuterAccessContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#outerAccess}.
	 * @param ctx the parse tree
	 */
	void exitOuterAccess(GoshParser.OuterAccessContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#functionDefinition}.
	 * @param ctx the parse tree
	 */
	void enterFunctionDefinition(GoshParser.FunctionDefinitionContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#functionDefinition}.
	 * @param ctx the parse tree
	 */
	void exitFunctionDefinition(GoshParser.FunctionDefinitionContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#parameters}.
	 * @param ctx the parse tree
	 */
	void enterParameters(GoshParser.ParametersContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#parameters}.
	 * @param ctx the parse tree
	 */
	void exitParameters(GoshParser.ParametersContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#runStatement}.
	 * @param ctx the parse tree
	 */
	void enterRunStatement(GoshParser.RunStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#runStatement}.
	 * @param ctx the parse tree
	 */
	void exitRunStatement(GoshParser.RunStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#returnStatement}.
	 * @param ctx the parse tree
	 */
	void enterReturnStatement(GoshParser.ReturnStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#returnStatement}.
	 * @param ctx the parse tree
	 */
	void exitReturnStatement(GoshParser.ReturnStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#ifStmt}.
	 * @param ctx the parse tree
	 */
	void enterIfStmt(GoshParser.IfStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#ifStmt}.
	 * @param ctx the parse tree
	 */
	void exitIfStmt(GoshParser.IfStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#breakStmt}.
	 * @param ctx the parse tree
	 */
	void enterBreakStmt(GoshParser.BreakStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#breakStmt}.
	 * @param ctx the parse tree
	 */
	void exitBreakStmt(GoshParser.BreakStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#continueStmt}.
	 * @param ctx the parse tree
	 */
	void enterContinueStmt(GoshParser.ContinueStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#continueStmt}.
	 * @param ctx the parse tree
	 */
	void exitContinueStmt(GoshParser.ContinueStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#simpleStmt}.
	 * @param ctx the parse tree
	 */
	void enterSimpleStmt(GoshParser.SimpleStmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#simpleStmt}.
	 * @param ctx the parse tree
	 */
	void exitSimpleStmt(GoshParser.SimpleStmtContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#forClause}.
	 * @param ctx the parse tree
	 */
	void enterForClause(GoshParser.ForClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#forClause}.
	 * @param ctx the parse tree
	 */
	void exitForClause(GoshParser.ForClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#mulDivOP}.
	 * @param ctx the parse tree
	 */
	void enterMulDivOP(GoshParser.MulDivOPContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#mulDivOP}.
	 * @param ctx the parse tree
	 */
	void exitMulDivOP(GoshParser.MulDivOPContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#binOP}.
	 * @param ctx the parse tree
	 */
	void enterBinOP(GoshParser.BinOPContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#binOP}.
	 * @param ctx the parse tree
	 */
	void exitBinOP(GoshParser.BinOPContext ctx);
	/**
	 * Enter a parse tree produced by {@link GoshParser#unOP}.
	 * @param ctx the parse tree
	 */
	void enterUnOP(GoshParser.UnOPContext ctx);
	/**
	 * Exit a parse tree produced by {@link GoshParser#unOP}.
	 * @param ctx the parse tree
	 */
	void exitUnOP(GoshParser.UnOPContext ctx);
}