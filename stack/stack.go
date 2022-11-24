package stack

import (
	"fmt"
)

type Stack[T any] []T

func (stack *(Stack[T])) Push(v T) {
	*stack = append(*stack, v)
}

func (stack *Stack[T]) Pop() (T, error) {
	l := len(*stack)
	if l < 1 {
		var t T
		return t, fmt.Errorf("stack underflow")
	}

	v := (*stack)[l-1]
	*stack = (*stack)[:l-1]
	return v, nil
}

func (stack *Stack[T]) Peek() (*T, error) {
	l := len(*stack)
	if l < 1 {
		var t T
		return &t, fmt.Errorf("stack underflow")
	}
	v := &(*stack)[l-1]
	return v, nil
}

func (stack *Stack[T]) Empty() bool {
	return len(*stack) == 0
}

func (stack *Stack[T]) Swap() error {
	a, err := stack.Pop()
	if err != nil {
		return err
	}
	b, err := stack.Pop()
	if err != nil {
		return err
	}
	stack.Push(a)
	stack.Push(b)
	return nil
}
