package vm

import (
	"fmt"
	"os"
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
		machine, err := NewVM(os.Stdout)
		if err != nil {
			t.Fatal(err)
		}
		machine.Load(tst.input)
		err = machine.Execute()
		if err != nil {
			t.Fatal(err)
		}

		if machine.stack[0].Val != tst.expected {
			t.Fatalf("can't do maffs for %+v got %+v", tst, machine.stack[0])
		}
	}
}

func ExampleDup() {
	machine, _ := NewVM(os.Stdout)
	machine.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 1},
		{StackoType: stackoval.StackoWord, Val: "dup"},
		{StackoType: stackoval.StackoWord, Val: "v"},
	})
	machine.Execute()
	// Output:
	// 1
	// 1
}

func ExampleSwap() {
	machine, _ := NewVM(os.Stdout)
	machine.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 1},
		{StackoType: stackoval.StackoInt, Val: 2},
		{StackoType: stackoval.StackoWord, Val: "swap"},
		{StackoType: stackoval.StackoWord, Val: "v"},
	})
	machine.Execute()
	// Output:
	// 1
	// 2
}

func ExampleOver() {
	machine, _ := NewVM(os.Stdout)
	machine.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 1},
		{StackoType: stackoval.StackoInt, Val: 2},
		{StackoType: stackoval.StackoWord, Val: "over"},
		{StackoType: stackoval.StackoWord, Val: "v"},
	})
	machine.Execute()
	// Output:
	// 1
	// 2
	// 1
}

func ExampleRot() {
	machine, _ := NewVM(os.Stdout)
	machine.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 1},
		{StackoType: stackoval.StackoInt, Val: 2},
		{StackoType: stackoval.StackoInt, Val: 3},
		{StackoType: stackoval.StackoWord, Val: "rot"},
		{StackoType: stackoval.StackoWord, Val: "v"},
	})
	machine.Execute()
	// Output:
	// 1
	// 3
	// 2
}

func ExampleDrop() {
	machine, _ := NewVM(os.Stdout)
	machine.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 1},
		{StackoType: stackoval.StackoInt, Val: 2},
		{StackoType: stackoval.StackoWord, Val: "drop"},
		{StackoType: stackoval.StackoWord, Val: "v"},
	})
	machine.Execute()
	// Output:
	// 1
}

func ExampleVM_PrintStack() {
	machine, _ := NewVM(os.Stdout)
	machine.Load([]stackoval.StackoVal{
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
	machine.Execute()
	// Output:
	// [foo [bar baz]]
	// [foo bar baz]
}

func ExampleDefVar() {
	machine, _ := NewVM(os.Stdout)
	machine.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 5},
		{StackoType: stackoval.StackoSymbol, Val: "foo"},
		{StackoType: stackoval.StackoWord, Val: "def"},
		{StackoType: stackoval.StackoWord, Val: "foo"},
		{StackoType: stackoval.StackoWord, Val: "foo"},
		{StackoType: stackoval.StackoWord, Val: "+"},
		{StackoType: stackoval.StackoWord, Val: "v"},
	})
	err := machine.Execute()
	if err != nil {
		fmt.Println(err)
	}
	// Output:
	// 10
}

func ExamplePrintTop() {
	machine, _ := NewVM(os.Stdout)
	machine.Load([]stackoval.StackoVal{
		{StackoType: stackoval.StackoInt, Val: 100},
		{StackoType: stackoval.StackoInt, Val: 20},
		{StackoType: stackoval.StackoWord, Val: "/"},
		{StackoType: stackoval.StackoWord, Val: "."},
	})
	err := machine.Execute()
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
			{StackoType: stackoval.StackoString, Val: "yes"},
			{StackoType: stackoval.StackoString, Val: "no"},
			{StackoType: stackoval.StackoBool, Val: true},
			{StackoType: stackoval.StackoWord, Val: "branch"},
		},
			out: stackoval.StackoVal{StackoType: stackoval.StackoString, Val: "yes"},
		},
		{in: []stackoval.StackoVal{
			{StackoType: stackoval.StackoString, Val: "yes"},
			{StackoType: stackoval.StackoString, Val: "no"},
			{StackoType: stackoval.StackoBool, Val: false},
			{StackoType: stackoval.StackoWord, Val: "branch"},
		},
			out: stackoval.StackoVal{StackoType: stackoval.StackoString, Val: "no"},
		},
		{in: []stackoval.StackoVal{
			{StackoType: stackoval.StackoBool, Val: true},
			{StackoType: stackoval.StackoFn, Val: []stackoval.StackoVal{
				{StackoType: stackoval.StackoString, Val: "yes"},
			}},
			{StackoType: stackoval.StackoFn, Val: []stackoval.StackoVal{
				{StackoType: stackoval.StackoString, Val: "no"},
			}},
			{StackoType: stackoval.StackoWord, Val: "if"},
		},
			out: stackoval.StackoVal{StackoType: stackoval.StackoString, Val: "yes"},
		},
		{in: []stackoval.StackoVal{
			{StackoType: stackoval.StackoBool, Val: false},
			{StackoType: stackoval.StackoFn, Val: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 10},
				{StackoType: stackoval.StackoInt, Val: 5},
				{StackoType: stackoval.StackoWord, Val: "*"},
			}},
			{StackoType: stackoval.StackoFn, Val: []stackoval.StackoVal{
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
			{StackoType: stackoval.StackoFn, Val: []stackoval.StackoVal{
				{StackoType: stackoval.StackoWord, Val: "*"},
			}},
			{StackoType: stackoval.StackoFn, Val: []stackoval.StackoVal{
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
			{StackoType: stackoval.StackoFn, Val: []stackoval.StackoVal{
				{StackoType: stackoval.StackoWord, Val: "*"},
			}},
			{StackoType: stackoval.StackoFn, Val: []stackoval.StackoVal{
				{StackoType: stackoval.StackoWord, Val: "+"},
			}},
			{StackoType: stackoval.StackoWord, Val: "if"},
		},
			out: stackoval.StackoVal{StackoType: stackoval.StackoInt, Val: 100},
		},
		{in: []stackoval.StackoVal{
			{StackoType: stackoval.StackoInt, Val: 10},
			{StackoType: stackoval.StackoBool, Val: true},
			{StackoType: stackoval.StackoFn, Val: []stackoval.StackoVal{{StackoType: stackoval.StackoNop}}},
			{StackoType: stackoval.StackoSymbol, Val: "nonsuch"},
			{StackoType: stackoval.StackoWord, Val: "if"},
		},
			out: stackoval.StackoVal{StackoType: stackoval.StackoInt, Val: 10},
		},
	}

	testTable(t, table)
}

