package parser

import (
	"github.com/danwhitford/stacko/stack"
	"github.com/danwhitford/stacko/tokeniser"
)

type Parser struct {
	tokens  []tokeniser.Token
	current int
	len     int
}

func NewParser(tokens []tokeniser.Token) Parser {
	return Parser{tokens, 0, len(tokens)}
}

func (parser Parser) Parse() ([]stack.StackoVal, error) {
	stackos := make([]stack.StackoVal, 0)

	for parser.current < parser.len {
		curr := parser.tokens[parser.current]
		switch curr.TT {
		case tokeniser.Tword:
			stackos = append(stackos, stack.StackoVal{stack.StackoWord, curr.V})
		case tokeniser.Tfloat:
			stackos = append(stackos, stack.StackoVal{stack.StackoFloat, curr.V})
		case tokeniser.Tstring:
			stackos = append(stackos, stack.StackoVal{stack.StackoString, curr.V})
		case tokeniser.Tint:
			stackos = append(stackos, stack.StackoVal{stack.StackoInt, curr.V})
		default:
			panic("wot u want")
		}

		parser.current++
	}

	return stackos, nil
}
