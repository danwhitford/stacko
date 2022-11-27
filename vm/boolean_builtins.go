package vm

import "fmt"

import "github.com/danwhitford/stacko/stackoval"

func (vm *VM) Eq() error {
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	b, err := vm.stack.Pop()
	if err != nil {
		return err
	}

	switch {
	case a.StackoType == stackoval.StackoInt && b.StackoType == stackoval.StackoInt:
		eq := b.Val.(int) == a.Val.(int)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoBool, Val: eq})
	case a.StackoType == stackoval.StackoString && b.StackoType == stackoval.StackoString:
		eq := b.Val.(string) == a.Val.(string)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoBool, Val: eq})
	default:
		return fmt.Errorf("== not defined for %s and %s", a.StackoType, b.StackoType)
	}

	// vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoBool, Val: eq})

	return nil
}

func (vm *VM) Gt() error {
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	b, err := vm.stack.Pop()
	if err != nil {
		return err
	}

	switch {
	case a.StackoType == stackoval.StackoInt && b.StackoType == stackoval.StackoInt:
		eq := b.Val.(int) > a.Val.(int)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoBool, Val: eq})
	case a.StackoType == stackoval.StackoString && b.StackoType == stackoval.StackoString:
		eq := b.Val.(string) > a.Val.(string)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoBool, Val: eq})
	default:
		return fmt.Errorf("> not defined for %s and %s", a.StackoType, b.StackoType)
	}

	// vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoBool, Val: eq})

	return nil
}

func (vm *VM) Lt() error {
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	b, err := vm.stack.Pop()
	if err != nil {
		return err
	}

	switch {
	case a.StackoType == stackoval.StackoInt && b.StackoType == stackoval.StackoInt:
		eq := b.Val.(int) < a.Val.(int)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoBool, Val: eq})
	case a.StackoType == stackoval.StackoString && b.StackoType == stackoval.StackoString:
		eq := b.Val.(string) < a.Val.(string)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoBool, Val: eq})
	default:
		return fmt.Errorf("< not defined for %s and %s", a.StackoType, b.StackoType)
	}

	// vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoBool, Val: eq})

	return nil
}
