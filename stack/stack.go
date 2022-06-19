package stack

import (
	"fmt"
	"strings"
)

type Stack struct {
	head *StackNode
}

type StackNode struct {
	val  int
	next *StackNode
}

func (stack *Stack) Push(val int) {
	node := StackNode{val: val, next: stack.head}
	stack.head = &node
}

func (stack *Stack) Pop() (int, error) {
	if stack.head == nil {
		return 0, fmt.Errorf("stack underflow")
	}
	ret := stack.head.val
	newHead := stack.head.next
	stack.head = newHead
	return ret, nil
}

func (stack Stack) String() string {
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

func (node StackNode) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("| %v\n", node.val))
	return sb.String()
}
