package parser

import (
	"testing"

	"github.com/danwhitford/stacko/stack"
	"github.com/danwhitford/stacko/tokeniser"
)

func compare(a, b []stack.StackoVal) bool {
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
	out []stack.StackoVal
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
		out []stack.StackoVal
	}{
		{
			[]tokeniser.Token{{TT: tokeniser.Tword, V: "+", Lexeme: "+"}},
			[]stack.StackoVal{{stack.StackoWord, "+"}},
		},
		{
			[]tokeniser.Token{{TT: tokeniser.Tint, V: 1, Lexeme: "1"}},
			[]stack.StackoVal{{stack.StackoInt, 1}},
		}, {
			[]tokeniser.Token{{TT: tokeniser.Tfloat, V: 12.7, Lexeme: "12.7"}},
			[]stack.StackoVal{{stack.StackoFloat, 12.7}},
		}, {
			[]tokeniser.Token{{TT: tokeniser.Tstring, V: "spam", Lexeme: "spam"}},
			[]stack.StackoVal{{stack.StackoString, "spam"}},
		}, {
			[]tokeniser.Token{
				{TT: tokeniser.Tword, V: "+", Lexeme: "+"},
				{TT: tokeniser.Tint, V: 1, Lexeme: "1"},
				{TT: tokeniser.Tstring, V: "spam", Lexeme: "spam"}},
			[]stack.StackoVal{{stack.StackoWord, "+"}, {stack.StackoInt, 1}, {stack.StackoString, "spam"}},
		},
	}
	testTable(t, table)
}
