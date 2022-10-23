package tokeniser

import "testing"

func compare(a, b []Token) bool {
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

func TestFloats(t *testing.T) {
	table := []struct {
		in  string
		out []Token
	}{
		{
			"1.0",
			[]Token{{Tfloat, 1.0, "1.0"}},
		},
		{
			"42.7",
			[]Token{{Tfloat, 42.7, "42.7"}},
		},
		{
			"1.1 2.2 3.3",
			[]Token{{Tfloat, 1.1, "1.1"}, {Tfloat, 2.2, "2.2"}, {Tfloat, 3.3, "3.3"}},
		},
	}
	testTable(t, table)
}

func TestStrings(t *testing.T) {
	table := []struct {
		in  string
		out []Token
	}{
		{
			`"foo"`,
			[]Token{{Tstring, "foo", "foo"}},
		},
		{
			`"foo bar"`,
			[]Token{{Tstring, "foo bar", "foo bar"}},
		},
		{
			`"foo" "bar" "baz"`,
			[]Token{{Tstring, "foo", "foo"}, {Tstring, "bar", "bar"}, {Tstring, "baz", "baz"}},
		},
	}
	testTable(t, table)
}