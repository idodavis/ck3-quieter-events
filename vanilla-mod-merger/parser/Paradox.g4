grammar Paradox;

file: (statement | COMMENT | WS)* EOF;

statement:
	assignment
	| macroDefinition
	| macroInsert
	| variableDefinition
	| calculation
	| valueOnly;

assignment: key comparator value;

macroDefinition: DEFINE key comparator blockBody;

macroInsert: INSERT key comparator blockBody;

variableDefinition:
	VAR comparator value
	| REGISTER_VARIABLE key comparator value;

calculation: CALC_START expr CALC_END;

blockBody: LBRACE blockContents RBRACE;

blockContents: (statement | COMMENT | WS)*;

valueOnly: value;

key: keyPart (WS+ keyPart)* # MultiKey;

keyPart: ID (COLON ID)? # ColonKeyPart;

value: NUMBER | STRING | ID | blockBody | calculation;

comparator: EQUALS | NOT_EQUALS | GTE | LTE | GT | LT;

expr:
	expr ADD expr
	| expr SUB expr
	| expr MUL expr
	| expr DIV expr
	| SUB expr
	| NUMBER
	| LPAREN expr RPAREN
	| ID;

ID: [a-zA-Z0-9_.$:]+;
NUMBER: '-'? [0-9]+ ('.' [0-9]+)?;
STRING: '"' (~["\r\n])* '"';

DEFINE: '@:define';
INSERT: '@:insert';
REGISTER_VARIABLE: '@:register_variable';
VAR: '@' [a-zA-Z0-9_]+;
CALC_START: '@[';
CALC_END: ']';

EQUALS: '=';
NOT_EQUALS: '!=';
GTE: '>=';
LTE: '<=';
GT: '>';
LT: '<';

LBRACE: '{';
RBRACE: '}';
LPAREN: '(';
RPAREN: ')';
ADD: '+';
SUB: '-';
MUL: '*';
DIV: '/';
COLON: ':';

COMMENT: '#' ~[\r\n]* -> skip;
WS: [ \t\r\n]+ -> skip;