// Code generated from .//Gosh.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Gosh

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type GoshParser struct {
	*antlr.BaseParser
}

var GoshParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func goshParserInit() {
	staticData := &GoshParserStaticData
	staticData.LiteralNames = []string{
		"", "'*'", "'/'", "'//'", "'%'", "'&'", "'|'", "'^'", "'>>'", "'<<'",
		"'~'", "'!'", "'--'", "'++'", "','", "'='", "'func'", "", "", "", "",
		"'break'", "'continue'", "'for'", "';'", "'outer.'", "'('", "')'", "'{'",
		"'}'", "'run'", "'return'", "'else'", "'if'", "'+'", "'-'", "'=='",
		"'!='", "'&&'", "'||'", "'<'", "'<='", "'>'", "'>='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "COMMA", "ASSIGN",
		"FUNC", "WS", "COMMENT", "TERMINATOR", "LINE_COMMENT", "BREAK", "CONTINUE",
		"FOR", "EOS", "OUTER", "L_PAREN", "R_PAREN", "L_CURLY", "R_CURLY", "RUN",
		"RETURN", "ELSE", "IF", "ADD", "SUB", "EQUAL", "NOTEQUAL", "AND", "OR",
		"LESS", "LESS_EQUAL", "GREATER", "GREATER_EQUAL", "Number", "ID", "Str",
	}
	staticData.RuleNames = []string{
		"program", "statements", "statement", "assignment", "lvalue", "rvalue",
		"constvalue", "functionCall", "arguments", "loop", "expression", "block",
		"outerAccess", "functionDefinition", "parameters", "runStatement", "returnStatement",
		"ifStmt", "breakStmt", "continueStmt", "simpleStmt", "forClause", "mulDivOP",
		"binOP", "unOP",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 46, 216, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 1,
		5, 1, 54, 8, 1, 10, 1, 12, 1, 57, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 70, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 1, 4, 5, 4, 79, 8, 4, 10, 4, 12, 4, 82, 9, 4, 1, 5, 1, 5, 1,
		5, 3, 5, 87, 8, 5, 1, 5, 1, 5, 5, 5, 91, 8, 5, 10, 5, 12, 5, 94, 9, 5,
		3, 5, 96, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 103, 8, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 5, 8, 110, 8, 8, 10, 8, 12, 8, 113, 9, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 1, 10, 1, 10, 3, 10, 130, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 10, 1, 10, 5, 10, 140, 8, 10, 10, 10, 12, 10, 143, 9, 10,
		1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 3, 13, 156, 8, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 5, 14,
		164, 8, 14, 10, 14, 12, 14, 167, 9, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 16, 1, 16, 5, 16, 176, 8, 16, 10, 16, 12, 16, 179, 9, 16, 1, 17,
		1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 187, 8, 17, 3, 17, 189, 8, 17,
		1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 3, 20, 197, 8, 20, 1, 21, 3,
		21, 200, 8, 21, 1, 21, 1, 21, 3, 21, 204, 8, 21, 1, 21, 1, 21, 3, 21, 208,
		8, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 0, 1, 20, 25, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 0, 4, 2, 0, 44, 44, 46, 46, 1, 0, 1, 2, 2, 0, 3, 9,
		34, 43, 2, 0, 10, 13, 35, 35, 222, 0, 50, 1, 0, 0, 0, 2, 55, 1, 0, 0, 0,
		4, 69, 1, 0, 0, 0, 6, 71, 1, 0, 0, 0, 8, 75, 1, 0, 0, 0, 10, 95, 1, 0,
		0, 0, 12, 97, 1, 0, 0, 0, 14, 99, 1, 0, 0, 0, 16, 106, 1, 0, 0, 0, 18,
		114, 1, 0, 0, 0, 20, 129, 1, 0, 0, 0, 22, 144, 1, 0, 0, 0, 24, 148, 1,
		0, 0, 0, 26, 151, 1, 0, 0, 0, 28, 160, 1, 0, 0, 0, 30, 168, 1, 0, 0, 0,
		32, 171, 1, 0, 0, 0, 34, 180, 1, 0, 0, 0, 36, 190, 1, 0, 0, 0, 38, 192,
		1, 0, 0, 0, 40, 196, 1, 0, 0, 0, 42, 199, 1, 0, 0, 0, 44, 209, 1, 0, 0,
		0, 46, 211, 1, 0, 0, 0, 48, 213, 1, 0, 0, 0, 50, 51, 3, 2, 1, 0, 51, 1,
		1, 0, 0, 0, 52, 54, 3, 4, 2, 0, 53, 52, 1, 0, 0, 0, 54, 57, 1, 0, 0, 0,
		55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 3, 1, 0, 0, 0, 57, 55, 1, 0,
		0, 0, 58, 70, 3, 6, 3, 0, 59, 70, 3, 14, 7, 0, 60, 70, 3, 20, 10, 0, 61,
		70, 3, 18, 9, 0, 62, 70, 3, 22, 11, 0, 63, 70, 3, 26, 13, 0, 64, 70, 3,
		30, 15, 0, 65, 70, 3, 32, 16, 0, 66, 70, 3, 34, 17, 0, 67, 70, 3, 36, 18,
		0, 68, 70, 3, 38, 19, 0, 69, 58, 1, 0, 0, 0, 69, 59, 1, 0, 0, 0, 69, 60,
		1, 0, 0, 0, 69, 61, 1, 0, 0, 0, 69, 62, 1, 0, 0, 0, 69, 63, 1, 0, 0, 0,
		69, 64, 1, 0, 0, 0, 69, 65, 1, 0, 0, 0, 69, 66, 1, 0, 0, 0, 69, 67, 1,
		0, 0, 0, 69, 68, 1, 0, 0, 0, 70, 5, 1, 0, 0, 0, 71, 72, 3, 8, 4, 0, 72,
		73, 5, 15, 0, 0, 73, 74, 3, 10, 5, 0, 74, 7, 1, 0, 0, 0, 75, 80, 5, 45,
		0, 0, 76, 77, 5, 14, 0, 0, 77, 79, 5, 45, 0, 0, 78, 76, 1, 0, 0, 0, 79,
		82, 1, 0, 0, 0, 80, 78, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 9, 1, 0, 0,
		0, 82, 80, 1, 0, 0, 0, 83, 96, 3, 14, 7, 0, 84, 87, 3, 20, 10, 0, 85, 87,
		3, 12, 6, 0, 86, 84, 1, 0, 0, 0, 86, 85, 1, 0, 0, 0, 87, 92, 1, 0, 0, 0,
		88, 89, 5, 14, 0, 0, 89, 91, 3, 10, 5, 0, 90, 88, 1, 0, 0, 0, 91, 94, 1,
		0, 0, 0, 92, 90, 1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94,
		92, 1, 0, 0, 0, 95, 83, 1, 0, 0, 0, 95, 86, 1, 0, 0, 0, 96, 11, 1, 0, 0,
		0, 97, 98, 7, 0, 0, 0, 98, 13, 1, 0, 0, 0, 99, 100, 5, 45, 0, 0, 100, 102,
		5, 26, 0, 0, 101, 103, 3, 16, 8, 0, 102, 101, 1, 0, 0, 0, 102, 103, 1,
		0, 0, 0, 103, 104, 1, 0, 0, 0, 104, 105, 5, 27, 0, 0, 105, 15, 1, 0, 0,
		0, 106, 111, 3, 20, 10, 0, 107, 108, 5, 14, 0, 0, 108, 110, 3, 20, 10,
		0, 109, 107, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109, 1, 0, 0, 0, 111,
		112, 1, 0, 0, 0, 112, 17, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0, 114, 115, 5,
		23, 0, 0, 115, 116, 3, 42, 21, 0, 116, 117, 3, 22, 11, 0, 117, 19, 1, 0,
		0, 0, 118, 119, 6, 10, -1, 0, 119, 130, 5, 45, 0, 0, 120, 130, 5, 44, 0,
		0, 121, 122, 3, 48, 24, 0, 122, 123, 3, 20, 10, 5, 123, 130, 1, 0, 0, 0,
		124, 125, 5, 26, 0, 0, 125, 126, 3, 20, 10, 0, 126, 127, 5, 27, 0, 0, 127,
		130, 1, 0, 0, 0, 128, 130, 3, 24, 12, 0, 129, 118, 1, 0, 0, 0, 129, 120,
		1, 0, 0, 0, 129, 121, 1, 0, 0, 0, 129, 124, 1, 0, 0, 0, 129, 128, 1, 0,
		0, 0, 130, 141, 1, 0, 0, 0, 131, 132, 10, 3, 0, 0, 132, 133, 3, 44, 22,
		0, 133, 134, 3, 20, 10, 4, 134, 140, 1, 0, 0, 0, 135, 136, 10, 2, 0, 0,
		136, 137, 3, 46, 23, 0, 137, 138, 3, 20, 10, 3, 138, 140, 1, 0, 0, 0, 139,
		131, 1, 0, 0, 0, 139, 135, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 139,
		1, 0, 0, 0, 141, 142, 1, 0, 0, 0, 142, 21, 1, 0, 0, 0, 143, 141, 1, 0,
		0, 0, 144, 145, 5, 28, 0, 0, 145, 146, 3, 2, 1, 0, 146, 147, 5, 29, 0,
		0, 147, 23, 1, 0, 0, 0, 148, 149, 5, 25, 0, 0, 149, 150, 5, 45, 0, 0, 150,
		25, 1, 0, 0, 0, 151, 152, 5, 16, 0, 0, 152, 153, 5, 45, 0, 0, 153, 155,
		5, 26, 0, 0, 154, 156, 3, 28, 14, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1,
		0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 158, 5, 27, 0, 0, 158, 159, 3, 22,
		11, 0, 159, 27, 1, 0, 0, 0, 160, 165, 5, 45, 0, 0, 161, 162, 5, 14, 0,
		0, 162, 164, 5, 45, 0, 0, 163, 161, 1, 0, 0, 0, 164, 167, 1, 0, 0, 0, 165,
		163, 1, 0, 0, 0, 165, 166, 1, 0, 0, 0, 166, 29, 1, 0, 0, 0, 167, 165, 1,
		0, 0, 0, 168, 169, 5, 30, 0, 0, 169, 170, 3, 22, 11, 0, 170, 31, 1, 0,
		0, 0, 171, 172, 5, 31, 0, 0, 172, 177, 3, 20, 10, 0, 173, 174, 5, 14, 0,
		0, 174, 176, 3, 20, 10, 0, 175, 173, 1, 0, 0, 0, 176, 179, 1, 0, 0, 0,
		177, 175, 1, 0, 0, 0, 177, 178, 1, 0, 0, 0, 178, 33, 1, 0, 0, 0, 179, 177,
		1, 0, 0, 0, 180, 181, 5, 33, 0, 0, 181, 182, 3, 20, 10, 0, 182, 188, 3,
		22, 11, 0, 183, 186, 5, 32, 0, 0, 184, 187, 3, 34, 17, 0, 185, 187, 3,
		22, 11, 0, 186, 184, 1, 0, 0, 0, 186, 185, 1, 0, 0, 0, 187, 189, 1, 0,
		0, 0, 188, 183, 1, 0, 0, 0, 188, 189, 1, 0, 0, 0, 189, 35, 1, 0, 0, 0,
		190, 191, 5, 21, 0, 0, 191, 37, 1, 0, 0, 0, 192, 193, 5, 22, 0, 0, 193,
		39, 1, 0, 0, 0, 194, 197, 3, 6, 3, 0, 195, 197, 3, 20, 10, 0, 196, 194,
		1, 0, 0, 0, 196, 195, 1, 0, 0, 0, 197, 41, 1, 0, 0, 0, 198, 200, 3, 40,
		20, 0, 199, 198, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0,
		201, 203, 5, 24, 0, 0, 202, 204, 3, 20, 10, 0, 203, 202, 1, 0, 0, 0, 203,
		204, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 207, 5, 24, 0, 0, 206, 208,
		3, 40, 20, 0, 207, 206, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 43, 1, 0,
		0, 0, 209, 210, 7, 1, 0, 0, 210, 45, 1, 0, 0, 0, 211, 212, 7, 2, 0, 0,
		212, 47, 1, 0, 0, 0, 213, 214, 7, 3, 0, 0, 214, 49, 1, 0, 0, 0, 20, 55,
		69, 80, 86, 92, 95, 102, 111, 129, 139, 141, 155, 165, 177, 186, 188, 196,
		199, 203, 207,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// GoshParserInit initializes any static state used to implement GoshParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGoshParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GoshParserInit() {
	staticData := &GoshParserStaticData
	staticData.once.Do(goshParserInit)
}

// NewGoshParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGoshParser(input antlr.TokenStream) *GoshParser {
	GoshParserInit()
	this := new(GoshParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GoshParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Gosh.g4"

	return this
}

// GoshParser tokens.
const (
	GoshParserEOF           = antlr.TokenEOF
	GoshParserT__0          = 1
	GoshParserT__1          = 2
	GoshParserT__2          = 3
	GoshParserT__3          = 4
	GoshParserT__4          = 5
	GoshParserT__5          = 6
	GoshParserT__6          = 7
	GoshParserT__7          = 8
	GoshParserT__8          = 9
	GoshParserT__9          = 10
	GoshParserT__10         = 11
	GoshParserT__11         = 12
	GoshParserT__12         = 13
	GoshParserCOMMA         = 14
	GoshParserASSIGN        = 15
	GoshParserFUNC          = 16
	GoshParserWS            = 17
	GoshParserCOMMENT       = 18
	GoshParserTERMINATOR    = 19
	GoshParserLINE_COMMENT  = 20
	GoshParserBREAK         = 21
	GoshParserCONTINUE      = 22
	GoshParserFOR           = 23
	GoshParserEOS           = 24
	GoshParserOUTER         = 25
	GoshParserL_PAREN       = 26
	GoshParserR_PAREN       = 27
	GoshParserL_CURLY       = 28
	GoshParserR_CURLY       = 29
	GoshParserRUN           = 30
	GoshParserRETURN        = 31
	GoshParserELSE          = 32
	GoshParserIF            = 33
	GoshParserADD           = 34
	GoshParserSUB           = 35
	GoshParserEQUAL         = 36
	GoshParserNOTEQUAL      = 37
	GoshParserAND           = 38
	GoshParserOR            = 39
	GoshParserLESS          = 40
	GoshParserLESS_EQUAL    = 41
	GoshParserGREATER       = 42
	GoshParserGREATER_EQUAL = 43
	GoshParserNumber        = 44
	GoshParserID            = 45
	GoshParserStr           = 46
)

// GoshParser rules.
const (
	GoshParserRULE_program            = 0
	GoshParserRULE_statements         = 1
	GoshParserRULE_statement          = 2
	GoshParserRULE_assignment         = 3
	GoshParserRULE_lvalue             = 4
	GoshParserRULE_rvalue             = 5
	GoshParserRULE_constvalue         = 6
	GoshParserRULE_functionCall       = 7
	GoshParserRULE_arguments          = 8
	GoshParserRULE_loop               = 9
	GoshParserRULE_expression         = 10
	GoshParserRULE_block              = 11
	GoshParserRULE_outerAccess        = 12
	GoshParserRULE_functionDefinition = 13
	GoshParserRULE_parameters         = 14
	GoshParserRULE_runStatement       = 15
	GoshParserRULE_returnStatement    = 16
	GoshParserRULE_ifStmt             = 17
	GoshParserRULE_breakStmt          = 18
	GoshParserRULE_continueStmt       = 19
	GoshParserRULE_simpleStmt         = 20
	GoshParserRULE_forClause          = 21
	GoshParserRULE_mulDivOP           = 22
	GoshParserRULE_binOP              = 23
	GoshParserRULE_unOP               = 24
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Statements() IStatementsContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GoshParserRULE_program)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Statements()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementsContext is an interface to support dynamic dispatch.
type IStatementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsStatementsContext differentiates from other interfaces.
	IsStatementsContext()
}

type StatementsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementsContext() *StatementsContext {
	var p = new(StatementsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_statements
	return p
}

func InitEmptyStatementsContext(p *StatementsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_statements
}

func (*StatementsContext) IsStatementsContext() {}

func NewStatementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementsContext {
	var p = new(StatementsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_statements

	return p
}

func (s *StatementsContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementsContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *StatementsContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *StatementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterStatements(s)
	}
}

func (s *StatementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitStatements(s)
	}
}

func (s *StatementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitStatements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) Statements() (localctx IStatementsContext) {
	localctx = NewStatementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GoshParserRULE_statements)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&52823112891392) != 0 {
		{
			p.SetState(52)
			p.Statement()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyAll(ctx *StatementContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CONTINUEContext struct {
	StatementContext
}

func NewCONTINUEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CONTINUEContext {
	var p = new(CONTINUEContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *CONTINUEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CONTINUEContext) ContinueStmt() IContinueStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IContinueStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IContinueStmtContext)
}

func (s *CONTINUEContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterCONTINUE(s)
	}
}

func (s *CONTINUEContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitCONTINUE(s)
	}
}

func (s *CONTINUEContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitCONTINUE(s)

	default:
		return t.VisitChildren(s)
	}
}

type RETURNContext struct {
	StatementContext
}

func NewRETURNContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RETURNContext {
	var p = new(RETURNContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *RETURNContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RETURNContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *RETURNContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterRETURN(s)
	}
}

func (s *RETURNContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitRETURN(s)
	}
}

func (s *RETURNContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitRETURN(s)

	default:
		return t.VisitChildren(s)
	}
}

type FORLOOPContext struct {
	StatementContext
}

func NewFORLOOPContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FORLOOPContext {
	var p = new(FORLOOPContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *FORLOOPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FORLOOPContext) Loop() ILoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILoopContext)
}

func (s *FORLOOPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterFORLOOP(s)
	}
}

func (s *FORLOOPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitFORLOOP(s)
	}
}

func (s *FORLOOPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitFORLOOP(s)

	default:
		return t.VisitChildren(s)
	}
}

type BREAKContext struct {
	StatementContext
}

func NewBREAKContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BREAKContext {
	var p = new(BREAKContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BREAKContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BREAKContext) BreakStmt() IBreakStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBreakStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBreakStmtContext)
}

func (s *BREAKContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterBREAK(s)
	}
}

func (s *BREAKContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitBREAK(s)
	}
}

func (s *BREAKContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitBREAK(s)

	default:
		return t.VisitChildren(s)
	}
}

