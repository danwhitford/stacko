package vm

import (
	"github.com/danwhitford/stacko/stackoval"
)

type InstructionFrame struct {
	Instructions       []stackoval.StackoVal
	Length             int
	InstructionPointer int
}
