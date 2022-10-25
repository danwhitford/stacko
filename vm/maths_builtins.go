
package vm

import 	"github.com/danwhitford/stacko/stack"


func Add(stk *stack.Stack)   error {
	a, err := stk.Pop()
	if err != nil {
		return err
	}
	b, err := stk.Pop()
	if err != nil {
		return err
	}

	switch {
	case a.StackoType == stack.StackoInt && b.StackoType == stack.StackoInt:
		stk.Push(stack.StackoVal{StackoType: stack.StackoInt, Val: b.Val.(int) + a.Val.(int)})
	case a.StackoType == stack.StackoInt && b.StackoType == stack.StackoFloat:
		aa := float64(a.Val.(int))
		bb := b.Val.(float64)
		stk.Push(stack.StackoVal{StackoType: stack.StackoFloat, Val: bb + aa})
	case a.StackoType == stack.StackoFloat && b.StackoType == stack.StackoInt:
		aa := a.Val.(float64)
		bb := float64(b.Val.(int))
		stk.Push(stack.StackoVal{StackoType: stack.StackoFloat, Val: bb + aa})
	case a.StackoType == stack.StackoFloat && b.StackoType == stack.StackoFloat:
		aa := a.Val.(float64)
		bb := b.Val.(float64)
		stk.Push(stack.StackoVal{StackoType: stack.StackoFloat, Val: bb + aa})
	}

	return nil
}

func Sub(stk *stack.Stack)   error {
	a, err := stk.Pop()
	if err != nil {
		return err
	}
	b, err := stk.Pop()
	if err != nil {
		return err
	}

	switch {
	case a.StackoType == stack.StackoInt && b.StackoType == stack.StackoInt:
		stk.Push(stack.StackoVal{StackoType: stack.StackoInt, Val: b.Val.(int) - a.Val.(int)})
	case a.StackoType == stack.StackoInt && b.StackoType == stack.StackoFloat:
		aa := float64(a.Val.(int))
		bb := b.Val.(float64)
		stk.Push(stack.StackoVal{StackoType: stack.StackoFloat, Val: bb - aa})
	case a.StackoType == stack.StackoFloat && b.StackoType == stack.StackoInt:
		aa := a.Val.(float64)
		bb := float64(b.Val.(int))
		stk.Push(stack.StackoVal{StackoType: stack.StackoFloat, Val: bb - aa})
	case a.StackoType == stack.StackoFloat && b.StackoType == stack.StackoFloat:
		aa := a.Val.(float64)
		bb := b.Val.(float64)
		stk.Push(stack.StackoVal{StackoType: stack.StackoFloat, Val: bb - aa})
	}

	return nil
}

func Mult(stk *stack.Stack)   error {
	a, err := stk.Pop()
	if err != nil {
		return err
	}
	b, err := stk.Pop()
	if err != nil {
		return err
	}

	switch {
	case a.StackoType == stack.StackoInt && b.StackoType == stack.StackoInt:
		stk.Push(stack.StackoVal{StackoType: stack.StackoInt, Val: b.Val.(int) * a.Val.(int)})
	case a.StackoType == stack.StackoInt && b.StackoType == stack.StackoFloat:
		aa := float64(a.Val.(int))
		bb := b.Val.(float64)
		stk.Push(stack.StackoVal{StackoType: stack.StackoFloat, Val: bb * aa})
	case a.StackoType == stack.StackoFloat && b.StackoType == stack.StackoInt:
		aa := a.Val.(float64)
		bb := float64(b.Val.(int))
		stk.Push(stack.StackoVal{StackoType: stack.StackoFloat, Val: bb * aa})
	case a.StackoType == stack.StackoFloat && b.StackoType == stack.StackoFloat:
		aa := a.Val.(float64)
		bb := b.Val.(float64)
		stk.Push(stack.StackoVal{StackoType: stack.StackoFloat, Val: bb * aa})
	}

	return nil
}

func Div(stk *stack.Stack)   error {
	a, err := stk.Pop()
	if err != nil {
		return err
	}
	b, err := stk.Pop()
	if err != nil {
		return err
	}

	switch {
	case a.StackoType == stack.StackoInt && b.StackoType == stack.StackoInt:
		stk.Push(stack.StackoVal{StackoType: stack.StackoInt, Val: b.Val.(int) / a.Val.(int)})
	case a.StackoType == stack.StackoInt && b.StackoType == stack.StackoFloat:
		aa := float64(a.Val.(int))
		bb := b.Val.(float64)
		stk.Push(stack.StackoVal{StackoType: stack.StackoFloat, Val: bb / aa})
	case a.StackoType == stack.StackoFloat && b.StackoType == stack.StackoInt:
		aa := a.Val.(float64)
		bb := float64(b.Val.(int))
		stk.Push(stack.StackoVal{StackoType: stack.StackoFloat, Val: bb / aa})
	case a.StackoType == stack.StackoFloat && b.StackoType == stack.StackoFloat:
		aa := a.Val.(float64)
		bb := b.Val.(float64)
		stk.Push(stack.StackoVal{StackoType: stack.StackoFloat, Val: bb / aa})
	}

	return nil
}

