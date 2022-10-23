package parser

import (
	"testing"

	"github.com/danwhitford/stacko/tokeniser"
)

func compare(a, b []StackoVal) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
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
			[]tokeniser.Token{{tokeniser.Tword, "+", "+"}},
			[]StackoVal{{StackoWord, "+"}},
		},
		{
			[]tokeniser.Token{{tokeniser.Tint, 1, "1"}},
			[]StackoVal{{StackoInt, 1}},
		}, {
			[]tokeniser.Token{{tokeniser.Tfloat, 12.7, "12.7"}},
			[]StackoVal{{StackoFloat, 12.7}},
		}, {
			[]tokeniser.Token{{tokeniser.Tstring, "spam", "spam"}},
			[]StackoVal{{StackoString, "spam"}},
		}, {
			[]tokeniser.Token{{tokeniser.Tword, "+", "+"},{tokeniser.Tint, 1, "1"},{tokeniser.Tstring, "spam", "spam"}},
			[]StackoVal{{StackoWord, "+"},{StackoInt, 1},{StackoString, "spam"}},
		},
	}
	testTable(t, table)
}
