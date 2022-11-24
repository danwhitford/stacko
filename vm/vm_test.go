package vm

import (
	"fmt"
	"testing"

	"github.com/danwhitford/stacko/stackoval"
	"github.com/google/go-cmp/cmp"
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

func ExampleDefVar() {
	vm := NewVM()
	vm.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 5},
		{StackoType: stackoval.StackoSymbol, Val: "foo"},
		{StackoType: stackoval.StackoWord, Val: "def"},
		{StackoType: stackoval.StackoWord, Val: "foo"},
		{StackoType: stackoval.StackoWord, Val: "foo"},
		{StackoType: stackoval.StackoWord, Val: "+"},
		{StackoType: stackoval.StackoWord, Val: "v"},
	})
	err := vm.Execute()
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 10
}

func ExamplePrintTop() {
	vm := NewVM()
	vm.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 100},
		{StackoType: stackoval.StackoInt, Val: 20},
		{StackoType: stackoval.StackoWord, Val: "/"},
		{StackoType: stackoval.StackoWord, Val: "."},
	})
	err := vm.Execute()
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 5
}

func TestIfStuff(t *testing.T) {
	table := []struct {
		in  []stackoval.StackoVal
		out stackoval.StackoVal
	}{
		{in: []stackoval.StackoVal{
			{StackoType: stackoval.StackoBool, Val: true},
			{StackoType: stackoval.StackoString, Val: "yes"},
			{StackoType: stackoval.StackoString, Val: "no"},
			{StackoType: stackoval.StackoWord, Val: "if"},
		},
			out: stackoval.StackoVal{StackoType: stackoval.StackoString, Val: "yes"},
		},
		{in: []stackoval.StackoVal{
			{StackoType: stackoval.StackoBool, Val: false},
			{StackoType: stackoval.StackoList, Val: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 10},
				{StackoType: stackoval.StackoInt, Val: 5},
				{StackoType: stackoval.StackoWord, Val: "*"},
			}},
			{StackoType: stackoval.StackoList, Val: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 10},
				{StackoType: stackoval.StackoInt, Val: 5},
				{StackoType: stackoval.StackoWord, Val: "+"},
			}},
			{StackoType: stackoval.StackoWord, Val: "if"},
		},
			out: stackoval.StackoVal{StackoType: stackoval.StackoInt, Val: 15},
		},
		{in: []stackoval.StackoVal{
			{StackoType: stackoval.StackoInt, Val: 10},
			{StackoType: stackoval.StackoInt, Val: 10},
			{StackoType: stackoval.StackoBool, Val: false},
			{StackoType: stackoval.StackoList, Val: []stackoval.StackoVal{
				{StackoType: stackoval.StackoWord, Val: "*"},
			}},
			{StackoType: stackoval.StackoList, Val: []stackoval.StackoVal{
				{StackoType: stackoval.StackoWord, Val: "+"},
			}},
			{StackoType: stackoval.StackoWord, Val: "if"},
		},
			out: stackoval.StackoVal{StackoType: stackoval.StackoInt, Val: 20},
		},
		{in: []stackoval.StackoVal{
			{StackoType: stackoval.StackoInt, Val: 10},
			{StackoType: stackoval.StackoInt, Val: 10},
			{StackoType: stackoval.StackoBool, Val: true},
			{StackoType: stackoval.StackoList, Val: []stackoval.StackoVal{
				{StackoType: stackoval.StackoWord, Val: "'*"},
			}},
			{StackoType: stackoval.StackoList, Val: []stackoval.StackoVal{
				{StackoType: stackoval.StackoWord, Val: "'+"},
			}},
			{StackoType: stackoval.StackoWord, Val: "if"},
		},
			out: stackoval.StackoVal{StackoType: stackoval.StackoInt, Val: 100},
		},
	}

	for _, test := range table {
		vm := NewVM()
		vm.Load(test.in)
		err := vm.Execute()
		if err != nil {
			t.Error(err)
		}
		top, err := vm.stack.Peek()
		if err != nil {
			t.Error(err)
		}
		diff := cmp.Diff(test.out, *top)
		if diff != "" {
			t.Errorf("(-want +got):\n%s", diff)
		}
	}
}
