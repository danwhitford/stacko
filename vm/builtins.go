package vm
//go:generate go run builtin_gen.go

import (
	"fmt"

	"github.com/danwhitford/stacko/stack"
)

func Drop(stk *stack.Stack) error {
	a, err := stk.Pop()
	if err != nil {
		return err
	}
	fmt.Printf("%v\n", a.Val)
	return nil
}

var builtins = map[string]func(*stack.Stack) error{
	"+": Add,
	"-": Sub,
	"*": Mult,
	"/": Div,
	".": Drop,
	"v": func(stk *stack.Stack) error {
		for i := len(*stk) - 1; i >= 0; i-- {
			fmt.Printf("%v\n", (*stk)[i].Val)
		}
		return nil
	},
	"dup": func(stk *stack.Stack) error {
		a, err := stk.Pop()
		if err != nil {
			return err
		}
		stk.Push(a)
		stk.Push(a)
		return nil
	},
	"swap": func(stk *stack.Stack) error {
		a, err := stk.Pop()
		if err != nil {
			return err
		}
		b, err := stk.Pop()
		if err != nil {
			return err
		}
		stk.Push(a)
		stk.Push(b)
		return nil
	},
	"over": func(stk *stack.Stack) error {
		a, err := stk.Pop()
		if err != nil {
			return err
		}
		b, err := stk.Pop()
		if err != nil {
			return err
		}
		stk.Push(b)
		stk.Push(a)
		stk.Push(b)
		return nil
	},
	"rot": func(stk *stack.Stack) error {
		a, err := stk.Pop()
		if err != nil {
			return err
		}
		b, err := stk.Pop()
		if err != nil {
			return err
		}
		c, err := stk.Pop()
		if err != nil {
			return err
		}
		stk.Push(b)
		stk.Push(a)
		stk.Push(c)
		return nil
	},
	"drop": func(stk *stack.Stack) error {
		_, err := stk.Pop()
		return err
	},
}
