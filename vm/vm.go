package vm

import (
	"fmt"

	"github.com/danwhitford/stacko/stack"
)

type VM struct {
	builtins       map[string]func(*stack.Stack) error
	userDict       map[string]stack.StackoVal
	instructions   []stack.StackoVal
	instructionPtr int
	len            int
	stack          stack.Stack
}

func NewVM() VM {
	return VM{
		builtins,
		make(map[string]stack.StackoVal),
		make([]stack.StackoVal, 0),
		0,
		0,
		make(stack.Stack, 0)}
}

func (vm *VM) Load(extras []stack.StackoVal) {
	vm.instructions = append(vm.instructions, extras...)
	vm.len += len(extras)
}

func (vm *VM) Execute() error {
	for vm.instructionPtr < vm.len {
		switch curr := vm.instructions[vm.instructionPtr]; curr.StackoType {
		case stack.StackoInt, stack.StackoFloat, stack.StackoString:
			vm.stack.Push(curr)
		case stack.StackoWord:
			f, ok := vm.builtins[curr.Val.(string)]
			if !ok {
				vm.instructionPtr++
				return fmt.Errorf("could not find word in dict: %s", curr.Val)
			}
			err := f(&vm.stack)			
			if err != nil {
				vm.instructionPtr++
				return fmt.Errorf("error while executing %v: %w", curr, err)
			}
		}
		vm.instructionPtr++
	}
	return nil
}
