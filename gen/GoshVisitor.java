// Generated from E:/project/go/gosh/compiler/parser/Gosh.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeVisitor;

/**
 * This interface defines a complete generic visitor for a parse tree produced
 * by {@link GoshParser}.
 *
 * @param <T> The return type of the visit operation. Use {@link Void} for
 * operations with no return type.
 */
public interface GoshVisitor<T> extends ParseTreeVisitor<T> {
	/**
	 * Visit a parse tree produced by {@link GoshParser#program}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitProgram(GoshParser.ProgramContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#statements}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitStatements(GoshParser.StatementsContext ctx);
	/**
	 * Visit a parse tree produced by the {@code ASSIGN}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitASSIGN(GoshParser.ASSIGNContext ctx);
	/**
	 * Visit a parse tree produced by the {@code FUNCCALL}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFUNCCALL(GoshParser.FUNCCALLContext ctx);
	/**
	 * Visit a parse tree produced by the {@code EXP}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitEXP(GoshParser.EXPContext ctx);
	/**
	 * Visit a parse tree produced by the {@code FORLOOP}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFORLOOP(GoshParser.FORLOOPContext ctx);
	/**
	 * Visit a parse tree produced by the {@code BLOCK}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBLOCK(GoshParser.BLOCKContext ctx);
	/**
	 * Visit a parse tree produced by the {@code FUNCDEF}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFUNCDEF(GoshParser.FUNCDEFContext ctx);
	/**
	 * Visit a parse tree produced by the {@code RUN}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRUN(GoshParser.RUNContext ctx);
	/**
	 * Visit a parse tree produced by the {@code RETURN}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRETURN(GoshParser.RETURNContext ctx);
	/**
	 * Visit a parse tree produced by the {@code IF}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIF(GoshParser.IFContext ctx);
	/**
	 * Visit a parse tree produced by the {@code BREAK}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBREAK(GoshParser.BREAKContext ctx);
	/**
	 * Visit a parse tree produced by the {@code CONTINUE}
	 * labeled alternative in {@link GoshParser#statement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitCONTINUE(GoshParser.CONTINUEContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#assignment}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitAssignment(GoshParser.AssignmentContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#lvalue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLvalue(GoshParser.LvalueContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#rvalue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRvalue(GoshParser.RvalueContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#constvalue}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitConstvalue(GoshParser.ConstvalueContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#functionCall}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionCall(GoshParser.FunctionCallContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#arguments}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitArguments(GoshParser.ArgumentsContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#loop}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitLoop(GoshParser.LoopContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#expression}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitExpression(GoshParser.ExpressionContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#block}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBlock(GoshParser.BlockContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#outerAccess}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitOuterAccess(GoshParser.OuterAccessContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#functionDefinition}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitFunctionDefinition(GoshParser.FunctionDefinitionContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#parameters}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitParameters(GoshParser.ParametersContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#runStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitRunStatement(GoshParser.RunStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#returnStatement}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitReturnStatement(GoshParser.ReturnStatementContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#ifStmt}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitIfStmt(GoshParser.IfStmtContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#breakStmt}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBreakStmt(GoshParser.BreakStmtContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#continueStmt}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitContinueStmt(GoshParser.ContinueStmtContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#simpleStmt}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitSimpleStmt(GoshParser.SimpleStmtContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#forClause}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitForClause(GoshParser.ForClauseContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#mulDivOP}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitMulDivOP(GoshParser.MulDivOPContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#binOP}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitBinOP(GoshParser.BinOPContext ctx);
	/**
	 * Visit a parse tree produced by {@link GoshParser#unOP}.
	 * @param ctx the parse tree
	 * @return the visitor result
	 */
	T visitUnOP(GoshParser.UnOPContext ctx);
}