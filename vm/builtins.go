package vm

//go:generate go run builtin_gen.go

import (
	"fmt"
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

func (vm *VM) execBuiltin(word string) (bool, error) {
	switch word {
	case "+":
		return true, vm.Add()
	case "-":
		return true, vm.Sub()
	case "*":
		return true, vm.Mult()
	case "/":
		return true, vm.Div()
	case "%":
		return true, vm.Mod()
	case ".":
		return true, vm.PrintTop()
	case "v":
		return true, vm.PrintStack()
	case "dup":
		return true, vm.Dup()
	case "swap":
		return true, vm.Swap()
	case "over":
		return true, vm.Over()
	case "rot":
		return true, vm.Rot()
	case "drop":
		return true, vm.Drop()
	default:
		return false, nil
	}
}

func (vm *VM) Swap() error {
	stack := &vm.stack
	a, err := stack.Pop()
	if err != nil {
		return err
	}
	b, err := stack.Pop()
	if err != nil {
		return err
	}
	stack.Push(a)
	stack.Push(b)
	return nil
}

func (vm *VM) Over() error {
	stack := &vm.stack
	a, err := stack.Pop()
	if err != nil {
		return err
	}
	b, err := stack.Pop()
	if err != nil {
		return err
	}
	stack.Push(b)
	stack.Push(a)
	stack.Push(b)
	return nil
}

func (vm *VM) Rot() error {
	stack := &vm.stack
	a, err := stack.Pop()
	if err != nil {
		return err
	}
	b, err := stack.Pop()
	if err != nil {
		return err
	}
	c, err := stack.Pop()
	if err != nil {
		return err
	}
	stack.Push(b)
	stack.Push(a)
	stack.Push(c)
	return nil
}

func (vm *VM) Drop() error {
	stack := &vm.stack
	_, err := stack.Pop()
	return err
}
