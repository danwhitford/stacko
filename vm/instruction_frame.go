package vm

import (
	"github.com/danwhitford/stacko/stackoval"
)

type InstructionFrame interface {
	GetNext() stackoval.StackoVal
	Advance() error
	Finished() bool
}

type RegularFrame struct {
	Instructions       []stackoval.StackoVal
	Length             int
	InstructionPointer int
}

func NewRegularFrame(instructions []stackoval.StackoVal) *RegularFrame {
	return &RegularFrame{
		Instructions:       instructions,
		Length:             len(instructions),
		InstructionPointer: 0,
	}
}

func (frame *RegularFrame) Advance() error {
	frame.InstructionPointer++
	return nil
}

func (frame *RegularFrame) Finished() bool {
	return frame.InstructionPointer >= frame.Length
}

func (frame *RegularFrame) GetNext() stackoval.StackoVal {
	instruction := frame.Instructions[frame.InstructionPointer]
	return instruction
}