func TestBuiltIns(t *testing.T) {
	table := []struct {
		in  []stackoval.StackoVal
		out stackoval.StackoVal
	}{
		{
			in: []stackoval.StackoVal{
				{StackoType: stackoval.StackoString, Val: "egg"},
				{StackoType: stackoval.StackoString, Val: "egg"},
				{StackoType: stackoval.StackoWord, Val: "="},
			},
			out: stackoval.StackoVal{StackoType: stackoval.StackoBool, Val: true},
		},
		{
			in: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 3},
				{StackoType: stackoval.StackoInt, Val: 4},
				{StackoType: stackoval.StackoWord, Val: "<"},
			},
			out: stackoval.StackoVal{StackoType: stackoval.StackoBool, Val: true},
		},
		{
			in: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 3},
				{StackoType: stackoval.StackoInt, Val: 4},
				{StackoType: stackoval.StackoWord, Val: ">"},
			},
			out: stackoval.StackoVal{StackoType: stackoval.StackoBool, Val: false},
		},
	}

	testTable(t, table)
}

func TestLoops(t *testing.T) {
	table := []struct {
		in  []stackoval.StackoVal
		out stackoval.StackoVal
	}{
		{
			in: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 10},
				{StackoType: stackoval.StackoWord, Val: "range"},
			},
			out: stackoval.StackoVal{
				StackoType: stackoval.StackoList,
				Val: []stackoval.StackoVal{
					{stackoval.StackoInt, 0},
					{stackoval.StackoInt, 1},
					{stackoval.StackoInt, 2},
					{stackoval.StackoInt, 3},
					{stackoval.StackoInt, 4},
					{stackoval.StackoInt, 5},
					{stackoval.StackoInt, 6},
					{stackoval.StackoInt, 7},
					{stackoval.StackoInt, 8},
					{stackoval.StackoInt, 9},
				},
			},
		},
		{
			in: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 0},
				{StackoType: stackoval.StackoFn, Val: []stackoval.StackoVal{
					{StackoType: stackoval.StackoInt, Val: 1},
					{StackoType: stackoval.StackoWord, Val: "+"},
				}},
				{StackoType: stackoval.StackoInt, Val: 5},
				{StackoType: stackoval.StackoWord, Val: "times"},
			},
			out: stackoval.StackoVal{
				StackoType: stackoval.StackoInt,
				Val:        5,
			},
		},
		{
			in: []stackoval.StackoVal{
				{StackoType: stackoval.StackoInt, Val: 0},
				{
					StackoType: stackoval.StackoList,
					Val: []stackoval.StackoVal{
						{StackoType: stackoval.StackoInt, Val: 0},
						{StackoType: stackoval.StackoInt, Val: 1},
						{StackoType: stackoval.StackoInt, Val: 2},
						{StackoType: stackoval.StackoInt, Val: 3},
						{StackoType: stackoval.StackoInt, Val: 4},
						{StackoType: stackoval.StackoInt, Val: 5},
					},
				},
				{StackoType: stackoval.StackoSymbol, Val: "+"},
				{StackoType: stackoval.StackoWord, Val: "each"},
			},
			out: stackoval.StackoVal{
				StackoType: stackoval.StackoInt,
				Val:        15,
			},
		},
		{
			in: []stackoval.StackoVal{
				{
					StackoType: stackoval.StackoFn,
					Val: []stackoval.StackoVal{
						{StackoType: stackoval.StackoInt, Val: 5},
						{StackoType: stackoval.StackoInt, Val: 4},
						{StackoType: stackoval.StackoWord, Val: "+"},
					},
				},
				{StackoType: stackoval.StackoWord, Val: "call"},
			},
			out: stackoval.StackoVal{
				StackoType: stackoval.StackoInt,
				Val:        9,
			},
		},
	}

	testTable(t, table)
}

func testTable(t *testing.T, table []struct {
	in  []stackoval.StackoVal
	out stackoval.StackoVal
}) {
	for _, test := range table {
		machine, err := NewVM(os.Stdout)
		if err != nil {
			t.Fatal(err)
		}
		machine.Load(test.in)
		err = machine.Execute()
		if err != nil {
			t.Fatalf("%s\n%s", test.in, err)
		}
		top, err := machine.stack.Peek()
		if err != nil {
			t.Fatal(err)
		}
		diff := cmp.Diff(test.out, *top)
		if diff != "" {
			t.Fatalf("failed %s\n(-want +got):\n%s", test.in, diff)
		}
		if !machine.returnStack.Empty() {
			t.Fatalf("should leave return stack empty but contains %v for %v", machine.returnStack, test.in)
		}
	}
}
