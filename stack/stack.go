package stack

import (
	"fmt"
	"strings"
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

func (val StackoVal) String() string {
	if val.StackoType == StackoList {
		var sb strings.Builder
		sb.WriteRune('[')
		for i, v := range val.Val.([]StackoVal) {
			sb.WriteString(fmt.Sprint(v.Val))
			if i < len(val.Val.([]StackoVal))-1 {
				sb.WriteRune(' ')
			}
		}
		sb.WriteRune(']')
		return sb.String()
	} else {
		return fmt.Sprintf("%v", val.Val)
	}
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
