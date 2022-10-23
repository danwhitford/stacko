package parser

import "github.com/danwhitford/stacko/tokeniser"

type Parser struct {
	tokens  []tokeniser.Token
	current int
	len     int
}

type stackoType int

const (
	StackoString = iota
	StackoInt
	StackoFloat
	StackoWord
)

type StackoVal struct {
	StackoType stackoType
	Val        interface{}
}

func NewParser(tokens []tokeniser.Token) Parser {
	return Parser{tokens, 0, len(tokens)}
}

func (parser Parser) Parse() ([]StackoVal, error) {
	stackos := make([]StackoVal, 0)

	for parser.current < parser.len {
		curr := parser.tokens[parser.current]
		switch curr.TT {
		case tokeniser.Tword:
			stackos = append(stackos, StackoVal{StackoWord, curr.V})
		case tokeniser.Tfloat:
			stackos = append(stackos, StackoVal{StackoFloat, curr.V})
		case tokeniser.Tstring:
			stackos = append(stackos, StackoVal{StackoString, curr.V})
		case tokeniser.Tint:
			stackos = append(stackos, StackoVal{StackoInt, curr.V})
		default:
			panic("wot u want")
		}

		parser.current++
	}

	return stackos, nil
}
