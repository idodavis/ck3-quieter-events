// Code generated from Paradox.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Paradox

import "github.com/antlr4-go/antlr/v4"

// ParadoxListener is a complete listener for a parse tree produced by ParadoxParser.
type ParadoxListener interface {
	antlr.ParseTreeListener

	// EnterFile is called when entering the file production.
	EnterFile(c *FileContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterAssignment is called when entering the assignment production.
	EnterAssignment(c *AssignmentContext)

	// EnterMacroDefinition is called when entering the macroDefinition production.
	EnterMacroDefinition(c *MacroDefinitionContext)

	// EnterMacroInsert is called when entering the macroInsert production.
	EnterMacroInsert(c *MacroInsertContext)

	// EnterVariableDefinition is called when entering the variableDefinition production.
	EnterVariableDefinition(c *VariableDefinitionContext)

	// EnterCalculation is called when entering the calculation production.
	EnterCalculation(c *CalculationContext)

	// EnterBlockBody is called when entering the blockBody production.
	EnterBlockBody(c *BlockBodyContext)

	// EnterBlockContents is called when entering the blockContents production.
	EnterBlockContents(c *BlockContentsContext)

	// EnterValueOnly is called when entering the valueOnly production.
	EnterValueOnly(c *ValueOnlyContext)

	// EnterMultiKey is called when entering the MultiKey production.
	EnterMultiKey(c *MultiKeyContext)

	// EnterColonKeyPart is called when entering the ColonKeyPart production.
	EnterColonKeyPart(c *ColonKeyPartContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterComparator is called when entering the comparator production.
	EnterComparator(c *ComparatorContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// ExitFile is called when exiting the file production.
	ExitFile(c *FileContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitAssignment is called when exiting the assignment production.
	ExitAssignment(c *AssignmentContext)

	// ExitMacroDefinition is called when exiting the macroDefinition production.
	ExitMacroDefinition(c *MacroDefinitionContext)

	// ExitMacroInsert is called when exiting the macroInsert production.
	ExitMacroInsert(c *MacroInsertContext)

	// ExitVariableDefinition is called when exiting the variableDefinition production.
	ExitVariableDefinition(c *VariableDefinitionContext)

	// ExitCalculation is called when exiting the calculation production.
	ExitCalculation(c *CalculationContext)

	// ExitBlockBody is called when exiting the blockBody production.
	ExitBlockBody(c *BlockBodyContext)

	// ExitBlockContents is called when exiting the blockContents production.
	ExitBlockContents(c *BlockContentsContext)

	// ExitValueOnly is called when exiting the valueOnly production.
	ExitValueOnly(c *ValueOnlyContext)

	// ExitMultiKey is called when exiting the MultiKey production.
	ExitMultiKey(c *MultiKeyContext)

	// ExitColonKeyPart is called when exiting the ColonKeyPart production.
	ExitColonKeyPart(c *ColonKeyPartContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitComparator is called when exiting the comparator production.
	ExitComparator(c *ComparatorContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)
}
