package parser

import (
	"testing"

	"github.com/danwhitford/stacko/stack"
	"github.com/danwhitford/stacko/tokeniser"
	"github.com/google/go-cmp/cmp"
)

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
		diff := cmp.Diff(tst.out, vals)
		if diff != "" {
			t.Error(diff)
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

func TestLists(t *testing.T) {
	table := []struct {
		in  []tokeniser.Token
		out []stack.StackoVal
	}{
		{
			[]tokeniser.Token{
				{TT: tokeniser.TLSqB, V: "[", Lexeme: "["},
				{TT: tokeniser.Tstring, V: "foo bar", Lexeme: "foo bar"},
				{TT: tokeniser.TRSqB, V: "]", Lexeme: "]"},
			},
			[]stack.StackoVal{
				{StackoType: stack.StackoList, Val: []stack.StackoVal{
					{StackoType: stack.StackoString, Val: "foo bar"},
				}},
			},
		},
		{
			[]tokeniser.Token{
				{TT: tokeniser.TLSqB, V: "[", Lexeme: "["},
				{TT: tokeniser.Tstring, V: "foo", Lexeme: "foo"},
				{TT: tokeniser.TLSqB, V: "[", Lexeme: "["},
				{TT: tokeniser.Tstring, V: "bar", Lexeme: "bar"},
				{TT: tokeniser.Tstring, V: "baz", Lexeme: "baz"},
				{TT: tokeniser.TRSqB, V: "]", Lexeme: "]"},
				{TT: tokeniser.TRSqB, V: "]", Lexeme: "]"},
			},
			[]stack.StackoVal{
				{StackoType: stack.StackoList, Val: []stack.StackoVal{
					{StackoType: stack.StackoString, Val: "foo"},
					{StackoType: stack.StackoList, Val: []stack.StackoVal{
						{StackoType: stack.StackoString, Val: "bar"},
						{StackoType: stack.StackoString, Val: "baz"},
					}}}},
			},
		},
	}
	testTable(t, table)
}
