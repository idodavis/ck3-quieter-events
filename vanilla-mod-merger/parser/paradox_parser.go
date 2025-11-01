// Code generated from Paradox.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Paradox

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

type ParadoxParser struct {
	*antlr.BaseParser
}

var ParadoxParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func paradoxParserInit() {
	staticData := &ParadoxParserStaticData
	staticData.LiteralNames = []string{
		"", "", "", "", "'@:define'", "'@:insert'", "'@:register_variable'",
		"", "'@['", "']'", "'='", "'!='", "'>='", "'<='", "'>'", "'<'", "'{'",
		"'}'", "'('", "')'", "'+'", "'-'", "'*'", "'/'", "':'",
	}
	staticData.SymbolicNames = []string{
		"", "ID", "NUMBER", "STRING", "DEFINE", "INSERT", "REGISTER_VARIABLE",
		"VAR", "CALC_START", "CALC_END", "EQUALS", "NOT_EQUALS", "GTE", "LTE",
		"GT", "LT", "LBRACE", "RBRACE", "LPAREN", "RPAREN", "ADD", "SUB", "MUL",
		"DIV", "COLON", "COMMENT", "WS",
	}
	staticData.RuleNames = []string{
		"file", "statement", "assignment", "macroDefinition", "macroInsert",
		"variableDefinition", "calculation", "blockBody", "blockContents", "valueOnly",
		"key", "keyPart", "value", "comparator", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 26, 146, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 1, 0,
		1, 0, 5, 0, 34, 8, 0, 10, 0, 12, 0, 37, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 47, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 72, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 5, 8, 85, 8, 8, 10, 8, 12, 8, 88,
		9, 8, 1, 9, 1, 9, 1, 10, 1, 10, 4, 10, 94, 8, 10, 11, 10, 12, 10, 95, 1,
		10, 5, 10, 99, 8, 10, 10, 10, 12, 10, 102, 9, 10, 1, 11, 1, 11, 1, 11,
		3, 11, 107, 8, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 114, 8, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 3, 14, 127, 8, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 5, 14, 141, 8, 14, 10, 14, 12, 14, 144,
		9, 14, 1, 14, 0, 1, 28, 15, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
		24, 26, 28, 0, 1, 1, 0, 10, 15, 156, 0, 35, 1, 0, 0, 0, 2, 46, 1, 0, 0,
		0, 4, 48, 1, 0, 0, 0, 6, 52, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0, 10, 71, 1,
		0, 0, 0, 12, 73, 1, 0, 0, 0, 14, 77, 1, 0, 0, 0, 16, 86, 1, 0, 0, 0, 18,
		89, 1, 0, 0, 0, 20, 91, 1, 0, 0, 0, 22, 103, 1, 0, 0, 0, 24, 113, 1, 0,
		0, 0, 26, 115, 1, 0, 0, 0, 28, 126, 1, 0, 0, 0, 30, 34, 3, 2, 1, 0, 31,
		34, 5, 25, 0, 0, 32, 34, 5, 26, 0, 0, 33, 30, 1, 0, 0, 0, 33, 31, 1, 0,
		0, 0, 33, 32, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36,
		1, 0, 0, 0, 36, 38, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 39, 5, 0, 0, 1,
		39, 1, 1, 0, 0, 0, 40, 47, 3, 4, 2, 0, 41, 47, 3, 6, 3, 0, 42, 47, 3, 8,
		4, 0, 43, 47, 3, 10, 5, 0, 44, 47, 3, 12, 6, 0, 45, 47, 3, 18, 9, 0, 46,
		40, 1, 0, 0, 0, 46, 41, 1, 0, 0, 0, 46, 42, 1, 0, 0, 0, 46, 43, 1, 0, 0,
		0, 46, 44, 1, 0, 0, 0, 46, 45, 1, 0, 0, 0, 47, 3, 1, 0, 0, 0, 48, 49, 3,
		20, 10, 0, 49, 50, 3, 26, 13, 0, 50, 51, 3, 24, 12, 0, 51, 5, 1, 0, 0,
		0, 52, 53, 5, 4, 0, 0, 53, 54, 3, 20, 10, 0, 54, 55, 3, 26, 13, 0, 55,
		56, 3, 14, 7, 0, 56, 7, 1, 0, 0, 0, 57, 58, 5, 5, 0, 0, 58, 59, 3, 20,
		10, 0, 59, 60, 3, 26, 13, 0, 60, 61, 3, 14, 7, 0, 61, 9, 1, 0, 0, 0, 62,
		63, 5, 7, 0, 0, 63, 64, 3, 26, 13, 0, 64, 65, 3, 24, 12, 0, 65, 72, 1,
		0, 0, 0, 66, 67, 5, 6, 0, 0, 67, 68, 3, 20, 10, 0, 68, 69, 3, 26, 13, 0,
		69, 70, 3, 24, 12, 0, 70, 72, 1, 0, 0, 0, 71, 62, 1, 0, 0, 0, 71, 66, 1,
		0, 0, 0, 72, 11, 1, 0, 0, 0, 73, 74, 5, 8, 0, 0, 74, 75, 3, 28, 14, 0,
		75, 76, 5, 9, 0, 0, 76, 13, 1, 0, 0, 0, 77, 78, 5, 16, 0, 0, 78, 79, 3,
		16, 8, 0, 79, 80, 5, 17, 0, 0, 80, 15, 1, 0, 0, 0, 81, 85, 3, 2, 1, 0,
		82, 85, 5, 25, 0, 0, 83, 85, 5, 26, 0, 0, 84, 81, 1, 0, 0, 0, 84, 82, 1,
		0, 0, 0, 84, 83, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 86,
		87, 1, 0, 0, 0, 87, 17, 1, 0, 0, 0, 88, 86, 1, 0, 0, 0, 89, 90, 3, 24,
		12, 0, 90, 19, 1, 0, 0, 0, 91, 100, 3, 22, 11, 0, 92, 94, 5, 26, 0, 0,
		93, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 95, 96, 1,
		0, 0, 0, 96, 97, 1, 0, 0, 0, 97, 99, 3, 22, 11, 0, 98, 93, 1, 0, 0, 0,
		99, 102, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 21,
		1, 0, 0, 0, 102, 100, 1, 0, 0, 0, 103, 106, 5, 1, 0, 0, 104, 105, 5, 24,
		0, 0, 105, 107, 5, 1, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0,
		107, 23, 1, 0, 0, 0, 108, 114, 5, 2, 0, 0, 109, 114, 5, 3, 0, 0, 110, 114,
		5, 1, 0, 0, 111, 114, 3, 14, 7, 0, 112, 114, 3, 12, 6, 0, 113, 108, 1,
		0, 0, 0, 113, 109, 1, 0, 0, 0, 113, 110, 1, 0, 0, 0, 113, 111, 1, 0, 0,
		0, 113, 112, 1, 0, 0, 0, 114, 25, 1, 0, 0, 0, 115, 116, 7, 0, 0, 0, 116,
		27, 1, 0, 0, 0, 117, 118, 6, 14, -1, 0, 118, 119, 5, 21, 0, 0, 119, 127,
		3, 28, 14, 4, 120, 127, 5, 2, 0, 0, 121, 122, 5, 18, 0, 0, 122, 123, 3,
		28, 14, 0, 123, 124, 5, 19, 0, 0, 124, 127, 1, 0, 0, 0, 125, 127, 5, 1,
		0, 0, 126, 117, 1, 0, 0, 0, 126, 120, 1, 0, 0, 0, 126, 121, 1, 0, 0, 0,
		126, 125, 1, 0, 0, 0, 127, 142, 1, 0, 0, 0, 128, 129, 10, 8, 0, 0, 129,
		130, 5, 20, 0, 0, 130, 141, 3, 28, 14, 9, 131, 132, 10, 7, 0, 0, 132, 133,
		5, 21, 0, 0, 133, 141, 3, 28, 14, 8, 134, 135, 10, 6, 0, 0, 135, 136, 5,
		22, 0, 0, 136, 141, 3, 28, 14, 7, 137, 138, 10, 5, 0, 0, 138, 139, 5, 23,
		0, 0, 139, 141, 3, 28, 14, 6, 140, 128, 1, 0, 0, 0, 140, 131, 1, 0, 0,
		0, 140, 134, 1, 0, 0, 0, 140, 137, 1, 0, 0, 0, 141, 144, 1, 0, 0, 0, 142,
		140, 1, 0, 0, 0, 142, 143, 1, 0, 0, 0, 143, 29, 1, 0, 0, 0, 144, 142, 1,
		0, 0, 0, 13, 33, 35, 46, 71, 84, 86, 95, 100, 106, 113, 126, 140, 142,
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

// ParadoxParserInit initializes any static state used to implement ParadoxParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewParadoxParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ParadoxParserInit() {
	staticData := &ParadoxParserStaticData
	staticData.once.Do(paradoxParserInit)
}

// NewParadoxParser produces a new parser instance for the optional input antlr.TokenStream.
func NewParadoxParser(input antlr.TokenStream) *ParadoxParser {
	ParadoxParserInit()
	this := new(ParadoxParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ParadoxParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Paradox.g4"

	return this
}

// ParadoxParser tokens.
const (
	ParadoxParserEOF               = antlr.TokenEOF
	ParadoxParserID                = 1
	ParadoxParserNUMBER            = 2
	ParadoxParserSTRING            = 3
	ParadoxParserDEFINE            = 4
	ParadoxParserINSERT            = 5
	ParadoxParserREGISTER_VARIABLE = 6
	ParadoxParserVAR               = 7
	ParadoxParserCALC_START        = 8
	ParadoxParserCALC_END          = 9
	ParadoxParserEQUALS            = 10
	ParadoxParserNOT_EQUALS        = 11
	ParadoxParserGTE               = 12
	ParadoxParserLTE               = 13
	ParadoxParserGT                = 14
	ParadoxParserLT                = 15
	ParadoxParserLBRACE            = 16
	ParadoxParserRBRACE            = 17
	ParadoxParserLPAREN            = 18
	ParadoxParserRPAREN            = 19
	ParadoxParserADD               = 20
	ParadoxParserSUB               = 21
	ParadoxParserMUL               = 22
	ParadoxParserDIV               = 23
	ParadoxParserCOLON             = 24
	ParadoxParserCOMMENT           = 25
	ParadoxParserWS                = 26
)

// ParadoxParser rules.
const (
	ParadoxParserRULE_file               = 0
	ParadoxParserRULE_statement          = 1
	ParadoxParserRULE_assignment         = 2
	ParadoxParserRULE_macroDefinition    = 3
	ParadoxParserRULE_macroInsert        = 4
	ParadoxParserRULE_variableDefinition = 5
	ParadoxParserRULE_calculation        = 6
	ParadoxParserRULE_blockBody          = 7
	ParadoxParserRULE_blockContents      = 8
	ParadoxParserRULE_valueOnly          = 9
	ParadoxParserRULE_key                = 10
	ParadoxParserRULE_keyPart            = 11
	ParadoxParserRULE_value              = 12
	ParadoxParserRULE_comparator         = 13
	ParadoxParserRULE_expr               = 14
)

// IFileContext is an interface to support dynamic dispatch.
type IFileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllCOMMENT() []antlr.TerminalNode
	COMMENT(i int) antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsFileContext differentiates from other interfaces.
	IsFileContext()
}

type FileContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileContext() *FileContext {
	var p = new(FileContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_file
	return p
}

func InitEmptyFileContext(p *FileContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_file
}

func (*FileContext) IsFileContext() {}

func NewFileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileContext {
	var p = new(FileContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_file

	return p
}

func (s *FileContext) GetParser() antlr.Parser { return s.parser }

func (s *FileContext) EOF() antlr.TerminalNode {
	return s.GetToken(ParadoxParserEOF, 0)
}

func (s *FileContext) AllStatement() []IStatementContext {
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

func (s *FileContext) Statement(i int) IStatementContext {
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

func (s *FileContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(ParadoxParserCOMMENT)
}

func (s *FileContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(ParadoxParserCOMMENT, i)
}

func (s *FileContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ParadoxParserWS)
}

func (s *FileContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ParadoxParserWS, i)
}

func (s *FileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterFile(s)
	}
}

func (s *FileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitFile(s)
	}
}

func (p *ParadoxParser) File() (localctx IFileContext) {
	localctx = NewFileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ParadoxParserRULE_file)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(35)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&100729342) != 0 {
		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ParadoxParserID, ParadoxParserNUMBER, ParadoxParserSTRING, ParadoxParserDEFINE, ParadoxParserINSERT, ParadoxParserREGISTER_VARIABLE, ParadoxParserVAR, ParadoxParserCALC_START, ParadoxParserLBRACE:
			{
				p.SetState(30)
				p.Statement()
			}

		case ParadoxParserCOMMENT:
			{
				p.SetState(31)
				p.Match(ParadoxParserCOMMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case ParadoxParserWS:
			{
				p.SetState(32)
				p.Match(ParadoxParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(38)
		p.Match(ParadoxParserEOF)
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

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Assignment() IAssignmentContext
	MacroDefinition() IMacroDefinitionContext
	MacroInsert() IMacroInsertContext
	VariableDefinition() IVariableDefinitionContext
	Calculation() ICalculationContext
	ValueOnly() IValueOnlyContext

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
	p.RuleIndex = ParadoxParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Assignment() IAssignmentContext {
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

func (s *StatementContext) MacroDefinition() IMacroDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroDefinitionContext)
}

func (s *StatementContext) MacroInsert() IMacroInsertContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMacroInsertContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMacroInsertContext)
}

func (s *StatementContext) VariableDefinition() IVariableDefinitionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDefinitionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionContext)
}

