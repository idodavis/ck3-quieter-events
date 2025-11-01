// Code generated from Paradox.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Paradox

import "github.com/antlr4-go/antlr/v4"

// BaseParadoxListener is a complete listener for a parse tree produced by ParadoxParser.
type BaseParadoxListener struct{}

var _ ParadoxListener = &BaseParadoxListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseParadoxListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseParadoxListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseParadoxListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseParadoxListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFile is called when production file is entered.
func (s *BaseParadoxListener) EnterFile(ctx *FileContext) {}

// ExitFile is called when production file is exited.
func (s *BaseParadoxListener) ExitFile(ctx *FileContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseParadoxListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseParadoxListener) ExitStatement(ctx *StatementContext) {}

// EnterAssignment is called when production assignment is entered.
func (s *BaseParadoxListener) EnterAssignment(ctx *AssignmentContext) {}

// ExitAssignment is called when production assignment is exited.
func (s *BaseParadoxListener) ExitAssignment(ctx *AssignmentContext) {}

// EnterMacroDefinition is called when production macroDefinition is entered.
func (s *BaseParadoxListener) EnterMacroDefinition(ctx *MacroDefinitionContext) {}

// ExitMacroDefinition is called when production macroDefinition is exited.
func (s *BaseParadoxListener) ExitMacroDefinition(ctx *MacroDefinitionContext) {}

// EnterMacroInsert is called when production macroInsert is entered.
func (s *BaseParadoxListener) EnterMacroInsert(ctx *MacroInsertContext) {}

// ExitMacroInsert is called when production macroInsert is exited.
func (s *BaseParadoxListener) ExitMacroInsert(ctx *MacroInsertContext) {}

// EnterVariableDefinition is called when production variableDefinition is entered.
func (s *BaseParadoxListener) EnterVariableDefinition(ctx *VariableDefinitionContext) {}

// ExitVariableDefinition is called when production variableDefinition is exited.
func (s *BaseParadoxListener) ExitVariableDefinition(ctx *VariableDefinitionContext) {}

// EnterCalculation is called when production calculation is entered.
func (s *BaseParadoxListener) EnterCalculation(ctx *CalculationContext) {}

// ExitCalculation is called when production calculation is exited.
func (s *BaseParadoxListener) ExitCalculation(ctx *CalculationContext) {}

// EnterBlockBody is called when production blockBody is entered.
func (s *BaseParadoxListener) EnterBlockBody(ctx *BlockBodyContext) {}

// ExitBlockBody is called when production blockBody is exited.
func (s *BaseParadoxListener) ExitBlockBody(ctx *BlockBodyContext) {}

// EnterBlockContents is called when production blockContents is entered.
func (s *BaseParadoxListener) EnterBlockContents(ctx *BlockContentsContext) {}

// ExitBlockContents is called when production blockContents is exited.
func (s *BaseParadoxListener) ExitBlockContents(ctx *BlockContentsContext) {}

// EnterValueOnly is called when production valueOnly is entered.
func (s *BaseParadoxListener) EnterValueOnly(ctx *ValueOnlyContext) {}

// ExitValueOnly is called when production valueOnly is exited.
func (s *BaseParadoxListener) ExitValueOnly(ctx *ValueOnlyContext) {}

// EnterMultiKey is called when production MultiKey is entered.
func (s *BaseParadoxListener) EnterMultiKey(ctx *MultiKeyContext) {}

// ExitMultiKey is called when production MultiKey is exited.
func (s *BaseParadoxListener) ExitMultiKey(ctx *MultiKeyContext) {}

// EnterColonKeyPart is called when production ColonKeyPart is entered.
func (s *BaseParadoxListener) EnterColonKeyPart(ctx *ColonKeyPartContext) {}

// ExitColonKeyPart is called when production ColonKeyPart is exited.
func (s *BaseParadoxListener) ExitColonKeyPart(ctx *ColonKeyPartContext) {}

// EnterValue is called when production value is entered.
func (s *BaseParadoxListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production value is exited.
func (s *BaseParadoxListener) ExitValue(ctx *ValueContext) {}

// EnterComparator is called when production comparator is entered.
func (s *BaseParadoxListener) EnterComparator(ctx *ComparatorContext) {}

// ExitComparator is called when production comparator is exited.
func (s *BaseParadoxListener) ExitComparator(ctx *ComparatorContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseParadoxListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseParadoxListener) ExitExpr(ctx *ExprContext) {}
