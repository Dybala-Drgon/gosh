grammar Gosh;

program: statements;

statements: statement*;

statement: assignment #ASSIGN|functionCall #FUNCCALL| expression #EXP| loop#FORLOOP|block#BLOCK
| functionDefinition #FUNCDEF| runStatement #RUN|returnStatement#RETURN|ifStmt#IF| breakStmt#BREAK| continueStmt #CONTINUE;

assignment: lvalue ASSIGN rvalue;
lvalue: ID (COMMA ID)*;
rvalue: functionCall|(expression | constvalue)(COMMA rvalue)*;
constvalue: Number|Str;

functionCall: ID L_PAREN arguments? R_PAREN;

arguments: expression (COMMA expression)*;

loop: FOR (forClause ) block;

expression: ID | Number | unOP expression | L_PAREN expression R_PAREN | expression mulDivOP expression |expression binOP expression | outerAccess ;

block: L_CURLY statements R_CURLY;

outerAccess: OUTER ID; // 外部作用域变量引用规则

functionDefinition: FUNC ID L_PAREN parameters? R_PAREN block;

parameters: ID (COMMA ID)*;

runStatement: RUN block;

returnStatement: RETURN expression (COMMA expression)*;

ifStmt:
	IF (expression) block (
		ELSE (ifStmt | block)
	)?;

// 词法规则
COMMA                  : ',';
ASSIGN                 :'=';
FUNC                   :'func';
WS                     : [ \t]+             -> channel(HIDDEN);
COMMENT                : '/*' .*? '*/'      -> channel(HIDDEN);
TERMINATOR             : [\r\n]+            -> channel(HIDDEN);
LINE_COMMENT           : '//' ~[\r\n]*      -> channel(HIDDEN);
BREAK                  : 'break';
CONTINUE               :'continue';
FOR                    : 'for';
EOS                    :';';
OUTER                  :'outer.';
L_PAREN                : '(';
R_PAREN                : ')';
L_CURLY                : '{';
R_CURLY                : '}';
RUN                    : 'run';
RETURN                 :'return';
ELSE                   : 'else';
IF                     : 'if';
ADD                    :'+';
SUB                    :'-';

breakStmt              :BREAK ;
continueStmt           :CONTINUE ;
simpleStmt             :assignment|expression;
forClause              :initStmt = simpleStmt? EOS expression? EOS postStmt = simpleStmt?;
Number               : [0-9] '.'? [0-9]*;
ID: [a-zA-Z]+('_'|Number)*;
Str:'"'[a-zA-Z]+'"';
//INT: [0-9]+;

mulDivOP :'*' | '/' ;
binOP : ADD | SUB | '//' | '%' | '&' | '|'
        | '^' | '>>' | '<<' | '<=' | '>=' | '<' | '>'
        | '==' | '!=' | '&&' | '||';

unOP : '~' | '!' | '-' | '--' | '++';