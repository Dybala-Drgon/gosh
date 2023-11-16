grammar Calc; // 定义规则名称

// Tokens 元素
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules // 入口规则
start : expr EOF; // 这里依赖expr规则

expr  // 具体的expr规则， 这里是个递归，可以自己调用自己，这里有先后顺序， 
   : expr op=('*'|'/') expr # MulDiv // 乘法
   | expr op=('+'|'-') expr # AddSub // 加法
   | NUMBER                 # Number // 数字
   ;
