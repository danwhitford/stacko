package stack

import "fmt"

type valType int

const (
	Tint = iota
	Tfloat
	Tstring
	Tword
)

type stackFrame struct {
	V  interface{}
	VT valType
}

type Stack []stackFrame

func (stack *Stack) Push(v stackFrame) {
	*stack = append(*stack, v)
}

func (stack *Stack) Pop() (stackFrame, error) {
	l := len(*stack)
	if l < 1 {
		return stackFrame{}, fmt.Errorf("stack underflow")
	}

	v := (*stack)[l-1]
	*stack = (*stack)[:l-1]
	return v, nil
}
