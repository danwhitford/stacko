package vm

import (
	"fmt"
	"io"
	"os"

	"github.com/danwhitford/stacko/parser"
	"github.com/danwhitford/stacko/stack"
	"github.com/danwhitford/stacko/stackoval"
	"github.com/danwhitford/stacko/tokeniser"
)

type VM struct {
	dictionary  map[string]stackoval.StackoVal
	callStack   stack.Stack[InstructionFrame]
	returnStack stack.Stack[stackoval.StackoVal]
	stack       stack.Stack[stackoval.StackoVal]
	prsr        parser.Parser
	tknsr       tokeniser.Tokeniser
	outF        io.Writer
}

func NewVM(r io.Writer) (VM, error) {
	dictionary := map[string]stackoval.StackoVal{
		"true": {StackoType: stackoval.StackoBool, Val: true},
	}
	instructions := make(stack.Stack[InstructionFrame], 0)
	newVm := VM{
		dictionary,
		instructions,
		make(stack.Stack[stackoval.StackoVal], 0),
		make(stack.Stack[stackoval.StackoVal], 0),
		parser.Parser{},
		tokeniser.Tokeniser{},
		r,
	}

	stdLibLoc := "/Users/danielwhitford/workspace/stacko/stdlib/lib.txt"
	f, err := os.Open(stdLibLoc)
	if err != nil {
		return newVm, err
	}
	libBytes, err := io.ReadAll(f)
	if err != nil {
		return newVm, err
	}
	newVm.tknsr = tokeniser.NewTokeniser(string(libBytes))
	tokens, err := newVm.tknsr.Tokenise()
	if err != nil {
		return newVm, err
	}
	newVm.prsr = parser.NewParser(tokens)
	libCompiled, err := newVm.prsr.Parse()
	if err != nil {
		return newVm, err
	}
	newVm.Load(libCompiled)
	err = newVm.Execute()
	if err != nil {
		return newVm, err
	}

	return newVm, nil
}

func (vm *VM) Reset() {
	vm.dictionary = map[string]stackoval.StackoVal{
		"true": {StackoType: stackoval.StackoBool, Val: true},
	}
	vm.callStack = make(stack.Stack[InstructionFrame], 0)
	vm.stack = make(stack.Stack[stackoval.StackoVal], 0)
}

func (vm *VM) Load(extras []stackoval.StackoVal) {
	frame := NewRegularFrame(extras)
	vm.callStack.Push(frame)
}

func (vm *VM) Execute() error {
	for !vm.callStack.Empty() {
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
	top, err := vm.callStack.Peek()
	if err != nil {
		return stackoval.StackoVal{}, fmt.Errorf("error getting next instruction: %w", err)
	}
	instruction := (*top).GetNext()
	(*top).Advance()
	if (*top).Finished() {
		_, err = vm.callStack.Pop()
		if err != nil {
			return stackoval.StackoVal{}, fmt.Errorf("error getting next instruction: %w", err)
		}
	}
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
