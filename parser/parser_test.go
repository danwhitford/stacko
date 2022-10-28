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
			[]stack.StackoVal{{StackoType: stack.StackoWord, Val: "+"}},
		},
		{
			[]tokeniser.Token{{TT: tokeniser.Tint, V: 1, Lexeme: "1"}},
			[]stack.StackoVal{{StackoType: stack.StackoInt, Val: 1}},
		}, {
			[]tokeniser.Token{{TT: tokeniser.Tfloat, V: 12.7, Lexeme: "12.7"}},
			[]stack.StackoVal{{StackoType: stack.StackoFloat, Val: 12.7}},
		}, {
			[]tokeniser.Token{{TT: tokeniser.Tstring, V: "spam", Lexeme: "spam"}},
			[]stack.StackoVal{{StackoType: stack.StackoString, Val: "spam"}},
		}, {
			[]tokeniser.Token{
				{TT: tokeniser.Tword, V: "+", Lexeme: "+"},
				{TT: tokeniser.Tint, V: 1, Lexeme: "1"},
				{TT: tokeniser.Tstring, V: "spam", Lexeme: "spam"}},
			[]stack.StackoVal{
				{StackoType: stack.StackoWord, Val: "+"}, 
				{StackoType: stack.StackoInt, Val: 1}, 
				{StackoType: stack.StackoString, Val: "spam"}},
		},
		{
			[]tokeniser.Token{
				{TT: tokeniser.Tword, V: "foo", Lexeme: "foo"},
				{TT: tokeniser.Tword, V: "bar", Lexeme: "bar"},
				{TT: tokeniser.Tword, V: "baz", Lexeme: "baz"},
			},
			[]stack.StackoVal{
				{StackoType: stack.StackoWord, Val: "foo"},
				{StackoType: stack.StackoWord, Val: "bar"},
				{StackoType: stack.StackoWord, Val: "baz"},
			},
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
		{
			[]tokeniser.Token{
				{TT: tokeniser.TLSqB, V: "[", Lexeme: "["},
				{TT: tokeniser.Tword, V: "foo", Lexeme: "foo"},
				{TT: tokeniser.Tword, V: "bar", Lexeme: "bar"},
				{TT: tokeniser.Tword, V: "baz", Lexeme: "baz"},
				{TT: tokeniser.TRSqB, V: "]", Lexeme: "]"},
			},
			[]stack.StackoVal{
				{StackoType: stack.StackoList, Val: []stack.StackoVal{
					{StackoType: stack.StackoWord, Val: "foo"},
					{StackoType: stack.StackoWord, Val: "bar"},
					{StackoType: stack.StackoWord, Val: "baz"},
				}},
			},
		},
	}
	testTable(t, table)
}
