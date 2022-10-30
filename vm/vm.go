package vm

import (
	"fmt"

	"github.com/danwhitford/stacko/stack"
	"github.com/danwhitford/stacko/stackoval"
)

type VM struct {
	dictionary   map[string][]stackoval.StackoVal
	instructions stack.Stack[InstructionFrame]
	stack        stack.Stack[stackoval.StackoVal]
}

func NewVM() VM {
	return VM{
		make(map[string][]stackoval.StackoVal, 0),
		make(stack.Stack[InstructionFrame], 0),
		make(stack.Stack[stackoval.StackoVal], 0)}
}

func (vm *VM) Load(extras []stackoval.StackoVal) {
	frame := InstructionFrame{
		Instructions:       extras,
		Length:             len(extras),
		InstructionPointer: 0,
	}
	vm.instructions.Push(frame)
}

func (vm *VM) Execute() error {
	for !vm.instructions.Empty() {
		currentFrame, err := vm.instructions.Pop()
		if err != nil {
			return err
		}
		for currentFrame.InstructionPointer < currentFrame.Length {
			switch curr := currentFrame.Instructions[currentFrame.InstructionPointer]; curr.StackoType {
			case stackoval.StackoWord:
				execd, err := vm.execBuiltin(curr.Val.(string))
				if err != nil {
					return fmt.Errorf("error while executing %v: %w", curr, err)
				}
				if !execd {
					return fmt.Errorf("couldn't find definition for word: %s", curr.Val)
				}
			default:
				vm.stack.Push(curr)
			}
			currentFrame.InstructionPointer++
		}
	}
	return nil
}