func (s *StatementContext) Calculation() ICalculationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculationContext)
}

func (s *StatementContext) ValueOnly() IValueOnlyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueOnlyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueOnlyContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *ParadoxParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ParadoxParserRULE_statement)
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Assignment()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.MacroDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.MacroInsert()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(43)
			p.VariableDefinition()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(44)
			p.Calculation()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(45)
			p.ValueOnly()
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
	Key() IKeyContext
	Comparator() IComparatorContext
	Value() IValueContext

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
	p.RuleIndex = ParadoxParserRULE_assignment
	return p
}

func InitEmptyAssignmentContext(p *AssignmentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_assignment
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *AssignmentContext) Comparator() IComparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *AssignmentContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *ParadoxParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ParadoxParserRULE_assignment)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Key()
	}
	{
		p.SetState(49)
		p.Comparator()
	}
	{
		p.SetState(50)
		p.Value()
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

// IMacroDefinitionContext is an interface to support dynamic dispatch.
type IMacroDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DEFINE() antlr.TerminalNode
	Key() IKeyContext
	Comparator() IComparatorContext
	BlockBody() IBlockBodyContext

	// IsMacroDefinitionContext differentiates from other interfaces.
	IsMacroDefinitionContext()
}

type MacroDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacroDefinitionContext() *MacroDefinitionContext {
	var p = new(MacroDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_macroDefinition
	return p
}

func InitEmptyMacroDefinitionContext(p *MacroDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_macroDefinition
}

func (*MacroDefinitionContext) IsMacroDefinitionContext() {}

func NewMacroDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MacroDefinitionContext {
	var p = new(MacroDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_macroDefinition

	return p
}

func (s *MacroDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *MacroDefinitionContext) DEFINE() antlr.TerminalNode {
	return s.GetToken(ParadoxParserDEFINE, 0)
}

func (s *MacroDefinitionContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *MacroDefinitionContext) Comparator() IComparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *MacroDefinitionContext) BlockBody() IBlockBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *MacroDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MacroDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterMacroDefinition(s)
	}
}

