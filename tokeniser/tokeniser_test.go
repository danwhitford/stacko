package tokeniser

import "testing"

func compare(a, b []Token) bool {
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

func testTable(t *testing.T, table []struct{in string; out []Token }) {
	var tokeniser Tokeniser
	for _, tst := range table {
		tokeniser = NewTokeniser(tst.in)
		tokens, err := tokeniser.Tokenise()
		if err != nil {
			t.Error(err)
		}
		if !compare(tokens, tst.out) {
			t.Errorf("expected %+v got %+v", tst.out, tokens)
		}
	}
}

func TestWords(t *testing.T) {
	table := []struct {
		in  string
		out []Token
	}{
		{
			"+",
			[]Token{{Tword, "+", "+"}},
		},
		{
			"spoon",
			[]Token{{Tword, "spoon", "spoon"}},
		},
		{
			"egg and spoon",
			[]Token{{Tword, "egg", "egg"}, {Tword, "and", "and"}, {Tword, "spoon", "spoon"}},
		},
	}
	testTable(t, table)
}

func TestInts(t *testing.T) {
	table := []struct {
		in  string
		out []Token
	}{
		{
			"1",
			[]Token{{Tint, 1, "1"}},
		},
		{
			"427",
			[]Token{{Tint, 427, "427"}},
		},
		{
			"1 2 3",
			[]Token{{Tint, 1, "1"}, {Tint, 2, "2"}, {Tint, 3, "3"}},
		},
	}
	testTable(t, table)
}