package vm

import (
	"fmt"

	"github.com/danwhitford/stacko/stack"
	"github.com/danwhitford/stacko/stackoval"
)

type VM struct {
	dictionary   map[string]stackoval.StackoVal
	instructions stack.Stack[InstructionFrame]
	stack        stack.Stack[stackoval.StackoVal]
}

func NewVM() VM {
	dictionary := map[string]stackoval.StackoVal{
		"true": {StackoType: stackoval.StackoBool, Val: true},
	}
	instructions := make(stack.Stack[InstructionFrame], 0)

	return VM{
		dictionary,
		instructions,
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
		err := vm.executeInstructionFrame()
		if err != nil {
			return err
		}
	}
	return nil
}

func (vm *VM) executeInstructionFrame() error {
	currentFrame, err := vm.instructions.Peek()

	if err != nil {
		return err
	}

	for currentFrame.InstructionPointer < currentFrame.Length {
		switch curr := currentFrame.Instructions[currentFrame.InstructionPointer]; curr.StackoType {
		case stackoval.StackoWord:
			execd, err := vm.execBuiltin(curr.Val.(string))
			if err != nil {
				currentFrame.InstructionPointer++
				if _, ok := err.(*DoNotPop); ok {
					currentFrame.InstructionPointer = 1000
					vm.instructions.Swap()
					vm.instructions.Pop()
					vm.instructions.Push(*currentFrame)
					vm.instructions.Swap()
					return nil
				}
				return fmt.Errorf("error while executing '%v': %w", curr, err)
			}
			if !execd {
				userWord, prs := vm.dictionary[curr.Val.(string)]
				if !prs {
					currentFrame.InstructionPointer++
					return fmt.Errorf("couldn't find definition for word: %s", curr.Val)
				}
				switch userWord.StackoType {
				case stackoval.StackoList:
					frame := InstructionFrame{
						userWord.Val.([]stackoval.StackoVal),
						len(userWord.Val.([]stackoval.StackoVal)),
						0,
					}
					currentFrame.InstructionPointer++
					vm.instructions.Push(frame)
				default:
					frame := InstructionFrame{
						[]stackoval.StackoVal{userWord},
						1,
						0,
					}
					currentFrame.InstructionPointer++
					vm.instructions.Push(frame)
				}
				return nil
			}
		default:
			vm.stack.Push(curr)
		}

		currentFrame.InstructionPointer++
	}

	vm.instructions.Pop()
	return nil
}