func (s *MacroDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitMacroDefinition(s)
	}
}

func (p *ParadoxParser) MacroDefinition() (localctx IMacroDefinitionContext) {
	localctx = NewMacroDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ParadoxParserRULE_macroDefinition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(52)
		p.Match(ParadoxParserDEFINE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Key()
	}
	{
		p.SetState(54)
		p.Comparator()
	}
	{
		p.SetState(55)
		p.BlockBody()
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

// IMacroInsertContext is an interface to support dynamic dispatch.
type IMacroInsertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INSERT() antlr.TerminalNode
	Key() IKeyContext
	Comparator() IComparatorContext
	BlockBody() IBlockBodyContext

	// IsMacroInsertContext differentiates from other interfaces.
	IsMacroInsertContext()
}

type MacroInsertContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMacroInsertContext() *MacroInsertContext {
	var p = new(MacroInsertContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_macroInsert
	return p
}

func InitEmptyMacroInsertContext(p *MacroInsertContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_macroInsert
}

func (*MacroInsertContext) IsMacroInsertContext() {}

func NewMacroInsertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MacroInsertContext {
	var p = new(MacroInsertContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_macroInsert

	return p
}

func (s *MacroInsertContext) GetParser() antlr.Parser { return s.parser }

func (s *MacroInsertContext) INSERT() antlr.TerminalNode {
	return s.GetToken(ParadoxParserINSERT, 0)
}

func (s *MacroInsertContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *MacroInsertContext) Comparator() IComparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *MacroInsertContext) BlockBody() IBlockBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *MacroInsertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MacroInsertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MacroInsertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterMacroInsert(s)
	}
}

