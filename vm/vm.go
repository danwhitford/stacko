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
		instruction, err := vm.getNextInstruction()
		if err != nil {
			return err
		}
		err = vm.executeInstruction(instruction)
		if err != nil {
				return err		
		}
	}
	return nil
}

func (vm *VM) getNextInstruction() (stackoval.StackoVal, error) {
	top, err := vm.instructions.Peek()
	if err != nil {
		return stackoval.StackoVal{}, fmt.Errorf("error getting next instruction: %w", err)
	}
	instruction := top.Instructions[top.InstructionPointer]
	top.Advance()
	if top.InstructionPointer >= top.Length {
		vm.instructions.Pop()
	}
	return instruction, nil
}

func (vm *VM) executeInstruction(curr stackoval.StackoVal) error {
	switch curr.StackoType {
	case stackoval.StackoWord:
		execd, err := vm.execBuiltin(curr.Val.(string))
		if err != nil {
			return fmt.Errorf("error while executing '%v': %w", curr, err)
		}
		if !execd {
			userWord, prs := vm.dictionary[curr.Val.(string)]
			if !prs {
				return fmt.Errorf("couldn't find definition for word: %s", curr.Val)
			}
			switch userWord.StackoType {
			case stackoval.StackoList:
				frame := InstructionFrame{
					userWord.Val.([]stackoval.StackoVal),
					len(userWord.Val.([]stackoval.StackoVal)),
					0,
				}
				vm.instructions.Push(frame)
			default:
				frame := InstructionFrame{
					[]stackoval.StackoVal{userWord},
					1,
					0,
				}
				vm.instructions.Push(frame)
			}
			return nil
		}
	default:
		vm.stack.Push(curr)
	}
	return nil
}
