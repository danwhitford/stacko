package vm

import (
	"github.com/danwhitford/stacko/stackoval"
)

type InstructionFrame struct {
	Instructions       []stackoval.StackoVal
	Length             int
	InstructionPointer int
	LoopCounter        int
}

func (frame *InstructionFrame) Advance() error {
	frame.InstructionPointer++
	return nil
}

func (frame *InstructionFrame) Finished() bool {
	if frame.InstructionPointer >= frame.Length {
		if frame.LoopCounter > 0 {
			frame.LoopCounter--
			frame.InstructionPointer = 0
			return false
		}
	}
	return true
}