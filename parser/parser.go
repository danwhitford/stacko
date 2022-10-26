package parser

import (
	"fmt"

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
		thing, err := parser.parseToken()
		if err != nil {
			return stackos, err
		}
		stackos = append(stackos, thing)
		parser.current++
	}

	return stackos, nil
}

func (parser *Parser) parseToken() (stack.StackoVal, error) {
	curr := parser.tokens[parser.current]

	switch curr.TT {
	case tokeniser.Tword:
		return stack.StackoVal{StackoType: stack.StackoWord, Val: curr.V}, nil
	case tokeniser.Tfloat:
		return stack.StackoVal{StackoType: stack.StackoFloat, Val: curr.V}, nil
	case tokeniser.Tstring:
		return stack.StackoVal{StackoType: stack.StackoString, Val: curr.V}, nil
	case tokeniser.Tint:
		return stack.StackoVal{StackoType: stack.StackoInt, Val: curr.V}, nil
	case tokeniser.TLSqB:
		return parser.readList()
	default:
		return stack.StackoVal{}, fmt.Errorf("unrecognised token type: %+v", curr)
	}
}

func (parser *Parser) readList() (stack.StackoVal, error) {
	listEls := make([]stack.StackoVal, 0)
	parser.current++ // Eat the opening '['
	for parser.current < parser.len {
		curr := parser.tokens[parser.current]
		if curr.TT == tokeniser.TRSqB {
			return stack.StackoVal{StackoType: stack.StackoList, Val: listEls}, nil
		}
		thing, err := parser.parseToken()
		if err != nil {
			return stack.StackoVal{}, err
		}
		listEls = append(listEls,thing)
		parser.current++
	}

	return stack.StackoVal{}, fmt.Errorf("unexpected end of input while parsing list")
}