func (s *MacroInsertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitMacroInsert(s)
	}
}

func (p *ParadoxParser) MacroInsert() (localctx IMacroInsertContext) {
	localctx = NewMacroInsertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ParadoxParserRULE_macroInsert)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(ParadoxParserINSERT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Key()
	}
	{
		p.SetState(59)
		p.Comparator()
	}
	{
		p.SetState(60)
		p.BlockBody()
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

// IVariableDefinitionContext is an interface to support dynamic dispatch.
type IVariableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR() antlr.TerminalNode
	Comparator() IComparatorContext
	Value() IValueContext
	REGISTER_VARIABLE() antlr.TerminalNode
	Key() IKeyContext

	// IsVariableDefinitionContext differentiates from other interfaces.
	IsVariableDefinitionContext()
}

type VariableDefinitionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionContext() *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_variableDefinition
	return p
}

func InitEmptyVariableDefinitionContext(p *VariableDefinitionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_variableDefinition
}

func (*VariableDefinitionContext) IsVariableDefinitionContext() {}

func NewVariableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_variableDefinition

	return p
}

func (s *VariableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionContext) VAR() antlr.TerminalNode {
	return s.GetToken(ParadoxParserVAR, 0)
}

func (s *VariableDefinitionContext) Comparator() IComparatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IComparatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IComparatorContext)
}

