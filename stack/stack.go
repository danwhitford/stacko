package stack

import (
	"fmt"
	"strings"
)

type Stack[T any] struct {
	head *StackNode[T]
}

type StackNode[T any] struct {
	val  T
	next *StackNode[T]
}

func (stack *Stack[T]) Push(val T) {
	node := StackNode[T]{val: val, next: stack.head}
	stack.head = &node
}

func (stack *Stack[T]) Pop() (T, error) {
	var ret T
	if stack.head == nil {
		return ret, fmt.Errorf("stack underflow")
	}
	ret = stack.head.val
	newHead := stack.head.next
	stack.head = newHead
	return ret, nil
}

func (stack Stack[T]) String() string {
	if stack.head == nil {
		return "<empty stack>"
	}
	var sb strings.Builder
	sb.WriteString("*--TOP--\n")
	cur := stack.head
	for cur != nil {
		sb.WriteString(cur.String())
		cur = cur.next
	}
	sb.WriteString("*-------")
	return sb.String()
}

func (node StackNode[T]) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("| %v\n", node.val))
	return sb.String()
}
