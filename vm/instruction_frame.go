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

type LoopInstructionFrame struct {
	Instructions       []stackoval.StackoVal
	Length             int
	InstructionPointer int
	LoopCounter        int
}

func NewLoopInstructionFrame(instructions []stackoval.StackoVal, counter int) *LoopInstructionFrame {
	return &LoopInstructionFrame{
		Instructions:       instructions,
		Length:             len(instructions),
		InstructionPointer: 0,
		LoopCounter:        counter,
	}
}

func (frame *LoopInstructionFrame) Finished() bool {
	if frame.InstructionPointer >= frame.Length {
		if frame.LoopCounter > 0 {
			frame.LoopCounter--
			frame.InstructionPointer = 0
			return false
		}

		return true
	}
	return false
}

func (frame *LoopInstructionFrame) GetNext() stackoval.StackoVal {
	instruction := frame.Instructions[frame.InstructionPointer]
	return instruction
}

func (frame *LoopInstructionFrame) Advance() error {
	frame.InstructionPointer++
	return nil
}