func (s *VariableDefinitionContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *VariableDefinitionContext) REGISTER_VARIABLE() antlr.TerminalNode {
	return s.GetToken(ParadoxParserREGISTER_VARIABLE, 0)
}

func (s *VariableDefinitionContext) Key() IKeyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyContext)
}

func (s *VariableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterVariableDefinition(s)
	}
}

func (s *VariableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitVariableDefinition(s)
	}
}

func (p *ParadoxParser) VariableDefinition() (localctx IVariableDefinitionContext) {
	localctx = NewVariableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ParadoxParserRULE_variableDefinition)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParadoxParserVAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(ParadoxParserVAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(63)
			p.Comparator()
		}
		{
			p.SetState(64)
			p.Value()
		}

	case ParadoxParserREGISTER_VARIABLE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(66)
			p.Match(ParadoxParserREGISTER_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Key()
		}
		{
			p.SetState(68)
			p.Comparator()
		}
		{
			p.SetState(69)
			p.Value()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// ICalculationContext is an interface to support dynamic dispatch.
type ICalculationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CALC_START() antlr.TerminalNode
	Expr() IExprContext
	CALC_END() antlr.TerminalNode

	// IsCalculationContext differentiates from other interfaces.
	IsCalculationContext()
}

type CalculationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCalculationContext() *CalculationContext {
	var p = new(CalculationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_calculation
	return p
}

func InitEmptyCalculationContext(p *CalculationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_calculation
}

func (*CalculationContext) IsCalculationContext() {}

func NewCalculationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CalculationContext {
	var p = new(CalculationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_calculation

	return p
}

func (s *CalculationContext) GetParser() antlr.Parser { return s.parser }

func (s *CalculationContext) CALC_START() antlr.TerminalNode {
	return s.GetToken(ParadoxParserCALC_START, 0)
}

func (s *CalculationContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CalculationContext) CALC_END() antlr.TerminalNode {
	return s.GetToken(ParadoxParserCALC_END, 0)
}

func (s *CalculationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CalculationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CalculationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterCalculation(s)
	}
}

func (s *CalculationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitCalculation(s)
	}
}

func (p *ParadoxParser) Calculation() (localctx ICalculationContext) {
	localctx = NewCalculationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ParadoxParserRULE_calculation)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(ParadoxParserCALC_START)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(74)
		p.expr(0)
	}
	{
		p.SetState(75)
		p.Match(ParadoxParserCALC_END)
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

// IBlockBodyContext is an interface to support dynamic dispatch.
type IBlockBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LBRACE() antlr.TerminalNode
	BlockContents() IBlockContentsContext
	RBRACE() antlr.TerminalNode

	// IsBlockBodyContext differentiates from other interfaces.
	IsBlockBodyContext()
}

type BlockBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockBodyContext() *BlockBodyContext {
	var p = new(BlockBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_blockBody
	return p
}

func InitEmptyBlockBodyContext(p *BlockBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_blockBody
}

func (*BlockBodyContext) IsBlockBodyContext() {}

func NewBlockBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockBodyContext {
	var p = new(BlockBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_blockBody

	return p
}

func (s *BlockBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockBodyContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ParadoxParserLBRACE, 0)
}

func (s *BlockBodyContext) BlockContents() IBlockContentsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContentsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContentsContext)
}

