package vm

import (
	"fmt"
	"io"

	"github.com/danwhitford/stacko/stack"
	"github.com/danwhitford/stacko/stackoval"
)

type VM struct {
	dictionary   map[string]stackoval.StackoVal
	instructions stack.Stack[InstructionFrame]
	stack        stack.Stack[stackoval.StackoVal]
	outF         io.Writer
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

	return instruction, nil
}

func (vm *VM) executeInstruction(curr stackoval.StackoVal) error {
	if curr.StackoType == stackoval.StackoWord && curr.Val.(string) == "if" {
		vm.advanceInstruction()
		err := vm.execIf()
		return err
	}

	vm.advanceInstruction()
	switch curr.StackoType {
	case stackoval.StackoWord:
		execd, err := vm.execBuiltin(curr.Val.(string))
		if err != nil {
			return fmt.Errorf("error while executing '%v': %w", curr, err)
		}
		if !execd {
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
		_, err = vm.instructions.Pop()
		if err != nil {
			return fmt.Errorf("error getting next instruction: %w", err)
		}
	}
	return nil
}

func (vm *VM) execIf() error {
	stack := &vm.stack

	falseBranch, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting false branch: %w", err)
	}
	trueBranch, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting true branch: %w", err)
	}
	condition, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting condition: %w", err)
	}
	var branch stackoval.StackoVal
	if condition.Val == true {
		branch = trueBranch
	} else {
		branch = falseBranch
	}

	next := listise(branch)
	vm.Load(next)
	return nil
}
