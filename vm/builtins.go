package vm

//go:generate go run builtin_gen.go

import (
	"fmt"

	"github.com/danwhitford/stacko/stack"
	"github.com/danwhitford/stacko/stackoval"
)

func (vm *VM) PrintTop() error {
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	fmt.Fprintln(vm.outF, a)
	return nil
}

func (vm *VM) PrintStack() error {
	for i := len(vm.stack) - 1; i >= 0; i-- {
		fmt.Fprintln(vm.outF, vm.stack[i])
	}
	return nil
}

func (vm *VM) Dup() error {
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	vm.stack.Push(a)
	vm.stack.Push(a)
	return nil
}

func (vm *VM) execBuiltin(word string) (bool, error) {
	switch word {
	case "+":
		return true, vm.Add()
	case "-":
		return true, vm.Sub()
	case "*":
		return true, vm.Mult()
	case "/":
		return true, vm.Div()
	case "%":
		return true, vm.Mod()
	case ".":
		return true, vm.PrintTop()
	case "v":
		return true, vm.PrintStack()
	case "dup":
		return true, vm.Dup()
	case "swap":
		return true, vm.Swap()
	case "over":
		return true, vm.Over()
	case "rot":
		return true, vm.Rot()
	case "drop":
		return true, vm.Drop()
	case "def":
		return true, vm.Def()
	case "dbg":
		return true, vm.Dbg()
	case "=":
		return true, vm.Eq()
	case ">":
		return true, vm.Gt()
	case "<":
		return true, vm.Lt()
	case "if":
		return true, vm.If()
	case "clear":
		return true, vm.Clear()
	case "nuke":
		return true, vm.Nuke()
	case "range":
		return true, vm.Range()
	case "each":
		return true, vm.Each()
	case "times":
		return true, vm.Times()
	default:
		return false, nil
	}
}

func (vm *VM) Times() error {
	stack := &vm.stack
	a, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting false branch: %w", err)
	}
	if a.StackoType != stackoval.StackoInt {
		return fmt.Errorf("wanted int for range but got %+v", a)
	}
	lim := a.Val.(int)

	b, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting false branch: %w", err)
	}
	if b.StackoType != stackoval.StackoFn {
		return fmt.Errorf("wanted function for range but got %+v", a)
	}

	next := listise(b)
	frame := NewRegularFrame(next, lim-1)
	vm.instructions.Push(frame)
	return nil
}

func (vm *VM) Each() error {
	// stack := &vm.stack
	// a, err := stack.Pop()
	// if err != nil {
	// 	return fmt.Errorf("error getting false branch: %w", err)
	// }
	// if a.StackoType != stackoval.StackoList {
	// 	return fmt.Errorf("wanted list for range but got %+v", a)
	// }
	// lim := a.Val.(int)
	// vm.loopControlStack.Push(LoopControlFrame{lim})
	return nil
}

func (vm *VM) Range() error {
	stack := &vm.stack
	a, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting false branch: %w", err)
	}
	if a.StackoType != stackoval.StackoInt {
		return fmt.Errorf("wanted int for range but got %+v", a)
	}
	lim := a.Val.(int)
	ret := make([]stackoval.StackoVal, lim)
	for i := range ret {
		ret[i] = stackoval.StackoVal{StackoType: stackoval.StackoInt, Val: i}
	}
	val := stackoval.StackoVal{StackoType: stackoval.StackoList, Val: ret}
	stack.Push(val)
	return nil
}

func (vm *VM) Nuke() error {
	vm.Reset()
	return nil
}

func (vm *VM) Clear() error {
	vm.stack = make(stack.Stack[stackoval.StackoVal], 0)
	return nil
}

func (vm *VM) If() error {
	stack := &vm.stack

	falseBranch, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting false branch: %w", err)
	}
	trueBranch, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting true branch: %w", err)
	}
	condition, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting condition: %w", err)
	}
	var branch stackoval.StackoVal
	if condition.Val == true {
		branch = trueBranch
	} else {
		branch = falseBranch
	}

	fmt.Printf("%+v\n%+v\n%+v\n", falseBranch, trueBranch, branch)
	next := listise(branch)
	vm.Load(next)
	return nil
}

func listise(val stackoval.StackoVal) []stackoval.StackoVal {
	switch val.StackoType {
	case stackoval.StackoFn:
		casted := val.Val.([]stackoval.StackoVal)
		return casted
	case stackoval.StackoSymbol:
		return []stackoval.StackoVal{{StackoType: stackoval.StackoWord, Val: val.Val}}
	default:
		return []stackoval.StackoVal{val}
	}
}

func (vm *VM) Dbg() error {
	fmt.Fprintf(vm.outF, "%+v\n", *vm)
	return nil
}

func (vm *VM) Swap() error {
	stack := &vm.stack
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

func (vm *VM) Over() error {
	stack := &vm.stack
	a, err := stack.Pop()
	if err != nil {
		return err
	}
	b, err := stack.Pop()
	if err != nil {
		return err
	}
	stack.Push(b)
	stack.Push(a)
	stack.Push(b)
	return nil
}

func (vm *VM) Rot() error {
	stack := &vm.stack
	a, err := stack.Pop()
	if err != nil {
		return err
	}
	b, err := stack.Pop()
	if err != nil {
		return err
	}
	c, err := stack.Pop()
	if err != nil {
		return err
	}
	stack.Push(b)
	stack.Push(a)
	stack.Push(c)
	return nil
}

func (vm *VM) Drop() error {
	stack := &vm.stack
	_, err := stack.Pop()
	return err
}

func (vm *VM) Def() error {
	word, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	val, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	vm.dictionary[word.Val.(string)] = val
	return nil
}
