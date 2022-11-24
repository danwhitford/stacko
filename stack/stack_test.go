package stack

import (
	"testing"

	"github.com/danwhitford/stacko/stackoval"
)

func TestPushBasic(t *testing.T) {
	var stack Stack[stackoval.StackoVal]
	stack.Push(stackoval.StackoVal{stackoval.StackoInt, 1})
	stack.Push(stackoval.StackoVal{stackoval.StackoFloat, 5.5})
	stack.Push(stackoval.StackoVal{stackoval.StackoFloat, 5.5})
}

func TestPopBasic(t *testing.T) {
	tests := []stackoval.StackoVal{
		{stackoval.StackoInt, 1},
		{stackoval.StackoFloat, 5.5},
		{stackoval.StackoFloat, 5.5}}
	var stack Stack[stackoval.StackoVal]

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
	tests := []stackoval.StackoVal{
		{stackoval.StackoInt, 1},
		{stackoval.StackoFloat, 5.5},
		{stackoval.StackoFloat, 5.5}}
	var stack Stack[stackoval.StackoVal]

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

func TestPeek(t *testing.T) {
	tests := []stackoval.StackoVal{
		{stackoval.StackoInt, 1},
		{stackoval.StackoFloat, 3.5},
		{stackoval.StackoFloat, 5.5}}
	var stack Stack[stackoval.StackoVal]

	for _, tst := range tests {
		stack.Push(tst)
	}
	peeked, err := stack.Peek()
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 10; i++ {
		if peeked.Val.(float64) != 5.5 {
			t.Errorf("wanted 5.5 got %v", peeked)
		}
	}
}

func TestEmpty(t *testing.T) {
	var stack Stack[int]
	if !stack.Empty() {
		t.Fatal("should be empty")
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Empty() {
		t.Fatal("should be unempty")
	}
	stack.Pop()
	stack.Pop()
	stack.Pop()
	if !stack.Empty() {
		t.Fatal("should be empty")
	}
}
