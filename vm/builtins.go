package vm

//go:generate go run builtin_gen.go

import (
	"fmt"

	"github.com/danwhitford/stacko/stack"
)

func (vm *VM) PrintTop() error {
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	fmt.Println(a)
	return nil
}

func (vm *VM) PrintStack() error {
	for i := len(vm.stack) - 1; i >= 0; i-- {
		fmt.Println(vm.stack[i])
	}
	return nil
}

func (vm *VM) Dup() error {
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	vm.stack.Push(a)
	vm.stack.Push(a)
	return nil
}

func (vm *VM) execBuiltin() (bool, error) {
	switch vm.instructions[vm.instructionPtr].Val.(string) {
	case "+":
		return true, vm.Add()
	case "-":
		return true, vm.Sub()
	case "*":
		return true, vm.Mult()
	case "/":
		return true, vm.Div()
	case ".":
		return true, vm.PrintTop()
	case "v":
		return true, vm.PrintStack()
	case "dup":
		return true, vm.Dup()
	default:
		return false, nil
	}
}

var builtins = map[string]func(*stack.Stack) error{
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
