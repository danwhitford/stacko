package tokeniser

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//go:generate stringer -type=TokenType
type TokenType int

const (
	Tstring TokenType = iota
	Tint
	Tfloat
	Tword
	TLSqB
	TRSqB
	Tsymbol
	TLB
	TRB
	TLBrace
	TRBrace
)

type Token struct {
	TT     TokenType
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

		switch {
		case unicode.IsSpace(curr):
			t.current++
			continue
		case unicode.IsDigit(curr):
			{
				token, err := t.readNumber()
				if err != nil {
					return tokens, err
				}
				tokens = append(tokens, token)
			}
		case (curr == '"'):
			{
				token, err := t.readString()
				if err != nil {
					return tokens, err
				}
				tokens = append(tokens, token)
			}
		case (curr == '['):
			{
				token := Token{TLSqB, "[", "["}
				tokens = append(tokens, token)
			}
		case (curr == ']'):
			{
				token := Token{TRSqB, "]", "]"}
				tokens = append(tokens, token)
			}
		case (curr == '('):
			{
				token := Token{TLB, "(", "("}
				tokens = append(tokens, token)
			}
		case (curr == ')'):
			{
				token := Token{TRB, ")", ")"}
				tokens = append(tokens, token)
			}
		case (curr == '{'):
			{
				token := Token{TLBrace, "{", "{"}
				tokens = append(tokens, token)
			}
		case (curr == '}'):
			{
				token := Token{TRBrace, "}", "}"}
				tokens = append(tokens, token)
			}
		case (curr == '\''):
			token, err := t.readSymbol()
			if err != nil {
				return tokens, err
			}
			tokens = append(tokens, token)
		default:
			{
				token, err := t.readWord()
				if err != nil {
					return tokens, err
				}
				tokens = append(tokens, token)
			}
		}

		t.current++
	}

	return tokens, nil
}

func (t *Tokeniser) readSymbol() (Token, error) {
	t.current++ // Eat the opening quote
	var sb strings.Builder
	for t.current < t.len {
		curr := t.src[t.current]
		if unicode.IsSpace(curr) || curr == ']' || curr == ')' || curr == '}' {
			t.current-- // Put back so curr is processed in main loop
			return Token{Tsymbol, sb.String(), fmt.Sprintf("'%s", sb.String())}, nil
		}
		sb.WriteRune(curr)
		t.current++
	}
	return Token{Tsymbol, sb.String(), fmt.Sprintf("'%s", sb.String())}, nil
}

func (t *Tokeniser) readWord() (Token, error) {
	var sb strings.Builder
	for t.current < t.len {
		curr := t.src[t.current]
		if unicode.IsSpace(curr) || curr == ']' || curr == ')' || curr == '}' {
			t.current-- // Put back so curr is processed in main loop
			return Token{Tword, sb.String(), sb.String()}, nil
		}
		sb.WriteRune(curr)
		t.current++
	}

	return Token{Tword, sb.String(), sb.String()}, nil
}

func (t *Tokeniser) readNumber() (Token, error) {
	var sb strings.Builder
	for t.current < t.len {
		curr := t.src[t.current]
		if unicode.IsSpace(curr) {
			return stringToNumberToken(sb.String())
		} else if curr == ']' || curr == ')' {
			t.current--
			return stringToNumberToken(sb.String())
		}
		sb.WriteRune(curr)
		t.current++
	}

	return stringToNumberToken(sb.String())
}

func stringToNumberToken(s string) (Token, error) {
	v, err := strconv.Atoi(s)
	if err != nil {
		f, ferr := strconv.ParseFloat(s, 64)
		if ferr != nil {
			return Token{}, fmt.Errorf("%v / %v", err, ferr)
		}
		return Token{Tfloat, f, s}, nil
	}
	return Token{Tint, v, s}, nil
}

func (t *Tokeniser) readString() (Token, error) {
	var sb strings.Builder
	t.current++ // Eat the opening '"'
	for t.current < t.len {
		curr := t.src[t.current]
		if curr == '"' {
			return Token{Tstring, sb.String(), sb.String()}, nil
		}
		sb.WriteRune(curr)
		t.current++
	}

	return Token{}, fmt.Errorf("unexpected end of input while reading string")
}
