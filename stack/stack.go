package stack

import (
	"fmt"

)

type stackoType int

const (
	StackoString = iota
	StackoInt
	StackoFloat
	StackoWord
)

type StackoVal struct {
	StackoType stackoType
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
