package tokeniser

import (
	"strconv"
	"strings"
	"unicode"
)

type tokenType int

const (
	Tstring = iota
	Tint
	Tfloat
	Tword
)

type Token struct {
	TT     tokenType
	V      interface{}
	Lexeme string
}

type Tokeniser struct {
	src     []rune
	current int
	len     int
}

func NewTokeniser(src string) Tokeniser {
	ra := []rune(src)
	return Tokeniser{ra, 0, len(ra)}
}

func (t Tokeniser) Tokenise() ([]Token, error) {
	tokens := make([]Token, 0)
	for t.current < t.len {
		curr := t.src[t.current]

		if unicode.IsDigit(curr) {
			token, err := t.readInt()
			if err != nil {
				return tokens, err
			}
			tokens = append(tokens, token)
		} else {
			token, err := t.readWord()
			if err != nil {
				return tokens, err
			}
			tokens = append(tokens, token)
		}

		t.current++
	}

	return tokens, nil
}

func (t *Tokeniser) readWord() (Token, error) {
	var sb strings.Builder
	for t.current < t.len {
		curr := t.src[t.current]
		if unicode.IsSpace(curr) {
			return Token{Tword, sb.String(), sb.String()}, nil
		}
		sb.WriteRune(curr)
		t.current++
	}

	return Token{Tword, sb.String(), sb.String()}, nil
}

func (t *Tokeniser) readInt() (Token, error) {
	var sb strings.Builder
	for t.current < t.len {
		curr := t.src[t.current]
		if unicode.IsSpace(curr) {
			return stringToIntToken(sb.String())
		}
		sb.WriteRune(curr)
		t.current++
	}

	return stringToIntToken(sb.String())
}

func stringToIntToken(s string) (Token, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		return Token{}, err
	}
	return Token{Tint, v, s}, nil
}