func (s *BlockBodyContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ParadoxParserRBRACE, 0)
}

func (s *BlockBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterBlockBody(s)
	}
}

func (s *BlockBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitBlockBody(s)
	}
}

func (p *ParadoxParser) BlockBody() (localctx IBlockBodyContext) {
	localctx = NewBlockBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ParadoxParserRULE_blockBody)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)
		p.Match(ParadoxParserLBRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(78)
		p.BlockContents()
	}
	{
		p.SetState(79)
		p.Match(ParadoxParserRBRACE)
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

// IBlockContentsContext is an interface to support dynamic dispatch.
type IBlockContentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	AllCOMMENT() []antlr.TerminalNode
	COMMENT(i int) antlr.TerminalNode
	AllWS() []antlr.TerminalNode
	WS(i int) antlr.TerminalNode

	// IsBlockContentsContext differentiates from other interfaces.
	IsBlockContentsContext()
}

type BlockContentsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContentsContext() *BlockContentsContext {
	var p = new(BlockContentsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_blockContents
	return p
}

func InitEmptyBlockContentsContext(p *BlockContentsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_blockContents
}

func (*BlockContentsContext) IsBlockContentsContext() {}

func NewBlockContentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContentsContext {
	var p = new(BlockContentsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_blockContents

	return p
}

func (s *BlockContentsContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContentsContext) AllStatement() []IStatementContext {
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

func (s *BlockContentsContext) Statement(i int) IStatementContext {
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

func (s *BlockContentsContext) AllCOMMENT() []antlr.TerminalNode {
	return s.GetTokens(ParadoxParserCOMMENT)
}

func (s *BlockContentsContext) COMMENT(i int) antlr.TerminalNode {
	return s.GetToken(ParadoxParserCOMMENT, i)
}

func (s *BlockContentsContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ParadoxParserWS)
}

func (s *BlockContentsContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ParadoxParserWS, i)
}

func (s *BlockContentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterBlockContents(s)
	}
}

func (s *BlockContentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitBlockContents(s)
	}
}

func (p *ParadoxParser) BlockContents() (localctx IBlockContentsContext) {
	localctx = NewBlockContentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ParadoxParserRULE_blockContents)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&100729342) != 0 {
		p.SetState(84)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ParadoxParserID, ParadoxParserNUMBER, ParadoxParserSTRING, ParadoxParserDEFINE, ParadoxParserINSERT, ParadoxParserREGISTER_VARIABLE, ParadoxParserVAR, ParadoxParserCALC_START, ParadoxParserLBRACE:
			{
				p.SetState(81)
				p.Statement()
			}

		case ParadoxParserCOMMENT:
			{
				p.SetState(82)
				p.Match(ParadoxParserCOMMENT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		case ParadoxParserWS:
			{
				p.SetState(83)
				p.Match(ParadoxParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(88)
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

// IValueOnlyContext is an interface to support dynamic dispatch.
type IValueOnlyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Value() IValueContext

	// IsValueOnlyContext differentiates from other interfaces.
	IsValueOnlyContext()
}

type ValueOnlyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueOnlyContext() *ValueOnlyContext {
	var p = new(ValueOnlyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_valueOnly
	return p
}

func InitEmptyValueOnlyContext(p *ValueOnlyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_valueOnly
}

func (*ValueOnlyContext) IsValueOnlyContext() {}

func NewValueOnlyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueOnlyContext {
	var p = new(ValueOnlyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_valueOnly

	return p
}

func (s *ValueOnlyContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueOnlyContext) Value() IValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ValueOnlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueOnlyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueOnlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterValueOnly(s)
	}
}

func (s *ValueOnlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitValueOnly(s)
	}
}

func (p *ParadoxParser) ValueOnly() (localctx IValueOnlyContext) {
	localctx = NewValueOnlyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ParadoxParserRULE_valueOnly)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(89)
		p.Value()
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

// IKeyContext is an interface to support dynamic dispatch.
type IKeyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsKeyContext differentiates from other interfaces.
	IsKeyContext()
}

type KeyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyContext() *KeyContext {
	var p = new(KeyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_key
	return p
}

func InitEmptyKeyContext(p *KeyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_key
}

func (*KeyContext) IsKeyContext() {}

func NewKeyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyContext {
	var p = new(KeyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_key

	return p
}

func (s *KeyContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyContext) CopyAll(ctx *KeyContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *KeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultiKeyContext struct {
	KeyContext
}

func NewMultiKeyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiKeyContext {
	var p = new(MultiKeyContext)

	InitEmptyKeyContext(&p.KeyContext)
	p.parser = parser
	p.CopyAll(ctx.(*KeyContext))

	return p
}

func (s *MultiKeyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiKeyContext) AllKeyPart() []IKeyPartContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IKeyPartContext); ok {
			len++
		}
	}

	tst := make([]IKeyPartContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IKeyPartContext); ok {
			tst[i] = t.(IKeyPartContext)
			i++
		}
	}

	return tst
}

func (s *MultiKeyContext) KeyPart(i int) IKeyPartContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyPartContext); ok {
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

	return t.(IKeyPartContext)
}

