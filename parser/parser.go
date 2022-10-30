package parser

import (
	"fmt"

	"github.com/danwhitford/stacko/stackoval"
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

func (parser Parser) Parse() ([]stackoval.StackoVal, error) {
	stackos := make([]stackoval.StackoVal, 0)

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

func (parser *Parser) parseToken() (stackoval.StackoVal, error) {
	curr := parser.tokens[parser.current]

	switch curr.TT {
	case tokeniser.Tword:
		return stackoval.StackoVal{StackoType: stackoval.StackoWord, Val: curr.V}, nil
	case tokeniser.Tfloat:
		return stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: curr.V}, nil
	case tokeniser.Tstring:
		return stackoval.StackoVal{StackoType: stackoval.StackoString, Val: curr.V}, nil
	case tokeniser.Tint:
		return stackoval.StackoVal{StackoType: stackoval.StackoInt, Val: curr.V}, nil
	case tokeniser.TLSqB:
		return parser.readList()
	default:
		return stackoval.StackoVal{}, fmt.Errorf("unrecognised token type: %+v", curr)
	}
}

func (parser *Parser) readList() (stackoval.StackoVal, error) {
	listEls := make([]stackoval.StackoVal, 0)
	parser.current++ // Eat the opening '['
	for parser.current < parser.len {
		curr := parser.tokens[parser.current]
		if curr.TT == tokeniser.TRSqB {
			return stackoval.StackoVal{StackoType: stackoval.StackoList, Val: listEls}, nil
		}
		thing, err := parser.parseToken()
		if err != nil {
			return stackoval.StackoVal{}, err
		}
		listEls = append(listEls, thing)
		parser.current++
	}

	return stackoval.StackoVal{}, fmt.Errorf("unexpected end of input while parsing list")
}
