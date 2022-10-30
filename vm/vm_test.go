package vm

import (
	"testing"

	"github.com/danwhitford/stacko/stackoval"
)

func TestMaths(t *testing.T) {
	tests := []struct {
		input    []stackoval.StackoVal
		expected interface{}
	}{
		{
			input: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 1},
				{StackoType: stackoval.StackoInt, Val: 2},
				{StackoType: stackoval.StackoWord, Val: "+"}},
			expected: 3,
		},
		{
			input: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 1},
				{StackoType: stackoval.StackoFloat, Val: 5.5},
				{StackoType: stackoval.StackoWord, Val: "+"}},
			expected: 6.5,
		},
		{
			input: []stackoval.StackoVal{
				{StackoType: stackoval.StackoFloat, Val: 1.0},
				{StackoType: stackoval.StackoInt, Val: 5},
				{StackoType: stackoval.StackoWord, Val: "+"}},
			expected: 6.0,
		},
		{
			input: []stackoval.StackoVal{
				{StackoType: stackoval.StackoFloat, Val: 1.3},
				{StackoType: stackoval.StackoFloat, Val: 5.5},
				{StackoType: stackoval.StackoWord, Val: "+"}},
			expected: 6.8,
		},
		{
			input: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 7},
				{StackoType: stackoval.StackoInt, Val: 4},
				{StackoType: stackoval.StackoWord, Val: "-"}},
			expected: 3,
		},
		{
			input: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 7},
				{StackoType: stackoval.StackoInt, Val: 4},
				{StackoType: stackoval.StackoWord, Val: "*"}},
			expected: 28,
		},
		{
			input: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 10},
				{StackoType: stackoval.StackoInt, Val: 5},
				{StackoType: stackoval.StackoWord, Val: "/"}},
			expected: 2,
		},
		{
			input: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 10},
				{StackoType: stackoval.StackoInt, Val: 3},
				{StackoType: stackoval.StackoWord, Val: "%"}},
			expected: 1,
		},
	}

	for _, tst := range tests {
		t.Log(tst)
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
	vm.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 1},
		{StackoType: stackoval.StackoWord, Val: "dup"},
		{StackoType: stackoval.StackoWord, Val: "v"},
	})
	vm.Execute()
	// Output:
	// 1
	// 1
}

func ExampleSwap() {
	vm := NewVM()
	vm.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 1},
		{StackoType: stackoval.StackoInt, Val: 2},
		{StackoType: stackoval.StackoWord, Val: "swap"},
		{StackoType: stackoval.StackoWord, Val: "v"},
	})
	vm.Execute()
	// Output:
	// 1
	// 2
}

func ExampleOver() {
	vm := NewVM()
	vm.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 1},
		{StackoType: stackoval.StackoInt, Val: 2},
		{StackoType: stackoval.StackoWord, Val: "over"},
		{StackoType: stackoval.StackoWord, Val: "v"},
	})
	vm.Execute()
	// Output:
	// 1
	// 2
	// 1
}

func ExampleRot() {
	vm := NewVM()
	vm.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 1},
		{StackoType: stackoval.StackoInt, Val: 2},
		{StackoType: stackoval.StackoInt, Val: 3},
		{StackoType: stackoval.StackoWord, Val: "rot"},
		{StackoType: stackoval.StackoWord, Val: "v"},
	})
	vm.Execute()
	// Output:
	// 1
	// 3
	// 2
}

func ExampleDrop() {
	vm := NewVM()
	vm.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 1},
		{StackoType: stackoval.StackoInt, Val: 2},
		{StackoType: stackoval.StackoWord, Val: "drop"},
		{StackoType: stackoval.StackoWord, Val: "v"},
	})
	vm.Execute()
	// Output:
	// 1
}

func ExampleVM_PrintStack() {
	vm := NewVM()
	vm.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoList, Val: []stackoval.StackoVal{
			{StackoType: stackoval.StackoString, Val: "foo"},
			{StackoType: stackoval.StackoString, Val: "bar"},
			{StackoType: stackoval.StackoString, Val: "baz"},
		}},
		{StackoType: stackoval.StackoList, Val: []stackoval.StackoVal{
			{StackoType: stackoval.StackoString, Val: "foo"},
			{StackoType: stackoval.StackoList, Val: []stackoval.StackoVal{
				{StackoType: stackoval.StackoString, Val: "bar"},
				{StackoType: stackoval.StackoString, Val: "baz"},
			}}}},
		{StackoType: stackoval.StackoWord, Val: "v"},
	})
	vm.Execute()
	// Output:
	// [foo [bar baz]]
	// [foo bar baz]
}
