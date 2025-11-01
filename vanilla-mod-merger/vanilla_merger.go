package main

import (
	"fmt"

	"vanilla_merger/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Event struct {
	Name   string
	Fields map[string]string // or map[string]interface{} for nested blocks
}

type EventCollector struct {
	*parser.BaseParadoxListener
	Events    []Event
	currEvent *Event
	inEvent   bool
	depth     int // Track nesting
}

func NewEventCollector() *EventCollector {
	return &EventCollector{
		Events: []Event{},
	}
}

func (ec *EventCollector) EnterBlockBody(ctx *parser.BlockBodyContext) {
	ec.depth++
}

func (ec *EventCollector) ExitBlockBody(ctx *parser.BlockBodyContext) {
	ec.depth--
}

func (ec *EventCollector) EnterAssignment(ctx *parser.AssignmentContext) {
	key := ctx.Key().GetText()
	val := ctx.Value().GetText()
	// Only collect at top level and if value is a block
	if ec.depth == 0 && len(key) > 0 && len(val) > 0 && val[0] == '{' {
		ec.currEvent = &Event{Name: key, Fields: map[string]string{}}
		ec.inEvent = true
	}
}

func (ec *EventCollector) ExitAssignment(ctx *parser.AssignmentContext) {
	if ec.inEvent && ec.currEvent != nil {
		ec.Events = append(ec.Events, *ec.currEvent)
		ec.currEvent = nil
		ec.inEvent = false
	}
}

func main() {
	input, _ := antlr.NewFileStream("test-paradox.txt")
	lexer := parser.NewParadoxLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewParadoxParser(stream)
	tree := p.File()

	collector := NewEventCollector()
	antlr.ParseTreeWalkerDefault.Walk(collector, tree)

	for _, event := range collector.Events {
		fmt.Printf("Event: %s\n", event.Name)
		fmt.Printf("    Field Count: %d\n", len(event.Fields))
	}
}