type BLOCKContext struct {
	StatementContext
}

func NewBLOCKContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BLOCKContext {
	var p = new(BLOCKContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *BLOCKContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BLOCKContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *BLOCKContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterBLOCK(s)
	}
}

func (s *BLOCKContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitBLOCK(s)
	}
}

func (s *BLOCKContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitBLOCK(s)

	default:
		return t.VisitChildren(s)
	}
}

type RUNContext struct {
	StatementContext
}

func NewRUNContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RUNContext {
	var p = new(RUNContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *RUNContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RUNContext) RunStatement() IRunStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRunStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRunStatementContext)
}

func (s *RUNContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterRUN(s)
	}
}

func (s *RUNContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitRUN(s)
	}
}

func (s *RUNContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitRUN(s)

	default:
		return t.VisitChildren(s)
	}
}

type FUNCDEFContext struct {
	StatementContext
}

func NewFUNCDEFContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FUNCDEFContext {
	var p = new(FUNCDEFContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *FUNCDEFContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FUNCDEFContext) FunctionDefinition() IFunctionDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefinitionContext)
}

func (s *FUNCDEFContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterFUNCDEF(s)
	}
}

func (s *FUNCDEFContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitFUNCDEF(s)
	}
}

func (s *FUNCDEFContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitFUNCDEF(s)

	default:
		return t.VisitChildren(s)
	}
}

type ASSIGNContext struct {
	StatementContext
}

func NewASSIGNContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ASSIGNContext {
	var p = new(ASSIGNContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *ASSIGNContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ASSIGNContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *ASSIGNContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterASSIGN(s)
	}
}

func (s *ASSIGNContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitASSIGN(s)
	}
}

func (s *ASSIGNContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitASSIGN(s)

	default:
		return t.VisitChildren(s)
	}
}

type EXPContext struct {
	StatementContext
}

func NewEXPContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EXPContext {
	var p = new(EXPContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *EXPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EXPContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *EXPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterEXP(s)
	}
}

func (s *EXPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitEXP(s)
	}
}

func (s *EXPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitEXP(s)

	default:
		return t.VisitChildren(s)
	}
}

type IFContext struct {
	StatementContext
}

func NewIFContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IFContext {
	var p = new(IFContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *IFContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IFContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *IFContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterIF(s)
	}
}

func (s *IFContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitIF(s)
	}
}

func (s *IFContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitIF(s)

	default:
		return t.VisitChildren(s)
	}
}

type FUNCCALLContext struct {
	StatementContext
}

func NewFUNCCALLContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FUNCCALLContext {
	var p = new(FUNCCALLContext)

	InitEmptyStatementContext(&p.StatementContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatementContext))

	return p
}

func (s *FUNCCALLContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FUNCCALLContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FUNCCALLContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterFUNCCALL(s)
	}
}

func (s *FUNCCALLContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitFUNCCALL(s)
	}
}

func (s *FUNCCALLContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitFUNCCALL(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GoshParserRULE_statement)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewASSIGNContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(58)
			p.Assignment()
		}

	case 2:
		localctx = NewFUNCCALLContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(59)
			p.FunctionCall()
		}

	case 3:
		localctx = NewEXPContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)
			p.expression(0)
		}

	case 4:
		localctx = NewFORLOOPContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(61)
			p.Loop()
		}

	case 5:
		localctx = NewBLOCKContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(62)
			p.Block()
		}

	case 6:
		localctx = NewFUNCDEFContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(63)
			p.FunctionDefinition()
		}

	case 7:
		localctx = NewRUNContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(64)
			p.RunStatement()
		}

	case 8:
		localctx = NewRETURNContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(65)
			p.ReturnStatement()
		}

	case 9:
		localctx = NewIFContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(66)
			p.IfStmt()
		}

	case 10:
		localctx = NewBREAKContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(67)
			p.BreakStmt()
		}

	case 11:
		localctx = NewCONTINUEContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(68)
			p.ContinueStmt()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Lvalue() ILvalueContext
	ASSIGN() antlr.TerminalNode
	Rvalue() IRvalueContext

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Lvalue() ILvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILvalueContext)
}

func (s *AssignmentContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(GoshParserASSIGN, 0)
}

func (s *AssignmentContext) Rvalue() IRvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (s *AssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GoshParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Lvalue()
	}
	{
		p.SetState(72)
		p.Match(GoshParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(73)
		p.Rvalue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILvalueContext is an interface to support dynamic dispatch.
type ILvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsLvalueContext differentiates from other interfaces.
	IsLvalueContext()
}

type LvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLvalueContext() *LvalueContext {
	var p = new(LvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_lvalue
	return p
}

func InitEmptyLvalueContext(p *LvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_lvalue
}

func (*LvalueContext) IsLvalueContext() {}

func NewLvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LvalueContext {
	var p = new(LvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_lvalue

	return p
}

func (s *LvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *LvalueContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GoshParserID)
}

func (s *LvalueContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GoshParserID, i)
}

func (s *LvalueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoshParserCOMMA)
}

func (s *LvalueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoshParserCOMMA, i)
}

func (s *LvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterLvalue(s)
	}
}

func (s *LvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitLvalue(s)
	}
}

func (s *LvalueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitLvalue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) Lvalue() (localctx ILvalueContext) {
	localctx = NewLvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GoshParserRULE_lvalue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.Match(GoshParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoshParserCOMMA {
		{
			p.SetState(76)
			p.Match(GoshParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.Match(GoshParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRvalueContext is an interface to support dynamic dispatch.
type IRvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FunctionCall() IFunctionCallContext
	Expression() IExpressionContext
	Constvalue() IConstvalueContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	AllRvalue() []IRvalueContext
	Rvalue(i int) IRvalueContext

	// IsRvalueContext differentiates from other interfaces.
	IsRvalueContext()
}

type RvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRvalueContext() *RvalueContext {
	var p = new(RvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_rvalue
	return p
}

func InitEmptyRvalueContext(p *RvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_rvalue
}

func (*RvalueContext) IsRvalueContext() {}

func NewRvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RvalueContext {
	var p = new(RvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_rvalue

	return p
}

func (s *RvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *RvalueContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *RvalueContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RvalueContext) Constvalue() IConstvalueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstvalueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstvalueContext)
}

func (s *RvalueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoshParserCOMMA)
}

func (s *RvalueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoshParserCOMMA, i)
}

func (s *RvalueContext) AllRvalue() []IRvalueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRvalueContext); ok {
			len++
		}
	}

	tst := make([]IRvalueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRvalueContext); ok {
			tst[i] = t.(IRvalueContext)
			i++
		}
	}

	return tst
}

