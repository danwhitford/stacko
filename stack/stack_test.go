package stack

import (
	"testing"

)

func TestPushBasic(t *testing.T) {
	var stack Stack
	stack.Push(StackoVal{StackoInt, 1})
	stack.Push(StackoVal{StackoFloat, 5.5})
	stack.Push(StackoVal{StackoFloat, 5.5})
}

func TestPopBasic(t *testing.T) {
	tests := []StackoVal{
		{StackoInt, 1},
		{StackoFloat, 5.5},
		{StackoFloat, 5.5}}
	var stack Stack

	for _, tst := range tests {
		stack.Push(tst)
		v, err := stack.Pop()
		if err != nil {
			t.Error(err)
		}
		if v != tst {
			t.Errorf("got: %v\nexpected: %v\n", v, tst)
		}
	}
}

func TestPopMany(t *testing.T) {
	tests := []StackoVal{
		{StackoInt, 1},
		{StackoFloat, 5.5},
		{StackoFloat, 5.5}}
	var stack Stack

	for _, tst := range tests {
		stack.Push(tst)
	}

	for i := len(tests) - 1; i >= 0; i-- {
		v, err := stack.Pop()
		if err != nil {
			t.Error(err)
		}
		if v != tests[i] {
			t.Errorf("got: %v\texpected: %v\n", v, tests[i])
		}
	}
}
