package parser

import (
	"testing"

	"github.com/danwhitford/stacko/tokeniser"
)

func compare(a, b []StackoVal) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func testTable(t *testing.T, table []struct {
	in  []tokeniser.Token
	out []StackoVal
}) {
	var parser Parser
	for _, tst := range table {
		parser = NewParser(tst.in)
		vals, err := parser.Parse()
		if err != nil {
			t.Error(err)
		}
		if !compare(vals, tst.out) {
			t.Errorf("expected %+v got %+v", tst.out, vals)
		}
	}
}

func TestBasic(t *testing.T) {
	table := []struct {
		in  []tokeniser.Token
		out []StackoVal
	}{
		{
			[]tokeniser.Token{{TT: tokeniser.Tword, V: "+", Lexeme: "+"}},
			[]StackoVal{{StackoWord, "+"}},
		},
		{
			[]tokeniser.Token{{TT: tokeniser.Tint, V: 1, Lexeme: "1"}},
			[]StackoVal{{StackoInt, 1}},
		}, {
			[]tokeniser.Token{{TT: tokeniser.Tfloat, V: 12.7, Lexeme: "12.7"}},
			[]StackoVal{{StackoFloat, 12.7}},
		}, {
			[]tokeniser.Token{{TT: tokeniser.Tstring, V: "spam", Lexeme: "spam"}},
			[]StackoVal{{StackoString, "spam"}},
		}, {
			[]tokeniser.Token{
				{TT: tokeniser.Tword, V: "+", Lexeme: "+"},
				{TT: tokeniser.Tint, V: 1, Lexeme: "1"},
				{TT: tokeniser.Tstring, V: "spam", Lexeme: "spam"}},
			[]StackoVal{{StackoWord, "+"}, {StackoInt, 1}, {StackoString, "spam"}},
		},
	}
	testTable(t, table)
}
