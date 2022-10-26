package stack

import (
	"fmt"

)

//go:generate stringer -type=StackoType
type StackoType int

const (
	StackoString StackoType = iota
	StackoInt
	StackoFloat
	StackoWord
	StackoList
)

type StackoVal struct {
	StackoType StackoType
	Val        interface{}
}

type Stack []StackoVal

func (stack *Stack) Push(v StackoVal) {
	*stack = append(*stack, v)
}

func (stack *Stack) Pop() (StackoVal, error) {
	l := len(*stack)
	if l < 1 {
		return StackoVal{}, fmt.Errorf("stack underflow")
	}

	v := (*stack)[l-1]
	*stack = (*stack)[:l-1]
	return v, nil
}