func (s *MultiKeyContext) AllWS() []antlr.TerminalNode {
	return s.GetTokens(ParadoxParserWS)
}

func (s *MultiKeyContext) WS(i int) antlr.TerminalNode {
	return s.GetToken(ParadoxParserWS, i)
}

func (s *MultiKeyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterMultiKey(s)
	}
}

func (s *MultiKeyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitMultiKey(s)
	}
}

func (p *ParadoxParser) Key() (localctx IKeyContext) {
	localctx = NewKeyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ParadoxParserRULE_key)
	var _la int

	localctx = NewMultiKeyContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.KeyPart()
	}
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ParadoxParserWS {
		p.SetState(93)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == ParadoxParserWS {
			{
				p.SetState(92)
				p.Match(ParadoxParserWS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(95)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(97)
			p.KeyPart()
		}

		p.SetState(102)
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

// IKeyPartContext is an interface to support dynamic dispatch.
type IKeyPartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsKeyPartContext differentiates from other interfaces.
	IsKeyPartContext()
}

type KeyPartContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyPartContext() *KeyPartContext {
	var p = new(KeyPartContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_keyPart
	return p
}

func InitEmptyKeyPartContext(p *KeyPartContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_keyPart
}

func (*KeyPartContext) IsKeyPartContext() {}

func NewKeyPartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyPartContext {
	var p = new(KeyPartContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_keyPart

	return p
}

func (s *KeyPartContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyPartContext) CopyAll(ctx *KeyPartContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *KeyPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyPartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ColonKeyPartContext struct {
	KeyPartContext
}

func NewColonKeyPartContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColonKeyPartContext {
	var p = new(ColonKeyPartContext)

	InitEmptyKeyPartContext(&p.KeyPartContext)
	p.parser = parser
	p.CopyAll(ctx.(*KeyPartContext))

	return p
}

func (s *ColonKeyPartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColonKeyPartContext) AllID() []antlr.TerminalNode {
	return s.GetTokens(ParadoxParserID)
}

func (s *ColonKeyPartContext) ID(i int) antlr.TerminalNode {
	return s.GetToken(ParadoxParserID, i)
}

func (s *ColonKeyPartContext) COLON() antlr.TerminalNode {
	return s.GetToken(ParadoxParserCOLON, 0)
}

func (s *ColonKeyPartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterColonKeyPart(s)
	}
}

func (s *ColonKeyPartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitColonKeyPart(s)
	}
}

func (p *ParadoxParser) KeyPart() (localctx IKeyPartContext) {
	localctx = NewKeyPartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ParadoxParserRULE_keyPart)
	var _la int

	localctx = NewColonKeyPartContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(ParadoxParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ParadoxParserCOLON {
		{
			p.SetState(104)
			p.Match(ParadoxParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Match(ParadoxParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NUMBER() antlr.TerminalNode
	STRING() antlr.TerminalNode
	ID() antlr.TerminalNode
	BlockBody() IBlockBodyContext
	Calculation() ICalculationContext

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_value
	return p
}

func InitEmptyValueContext(p *ValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_value
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ParadoxParserNUMBER, 0)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(ParadoxParserSTRING, 0)
}

func (s *ValueContext) ID() antlr.TerminalNode {
	return s.GetToken(ParadoxParserID, 0)
}

func (s *ValueContext) BlockBody() IBlockBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockBodyContext)
}

func (s *ValueContext) Calculation() ICalculationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICalculationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICalculationContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *ParadoxParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ParadoxParserRULE_value)
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParadoxParserNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(ParadoxParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParadoxParserSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.Match(ParadoxParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParadoxParserID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(110)
			p.Match(ParadoxParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParadoxParserLBRACE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(111)
			p.BlockBody()
		}

	case ParadoxParserCALC_START:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(112)
			p.Calculation()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IComparatorContext is an interface to support dynamic dispatch.
type IComparatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQUALS() antlr.TerminalNode
	NOT_EQUALS() antlr.TerminalNode
	GTE() antlr.TerminalNode
	LTE() antlr.TerminalNode
	GT() antlr.TerminalNode
	LT() antlr.TerminalNode

	// IsComparatorContext differentiates from other interfaces.
	IsComparatorContext()
}

type ComparatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparatorContext() *ComparatorContext {
	var p = new(ComparatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_comparator
	return p
}

func InitEmptyComparatorContext(p *ComparatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_comparator
}

func (*ComparatorContext) IsComparatorContext() {}

func NewComparatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparatorContext {
	var p = new(ComparatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_comparator

	return p
}

func (s *ComparatorContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparatorContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(ParadoxParserEQUALS, 0)
}

func (s *ComparatorContext) NOT_EQUALS() antlr.TerminalNode {
	return s.GetToken(ParadoxParserNOT_EQUALS, 0)
}

func (s *ComparatorContext) GTE() antlr.TerminalNode {
	return s.GetToken(ParadoxParserGTE, 0)
}

func (s *ComparatorContext) LTE() antlr.TerminalNode {
	return s.GetToken(ParadoxParserLTE, 0)
}

func (s *ComparatorContext) GT() antlr.TerminalNode {
	return s.GetToken(ParadoxParserGT, 0)
}

func (s *ComparatorContext) LT() antlr.TerminalNode {
	return s.GetToken(ParadoxParserLT, 0)
}

func (s *ComparatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterComparator(s)
	}
}

func (s *ComparatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitComparator(s)
	}
}

func (p *ParadoxParser) Comparator() (localctx IComparatorContext) {
	localctx = NewComparatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ParadoxParserRULE_comparator)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&64512) != 0) {
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SUB() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext
	NUMBER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ID() antlr.TerminalNode
	ADD() antlr.TerminalNode
	MUL() antlr.TerminalNode
	DIV() antlr.TerminalNode

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ParadoxParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ParadoxParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) SUB() antlr.TerminalNode {
	return s.GetToken(ParadoxParserSUB, 0)
}

func (s *ExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ParadoxParserNUMBER, 0)
}

func (s *ExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ParadoxParserLPAREN, 0)
}

