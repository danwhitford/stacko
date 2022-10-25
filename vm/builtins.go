package vm

import (
	"github.com/danwhitford/stacko/stack"
)

var builtins = map[string]func(*stack.Stack) error {
	"+": func(stk *stack.Stack) error {
		a, err := stk.Pop()
		if err != nil {
			return err
		}
		b, err := stk.Pop()
		if err != nil {
			return err
		}

		switch {
		case a.StackoType == stack.StackoInt && b.StackoType == stack.StackoInt:
			stk.Push(stack.StackoVal{StackoType: stack.StackoInt, Val: a.Val.(int) + b.Val.(int)})
		}
		return nil
	},
}