func (s *RvalueContext) Rvalue(i int) IRvalueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRvalueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRvalueContext)
}

func (s *RvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterRvalue(s)
	}
}

func (s *RvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitRvalue(s)
	}
}

func (s *RvalueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitRvalue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) Rvalue() (localctx IRvalueContext) {
	localctx = NewRvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GoshParserRULE_rvalue)
	var _alt int

	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.FunctionCall()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(86)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(84)
				p.expression(0)
			}

		case 2:
			{
				p.SetState(85)
				p.Constvalue()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(88)
					p.Match(GoshParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(89)
					p.Rvalue()
				}

			}
			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstvalueContext is an interface to support dynamic dispatch.
type IConstvalueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Number() antlr.TerminalNode
	Str() antlr.TerminalNode

	// IsConstvalueContext differentiates from other interfaces.
	IsConstvalueContext()
}

type ConstvalueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstvalueContext() *ConstvalueContext {
	var p = new(ConstvalueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_constvalue
	return p
}

func InitEmptyConstvalueContext(p *ConstvalueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_constvalue
}

func (*ConstvalueContext) IsConstvalueContext() {}

func NewConstvalueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstvalueContext {
	var p = new(ConstvalueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_constvalue

	return p
}

func (s *ConstvalueContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstvalueContext) Number() antlr.TerminalNode {
	return s.GetToken(GoshParserNumber, 0)
}

func (s *ConstvalueContext) Str() antlr.TerminalNode {
	return s.GetToken(GoshParserStr, 0)
}

func (s *ConstvalueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstvalueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstvalueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterConstvalue(s)
	}
}

func (s *ConstvalueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitConstvalue(s)
	}
}

func (s *ConstvalueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitConstvalue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) Constvalue() (localctx IConstvalueContext) {
	localctx = NewConstvalueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GoshParserRULE_constvalue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoshParserNumber || _la == GoshParserStr) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	Arguments() IArgumentsContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(GoshParserID, 0)
}

func (s *FunctionCallContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GoshParserL_PAREN, 0)
}

func (s *FunctionCallContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GoshParserR_PAREN, 0)
}

func (s *FunctionCallContext) Arguments() IArgumentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GoshParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(99)
		p.Match(GoshParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(100)
		p.Match(GoshParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(102)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&52811018550272) != 0 {
		{
			p.SetState(101)
			p.Arguments()
		}

	}
	{
		p.SetState(104)
		p.Match(GoshParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_arguments
	return p
}

func InitEmptyArgumentsContext(p *ArgumentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_arguments
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentsContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoshParserCOMMA)
}

func (s *ArgumentsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoshParserCOMMA, i)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GoshParserRULE_arguments)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.expression(0)
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoshParserCOMMA {
		{
			p.SetState(107)
			p.Match(GoshParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.expression(0)
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILoopContext is an interface to support dynamic dispatch.
type ILoopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FOR() antlr.TerminalNode
	Block() IBlockContext
	ForClause() IForClauseContext

	// IsLoopContext differentiates from other interfaces.
	IsLoopContext()
}

type LoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLoopContext() *LoopContext {
	var p = new(LoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_loop
	return p
}

func InitEmptyLoopContext(p *LoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_loop
}

func (*LoopContext) IsLoopContext() {}

func NewLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopContext {
	var p = new(LoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_loop

	return p
}

func (s *LoopContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(GoshParserFOR, 0)
}

func (s *LoopContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *LoopContext) ForClause() IForClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForClauseContext)
}

func (s *LoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterLoop(s)
	}
}

func (s *LoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitLoop(s)
	}
}

func (s *LoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) Loop() (localctx ILoopContext) {
	localctx = NewLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GoshParserRULE_loop)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(GoshParserFOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(115)
		p.ForClause()
	}

	{
		p.SetState(116)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Number() antlr.TerminalNode
	UnOP() IUnOPContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	OuterAccess() IOuterAccessContext
	MulDivOP() IMulDivOPContext
	BinOP() IBinOPContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ID() antlr.TerminalNode {
	return s.GetToken(GoshParserID, 0)
}

func (s *ExpressionContext) Number() antlr.TerminalNode {
	return s.GetToken(GoshParserNumber, 0)
}

func (s *ExpressionContext) UnOP() IUnOPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnOPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnOPContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GoshParserL_PAREN, 0)
}

func (s *ExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GoshParserR_PAREN, 0)
}

func (s *ExpressionContext) OuterAccess() IOuterAccessContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOuterAccessContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOuterAccessContext)
}

func (s *ExpressionContext) MulDivOP() IMulDivOPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMulDivOPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMulDivOPContext)
}