func (s *ExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ParadoxParserRPAREN, 0)
}

func (s *ExprContext) ID() antlr.TerminalNode {
	return s.GetToken(ParadoxParserID, 0)
}

func (s *ExprContext) ADD() antlr.TerminalNode {
	return s.GetToken(ParadoxParserADD, 0)
}

func (s *ExprContext) MUL() antlr.TerminalNode {
	return s.GetToken(ParadoxParserMUL, 0)
}

func (s *ExprContext) DIV() antlr.TerminalNode {
	return s.GetToken(ParadoxParserDIV, 0)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ParadoxListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *ParadoxParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ParadoxParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, ParadoxParserRULE_expr, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ParadoxParserSUB:
		{
			p.SetState(118)
			p.Match(ParadoxParserSUB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.expr(4)
		}

	case ParadoxParserNUMBER:
		{
			p.SetState(120)
			p.Match(ParadoxParserNUMBER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParadoxParserLPAREN:
		{
			p.SetState(121)
			p.Match(ParadoxParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(122)
			p.expr(0)
		}
		{
			p.SetState(123)
			p.Match(ParadoxParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ParadoxParserID:
		{
			p.SetState(125)
			p.Match(ParadoxParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(140)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ParadoxParserRULE_expr)
				p.SetState(128)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(129)
					p.Match(ParadoxParserADD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(130)
					p.expr(9)
				}

			case 2:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ParadoxParserRULE_expr)
				p.SetState(131)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(132)
					p.Match(ParadoxParserSUB)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(133)
					p.expr(8)
				}

			case 3:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ParadoxParserRULE_expr)
				p.SetState(134)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(135)
					p.Match(ParadoxParserMUL)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(136)
					p.expr(7)
				}

			case 4:
				localctx = NewExprContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, ParadoxParserRULE_expr)
				p.SetState(137)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(138)
					p.Match(ParadoxParserDIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(139)
					p.expr(6)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
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

func (p *ParadoxParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 14:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ParadoxParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 5)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
