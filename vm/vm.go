package vm

import (
	"fmt"
	"io"

	"github.com/danwhitford/stacko/stack"
	"github.com/danwhitford/stacko/stackoval"
)

type VM struct {
	dictionary       map[string]stackoval.StackoVal
	instructions     stack.Stack[InstructionFrame]
	stack            stack.Stack[stackoval.StackoVal]
	outF             io.Writer
}

func NewVM(r io.Writer) VM {
	dictionary := map[string]stackoval.StackoVal{
		"true": {StackoType: stackoval.StackoBool, Val: true},
	}
	instructions := make(stack.Stack[InstructionFrame], 0)

	return VM{
		dictionary,
		instructions,
		make(stack.Stack[stackoval.StackoVal], 0),
		r}
}

func (vm *VM) Reset() {
	vm.dictionary = map[string]stackoval.StackoVal{
		"true": {StackoType: stackoval.StackoBool, Val: true},
	}
	vm.instructions = make(stack.Stack[InstructionFrame], 0)
	vm.stack = make(stack.Stack[stackoval.StackoVal], 0)
}

func (vm *VM) Load(extras []stackoval.StackoVal) {
	frame := InstructionFrame{
		Instructions:       extras,
		Length:             len(extras),
		InstructionPointer: 0,
		LoopCounter: 0,
	}
	vm.instructions.Push(frame)
}

func (vm *VM) Execute() error {
	for !vm.instructions.Empty() {
		instruction, err := vm.getNextInstruction()
		if err != nil {
			return fmt.Errorf("error getting next instruction %w", err)
		}
		if instruction.StackoType != stackoval.StackoNop {
			err := vm.executeInstruction(instruction)
			if err != nil {
				return fmt.Errorf("error executing instruction %w", err)
			}
		}
	}
	return nil
}

func (vm *VM) getNextInstruction() (stackoval.StackoVal, error) {
	top, err := vm.instructions.Peek()
	if err != nil {
		return stackoval.StackoVal{}, fmt.Errorf("error getting next instruction: %w", err)
	}
	if len(top.Instructions) == 0 {
		_, err = vm.instructions.Pop()
		if err != nil {
			return stackoval.StackoVal{}, fmt.Errorf("error getting next instruction: %w", err)
		}
		return stackoval.StackoVal{StackoType: stackoval.StackoNop}, nil
	}

	instruction := top.Instructions[top.InstructionPointer]
	vm.advanceInstruction()
	return instruction, nil
}

func (vm *VM) executeInstruction(curr stackoval.StackoVal) error {
	switch curr.StackoType {
	case stackoval.StackoWord:
		return vm.execWord(curr)
	default:
		vm.stack.Push(curr)
	}
	return nil
}

func (vm *VM) advanceInstruction() error {
	top, err := vm.instructions.Peek()
	if err != nil {
		return fmt.Errorf("error getting next instruction: %w", err)
	}

	top.Advance()
	if top.InstructionPointer >= top.Length {
		if top.LoopCounter > 0 {
			top.LoopCounter--
			top.InstructionPointer = 0
			return nil
		}
		_, err = vm.instructions.Pop()
		if err != nil {
			return fmt.Errorf("error getting next instruction: %w", err)
		}
	}
	return nil
}

func (vm *VM) execWord(curr stackoval.StackoVal) error {
	// Try builtins first
	execd, err := vm.execBuiltin(curr.Val.(string))
	if err != nil {
		return fmt.Errorf("error while executing '%v': %w", curr, err)
	}
	if !execd {
		// Must be a user defined word
		return vm.execUserWord(curr)
	}
	return nil
}

func (vm *VM) execUserWord(curr stackoval.StackoVal) error {
	userWordDef, prs := vm.dictionary[curr.Val.(string)]
	if !prs {
		return fmt.Errorf("couldn't find definition for word: %s", curr.Val)
	}
	switch userWordDef.StackoType {
	case stackoval.StackoFn:
		vm.Load(userWordDef.Val.([]stackoval.StackoVal))
	default:
		vm.Load([]stackoval.StackoVal{userWordDef})
	}
	return nil
}
