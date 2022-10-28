package vm

import (
	"testing"

	"github.com/danwhitford/stacko/stack"
)

func TestMaths(t *testing.T) {
	tests := []struct {
		input    []stack.StackoVal
		expected interface{}
	}{
		{
			input: []stack.StackoVal{
				{StackoType: stack.StackoInt, Val: 1},
				{StackoType: stack.StackoInt, Val: 2},
				{StackoType: stack.StackoWord, Val: "+"}},
			expected: 3,
		},
		{
			input: []stack.StackoVal{
				{StackoType: stack.StackoInt, Val: 1},
				{StackoType: stack.StackoFloat, Val: 5.5},
				{StackoType: stack.StackoWord, Val: "+"}},
			expected: 6.5,
		},
		{
			input: []stack.StackoVal{
				{StackoType: stack.StackoFloat, Val: 1.0},
				{StackoType: stack.StackoInt, Val: 5},
				{StackoType: stack.StackoWord, Val: "+"}},
			expected: 6.0,
		},
		{
			input: []stack.StackoVal{
				{StackoType: stack.StackoFloat, Val: 1.3},
				{StackoType: stack.StackoFloat, Val: 5.5},
				{StackoType: stack.StackoWord, Val: "+"}},
			expected: 6.8,
		},
		{
			input: []stack.StackoVal{
				{StackoType: stack.StackoInt, Val: 7},
				{StackoType: stack.StackoInt, Val: 4},
				{StackoType: stack.StackoWord, Val: "-"}},
			expected: 3,
		},
		{
			input: []stack.StackoVal{
				{StackoType: stack.StackoInt, Val: 7},
				{StackoType: stack.StackoInt, Val: 4},
				{StackoType: stack.StackoWord, Val: "*"}},
			expected: 28,
		},
		{
			input: []stack.StackoVal{
				{StackoType: stack.StackoInt, Val: 10},
				{StackoType: stack.StackoInt, Val: 5},
				{StackoType: stack.StackoWord, Val: "/"}},
			expected: 2,
		},
		{
			input: []stack.StackoVal{
				{StackoType: stack.StackoInt, Val: 10},
				{StackoType: stack.StackoInt, Val: 3},
				{StackoType: stack.StackoWord, Val: "%"}},
			expected: 1,
		},
	}

	for _, tst := range tests {
		vm := NewVM()
		vm.Load(tst.input)
		vm.Execute()

		if vm.stack[0].Val != tst.expected {
			t.Fatalf("can't do maffs for %+v got %+v", tst, vm.stack[0])
		}
	}
}

func ExampleDup() {
	vm := NewVM()
	vm.Load([]stack.StackoVal{
		{StackoType: stack.StackoInt, Val: 1},
		{StackoType: stack.StackoWord, Val: "dup"},
		{StackoType: stack.StackoWord, Val: "v"},
	})
	vm.Execute()
	// Output:
	// 1
	// 1
}

func ExampleSwap() {
	vm := NewVM()
	vm.Load([]stack.StackoVal{
		{StackoType: stack.StackoInt, Val: 1},
		{StackoType: stack.StackoInt, Val: 2},
		{StackoType: stack.StackoWord, Val: "swap"},
		{StackoType: stack.StackoWord, Val: "v"},
	})
	vm.Execute()
	// Output:
	// 1
	// 2
}

func ExampleOver() {
	vm := NewVM()
	vm.Load([]stack.StackoVal{
		{StackoType: stack.StackoInt, Val: 1},
		{StackoType: stack.StackoInt, Val: 2},
		{StackoType: stack.StackoWord, Val: "over"},
		{StackoType: stack.StackoWord, Val: "v"},
	})
	vm.Execute()
	// Output:
	// 1
	// 2
	// 1
}

func ExampleRot() {
	vm := NewVM()
	vm.Load([]stack.StackoVal{
		{StackoType: stack.StackoInt, Val: 1},
		{StackoType: stack.StackoInt, Val: 2},
		{StackoType: stack.StackoInt, Val: 3},
		{StackoType: stack.StackoWord, Val: "rot"},
		{StackoType: stack.StackoWord, Val: "v"},
	})
	vm.Execute()
	// Output:
	// 1
	// 3
	// 2
}

func ExampleDrop() {
	vm := NewVM()
	vm.Load([]stack.StackoVal{
		{StackoType: stack.StackoInt, Val: 1},
		{StackoType: stack.StackoInt, Val: 2},
		{StackoType: stack.StackoWord, Val: "drop"},
		{StackoType: stack.StackoWord, Val: "v"},
	})
	vm.Execute()
	// Output:
	// 1
}

func ExampleVM_PrintStack() {
	vm := NewVM()
	vm.Load([]stack.StackoVal{
		{StackoType: stack.StackoList, Val: []stack.StackoVal{
			{StackoType: stack.StackoString, Val: "foo"},
			{StackoType: stack.StackoString, Val: "bar"},
			{StackoType: stack.StackoString, Val: "baz"},
		}},
		{StackoType: stack.StackoList, Val: []stack.StackoVal{
			{StackoType: stack.StackoString, Val: "foo"},
			{StackoType: stack.StackoList, Val: []stack.StackoVal{
				{StackoType: stack.StackoString, Val: "bar"},
				{StackoType: stack.StackoString, Val: "baz"},
			}}}},
		{StackoType: stack.StackoWord, Val: "v"},
	})
	vm.Execute()
	// Output:
	// [foo [bar baz]]
	// [foo bar baz]
}