func (s *ExpressionContext) BinOP() IBinOPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBinOPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBinOPContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *GoshParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 20
	p.EnterRecursionRule(localctx, 20, GoshParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(129)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GoshParserID:
		{
			p.SetState(119)
			p.Match(GoshParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoshParserNumber:
		{
			p.SetState(120)
			p.Match(GoshParserNumber)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoshParserT__9, GoshParserT__10, GoshParserT__11, GoshParserT__12, GoshParserSUB:
		{
			p.SetState(121)
			p.UnOP()
		}
		{
			p.SetState(122)
			p.expression(5)
		}

	case GoshParserL_PAREN:
		{
			p.SetState(124)
			p.Match(GoshParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(125)
			p.expression(0)
		}
		{
			p.SetState(126)
			p.Match(GoshParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GoshParserOUTER:
		{
			p.SetState(128)
			p.OuterAccess()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(139)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GoshParserRULE_expression)
				p.SetState(131)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(132)
					p.MulDivOP()
				}
				{
					p.SetState(133)
					p.expression(4)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, GoshParserRULE_expression)
				p.SetState(135)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(136)
					p.BinOP()
				}
				{
					p.SetState(137)
					p.expression(3)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(143)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_CURLY() antlr.TerminalNode
	Statements() IStatementsContext
	R_CURLY() antlr.TerminalNode

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(GoshParserL_CURLY, 0)
}

func (s *BlockContext) Statements() IStatementsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementsContext)
}

func (s *BlockContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(GoshParserR_CURLY, 0)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GoshParserRULE_block)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(GoshParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Statements()
	}
	{
		p.SetState(146)
		p.Match(GoshParserR_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOuterAccessContext is an interface to support dynamic dispatch.
type IOuterAccessContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OUTER() antlr.TerminalNode
	ID() antlr.TerminalNode

	// IsOuterAccessContext differentiates from other interfaces.
	IsOuterAccessContext()
}

type OuterAccessContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOuterAccessContext() *OuterAccessContext {
	var p = new(OuterAccessContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_outerAccess
	return p
}

func InitEmptyOuterAccessContext(p *OuterAccessContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_outerAccess
}

func (*OuterAccessContext) IsOuterAccessContext() {}

func NewOuterAccessContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OuterAccessContext {
	var p = new(OuterAccessContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_outerAccess

	return p
}

func (s *OuterAccessContext) GetParser() antlr.Parser { return s.parser }

func (s *OuterAccessContext) OUTER() antlr.TerminalNode {
	return s.GetToken(GoshParserOUTER, 0)
}

func (s *OuterAccessContext) ID() antlr.TerminalNode {
	return s.GetToken(GoshParserID, 0)
}

func (s *OuterAccessContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OuterAccessContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OuterAccessContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterOuterAccess(s)
	}
}

func (s *OuterAccessContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitOuterAccess(s)
	}
}

func (s *OuterAccessContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitOuterAccess(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) OuterAccess() (localctx IOuterAccessContext) {
	localctx = NewOuterAccessContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GoshParserRULE_outerAccess)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.Match(GoshParserOUTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(149)
		p.Match(GoshParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDefinitionContext is an interface to support dynamic dispatch.
type IFunctionDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	ID() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	Block() IBlockContext
	Parameters() IParametersContext

	// IsFunctionDefinitionContext differentiates from other interfaces.
	IsFunctionDefinitionContext()
}

type FunctionDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefinitionContext() *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_functionDefinition
	return p
}

func InitEmptyFunctionDefinitionContext(p *FunctionDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_functionDefinition
}

func (*FunctionDefinitionContext) IsFunctionDefinitionContext() {}

func NewFunctionDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefinitionContext {
	var p = new(FunctionDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_functionDefinition

	return p
}

func (s *FunctionDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefinitionContext) FUNC() antlr.TerminalNode {
	return s.GetToken(GoshParserFUNC, 0)
}

func (s *FunctionDefinitionContext) ID() antlr.TerminalNode {
	return s.GetToken(GoshParserID, 0)
}

func (s *FunctionDefinitionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(GoshParserL_PAREN, 0)
}

func (s *FunctionDefinitionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(GoshParserR_PAREN, 0)
}

func (s *FunctionDefinitionContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDefinitionContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitFunctionDefinition(s)
	}
}

func (s *FunctionDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitFunctionDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) FunctionDefinition() (localctx IFunctionDefinitionContext) {
	localctx = NewFunctionDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GoshParserRULE_functionDefinition)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(GoshParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(152)
		p.Match(GoshParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(153)
		p.Match(GoshParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoshParserID {
		{
			p.SetState(154)
			p.Parameters()
		}

	}
	{
		p.SetState(157)
		p.Match(GoshParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(158)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllID() []antlr.TerminalNode
	ID(i int) antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_parameters
	return p
}

func InitEmptyParametersContext(p *ParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_parameters
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(GoshParserID)
}

func (s *ParametersContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(GoshParserID, i)
}

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoshParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoshParserCOMMA, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (s *ParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GoshParserRULE_parameters)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(GoshParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoshParserCOMMA {
		{
			p.SetState(161)
			p.Match(GoshParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(162)
			p.Match(GoshParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRunStatementContext is an interface to support dynamic dispatch.
type IRunStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RUN() antlr.TerminalNode
	Block() IBlockContext

	// IsRunStatementContext differentiates from other interfaces.
	IsRunStatementContext()
}

type RunStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRunStatementContext() *RunStatementContext {
	var p = new(RunStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_runStatement
	return p
}

func InitEmptyRunStatementContext(p *RunStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_runStatement
}

func (*RunStatementContext) IsRunStatementContext() {}

func NewRunStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RunStatementContext {
	var p = new(RunStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_runStatement

	return p
}

func (s *RunStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *RunStatementContext) RUN() antlr.TerminalNode {
	return s.GetToken(GoshParserRUN, 0)
}

func (s *RunStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *RunStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RunStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RunStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterRunStatement(s)
	}
}

func (s *RunStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitRunStatement(s)
	}
}

func (s *RunStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitRunStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) RunStatement() (localctx IRunStatementContext) {
	localctx = NewRunStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GoshParserRULE_runStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		p.Match(GoshParserRUN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(GoshParserRETURN, 0)
}

func (s *ReturnStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ReturnStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(GoshParserCOMMA)
}

func (s *ReturnStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(GoshParserCOMMA, i)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GoshParserRULE_returnStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(171)
		p.Match(GoshParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.expression(0)
	}
	p.SetState(177)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GoshParserCOMMA {
		{
			p.SetState(173)
			p.Match(GoshParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.expression(0)
		}

		p.SetState(179)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStmtContext is an interface to support dynamic dispatch.
type IIfStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	Expression() IExpressionContext
	ELSE() antlr.TerminalNode
	IfStmt() IIfStmtContext

	// IsIfStmtContext differentiates from other interfaces.
	IsIfStmtContext()
}

type IfStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStmtContext() *IfStmtContext {
	var p = new(IfStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_ifStmt
	return p
}

func InitEmptyIfStmtContext(p *IfStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_ifStmt
}

func (*IfStmtContext) IsIfStmtContext() {}

func NewIfStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStmtContext {
	var p = new(IfStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_ifStmt

	return p
}

func (s *IfStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStmtContext) IF() antlr.TerminalNode {
	return s.GetToken(GoshParserIF, 0)
}

func (s *IfStmtContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStmtContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStmtContext) ELSE() antlr.TerminalNode {
	return s.GetToken(GoshParserELSE, 0)
}

func (s *IfStmtContext) IfStmt() IIfStmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStmtContext)
}

func (s *IfStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterIfStmt(s)
	}
}

func (s *IfStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitIfStmt(s)
	}
}

func (s *IfStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitIfStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) IfStmt() (localctx IIfStmtContext) {
	localctx = NewIfStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GoshParserRULE_ifStmt)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(GoshParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(181)
		p.expression(0)
	}

	{
		p.SetState(182)
		p.Block()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == GoshParserELSE {
		{
			p.SetState(183)
			p.Match(GoshParserELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case GoshParserIF:
			{
				p.SetState(184)
				p.IfStmt()
			}

		case GoshParserL_CURLY:
			{
				p.SetState(185)
				p.Block()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBreakStmtContext is an interface to support dynamic dispatch.
type IBreakStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BREAK() antlr.TerminalNode

	// IsBreakStmtContext differentiates from other interfaces.
	IsBreakStmtContext()
}

type BreakStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBreakStmtContext() *BreakStmtContext {
	var p = new(BreakStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_breakStmt
	return p
}

func InitEmptyBreakStmtContext(p *BreakStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_breakStmt
}

func (*BreakStmtContext) IsBreakStmtContext() {}

func NewBreakStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreakStmtContext {
	var p = new(BreakStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_breakStmt

	return p
}

func (s *BreakStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *BreakStmtContext) BREAK() antlr.TerminalNode {
	return s.GetToken(GoshParserBREAK, 0)
}

func (s *BreakStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreakStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterBreakStmt(s)
	}
}

func (s *BreakStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitBreakStmt(s)
	}
}

func (s *BreakStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitBreakStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) BreakStmt() (localctx IBreakStmtContext) {
	localctx = NewBreakStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GoshParserRULE_breakStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.Match(GoshParserBREAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IContinueStmtContext is an interface to support dynamic dispatch.
type IContinueStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONTINUE() antlr.TerminalNode

	// IsContinueStmtContext differentiates from other interfaces.
	IsContinueStmtContext()
}

type ContinueStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyContinueStmtContext() *ContinueStmtContext {
	var p = new(ContinueStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_continueStmt
	return p
}

func InitEmptyContinueStmtContext(p *ContinueStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_continueStmt
}

func (*ContinueStmtContext) IsContinueStmtContext() {}

func NewContinueStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinueStmtContext {
	var p = new(ContinueStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_continueStmt

	return p
}

func (s *ContinueStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinueStmtContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(GoshParserCONTINUE, 0)
}

func (s *ContinueStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinueStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterContinueStmt(s)
	}
}

func (s *ContinueStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitContinueStmt(s)
	}
}

func (s *ContinueStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitContinueStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) ContinueStmt() (localctx IContinueStmtContext) {
	localctx = NewContinueStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GoshParserRULE_continueStmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(GoshParserCONTINUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISimpleStmtContext is an interface to support dynamic dispatch.
type ISimpleStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	Expression() IExpressionContext

	// IsSimpleStmtContext differentiates from other interfaces.
	IsSimpleStmtContext()
}

type SimpleStmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleStmtContext() *SimpleStmtContext {
	var p = new(SimpleStmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_simpleStmt
	return p
}

func InitEmptySimpleStmtContext(p *SimpleStmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_simpleStmt
}

func (*SimpleStmtContext) IsSimpleStmtContext() {}

func NewSimpleStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleStmtContext {
	var p = new(SimpleStmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_simpleStmt

	return p
}

func (s *SimpleStmtContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleStmtContext) Assignment() IAssignmentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *SimpleStmtContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SimpleStmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleStmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleStmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterSimpleStmt(s)
	}
}

func (s *SimpleStmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitSimpleStmt(s)
	}
}

func (s *SimpleStmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitSimpleStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) SimpleStmt() (localctx ISimpleStmtContext) {
	localctx = NewSimpleStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GoshParserRULE_simpleStmt)
	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(194)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(195)
			p.expression(0)
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForClauseContext is an interface to support dynamic dispatch.
type IForClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetInitStmt returns the initStmt rule contexts.
	GetInitStmt() ISimpleStmtContext

	// GetPostStmt returns the postStmt rule contexts.
	GetPostStmt() ISimpleStmtContext

	// SetInitStmt sets the initStmt rule contexts.
	SetInitStmt(ISimpleStmtContext)

	// SetPostStmt sets the postStmt rule contexts.
	SetPostStmt(ISimpleStmtContext)

	// Getter signatures
	AllEOS() []antlr.TerminalNode
	EOS(i int) antlr.TerminalNode
	Expression() IExpressionContext
	AllSimpleStmt() []ISimpleStmtContext
	SimpleStmt(i int) ISimpleStmtContext

	// IsForClauseContext differentiates from other interfaces.
	IsForClauseContext()
}

type ForClauseContext struct {
	antlr.BaseParserRuleContext
	parser   antlr.Parser
	initStmt ISimpleStmtContext
	postStmt ISimpleStmtContext
}

func NewEmptyForClauseContext() *ForClauseContext {
	var p = new(ForClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_forClause
	return p
}

func InitEmptyForClauseContext(p *ForClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_forClause
}

func (*ForClauseContext) IsForClauseContext() {}

func NewForClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForClauseContext {
	var p = new(ForClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_forClause

	return p
}

func (s *ForClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ForClauseContext) GetInitStmt() ISimpleStmtContext { return s.initStmt }

func (s *ForClauseContext) GetPostStmt() ISimpleStmtContext { return s.postStmt }

func (s *ForClauseContext) SetInitStmt(v ISimpleStmtContext) { s.initStmt = v }

func (s *ForClauseContext) SetPostStmt(v ISimpleStmtContext) { s.postStmt = v }

func (s *ForClauseContext) AllEOS() []antlr.TerminalNode {
	return s.GetTokens(GoshParserEOS)
}

func (s *ForClauseContext) EOS(i int) antlr.TerminalNode {
	return s.GetToken(GoshParserEOS, i)
}

func (s *ForClauseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForClauseContext) AllSimpleStmt() []ISimpleStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			len++
		}
	}

	tst := make([]ISimpleStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISimpleStmtContext); ok {
			tst[i] = t.(ISimpleStmtContext)
			i++
		}
	}

	return tst
}

func (s *ForClauseContext) SimpleStmt(i int) ISimpleStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISimpleStmtContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISimpleStmtContext)
}

func (s *ForClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterForClause(s)
	}
}

func (s *ForClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitForClause(s)
	}
}

func (s *ForClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitForClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) ForClause() (localctx IForClauseContext) {
	localctx = NewForClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GoshParserRULE_forClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(199)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&52811018550272) != 0 {
		{
			p.SetState(198)

			var _x = p.SimpleStmt()

			localctx.(*ForClauseContext).initStmt = _x
		}

	}
	{
		p.SetState(201)
		p.Match(GoshParserEOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&52811018550272) != 0 {
		{
			p.SetState(202)
			p.expression(0)
		}

	}
	{
		p.SetState(205)
		p.Match(GoshParserEOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(207)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&52811018550272) != 0 {
		{
			p.SetState(206)

			var _x = p.SimpleStmt()

			localctx.(*ForClauseContext).postStmt = _x
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMulDivOPContext is an interface to support dynamic dispatch.
type IMulDivOPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsMulDivOPContext differentiates from other interfaces.
	IsMulDivOPContext()
}

type MulDivOPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMulDivOPContext() *MulDivOPContext {
	var p = new(MulDivOPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_mulDivOP
	return p
}

func InitEmptyMulDivOPContext(p *MulDivOPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_mulDivOP
}

func (*MulDivOPContext) IsMulDivOPContext() {}

func NewMulDivOPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MulDivOPContext {
	var p = new(MulDivOPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_mulDivOP

	return p
}

func (s *MulDivOPContext) GetParser() antlr.Parser { return s.parser }
func (s *MulDivOPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivOPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MulDivOPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterMulDivOP(s)
	}
}

func (s *MulDivOPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitMulDivOP(s)
	}
}

func (s *MulDivOPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitMulDivOP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) MulDivOP() (localctx IMulDivOPContext) {
	localctx = NewMulDivOPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GoshParserRULE_mulDivOP)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(209)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GoshParserT__0 || _la == GoshParserT__1) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBinOPContext is an interface to support dynamic dispatch.
type IBinOPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ADD() antlr.TerminalNode
	SUB() antlr.TerminalNode
	LESS_EQUAL() antlr.TerminalNode
	GREATER_EQUAL() antlr.TerminalNode
	LESS() antlr.TerminalNode
	GREATER() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	NOTEQUAL() antlr.TerminalNode
	AND() antlr.TerminalNode
	OR() antlr.TerminalNode

	// IsBinOPContext differentiates from other interfaces.
	IsBinOPContext()
}

type BinOPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinOPContext() *BinOPContext {
	var p = new(BinOPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_binOP
	return p
}

func InitEmptyBinOPContext(p *BinOPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_binOP
}

func (*BinOPContext) IsBinOPContext() {}

func NewBinOPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinOPContext {
	var p = new(BinOPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_binOP

	return p
}

func (s *BinOPContext) GetParser() antlr.Parser { return s.parser }

func (s *BinOPContext) ADD() antlr.TerminalNode {
	return s.GetToken(GoshParserADD, 0)
}

func (s *BinOPContext) SUB() antlr.TerminalNode {
	return s.GetToken(GoshParserSUB, 0)
}

func (s *BinOPContext) LESS_EQUAL() antlr.TerminalNode {
	return s.GetToken(GoshParserLESS_EQUAL, 0)
}

func (s *BinOPContext) GREATER_EQUAL() antlr.TerminalNode {
	return s.GetToken(GoshParserGREATER_EQUAL, 0)
}

func (s *BinOPContext) LESS() antlr.TerminalNode {
	return s.GetToken(GoshParserLESS, 0)
}

func (s *BinOPContext) GREATER() antlr.TerminalNode {
	return s.GetToken(GoshParserGREATER, 0)
}

func (s *BinOPContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(GoshParserEQUAL, 0)
}

func (s *BinOPContext) NOTEQUAL() antlr.TerminalNode {
	return s.GetToken(GoshParserNOTEQUAL, 0)
}

func (s *BinOPContext) AND() antlr.TerminalNode {
	return s.GetToken(GoshParserAND, 0)
}

func (s *BinOPContext) OR() antlr.TerminalNode {
	return s.GetToken(GoshParserOR, 0)
}

func (s *BinOPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinOPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinOPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterBinOP(s)
	}
}

func (s *BinOPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitBinOP(s)
	}
}

func (s *BinOPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitBinOP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) BinOP() (localctx IBinOPContext) {
	localctx = NewBinOPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GoshParserRULE_binOP)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17575006176248) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnOPContext is an interface to support dynamic dispatch.
type IUnOPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUB() antlr.TerminalNode

	// IsUnOPContext differentiates from other interfaces.
	IsUnOPContext()
}

type UnOPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnOPContext() *UnOPContext {
	var p = new(UnOPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_unOP
	return p
}

func InitEmptyUnOPContext(p *UnOPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GoshParserRULE_unOP
}

func (*UnOPContext) IsUnOPContext() {}

func NewUnOPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnOPContext {
	var p = new(UnOPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GoshParserRULE_unOP

	return p
}

func (s *UnOPContext) GetParser() antlr.Parser { return s.parser }

func (s *UnOPContext) SUB() antlr.TerminalNode {
	return s.GetToken(GoshParserSUB, 0)
}

func (s *UnOPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnOPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnOPContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.EnterUnOP(s)
	}
}

func (s *UnOPContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GoshListener); ok {
		listenerT.ExitUnOP(s)
	}
}

func (s *UnOPContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GoshVisitor:
		return t.VisitUnOP(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GoshParser) UnOP() (localctx IUnOPContext) {
	localctx = NewUnOPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GoshParserRULE_unOP)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(213)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&34359753728) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *GoshParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 10:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GoshParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
