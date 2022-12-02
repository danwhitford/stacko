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
	case "clear":
		return true, vm.Clear()
	case "nuke":
		return true, vm.Nuke()
	case "call":
		return true, vm.Call()
	case "branch":
		return true, vm.Branch()
	case "rpush":
		return true, vm.ReturnPush()
	case "rpop":
		return true, vm.ReturnPop()
	case "append":
		return true, vm.Append()
	default:
		return false, nil
	}
}

func (vm *VM) Append() error {
	stack := &vm.stack
	a, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting func to call: %w", err)
	}
	b, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting func to call: %w", err)
	}
	if b.StackoType != stackoval.StackoList {
		return fmt.Errorf("need list for concat but got %+v", b.StackoType)
	}
	els := b.Val.([]stackoval.StackoVal)
	els = append(els, a)
	b.Val = els
	stack.Push(b)
	return nil
}

func (vm *VM) ReturnPush() error {
	stack := &vm.stack
	returnStack := &vm.returnStack
	a, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting func to call: %w", err)
	}
	returnStack.Push(a)
	return err
}

func (vm *VM) ReturnPop() error {
	stack := &vm.stack
	returnStack := &vm.returnStack
	a, err := returnStack.Pop()
	if err != nil {
		return fmt.Errorf("error getting func to call: %w", err)
	}
	stack.Push(a)
	return err
}

func (vm *VM) Branch() error {
	stack := &vm.stack
	a, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting func to call: %w", err)
	}
	if a.StackoType != stackoval.StackoBool {
		return fmt.Errorf("need boolean for branch but got %+v", a)
	}
	cond := a.Val.(bool)
	if cond {
		return vm.Drop()
	} else {
		err := vm.Swap()
		if err != nil {
			return err
		}
		return vm.Drop()
	}
}

func (vm *VM) Call() error {
	stack := &vm.stack
	a, err := stack.Pop()
	if err != nil {
		return fmt.Errorf("error getting func to call: %w", err)
	}
	if a.StackoType != stackoval.StackoFn {
		return fmt.Errorf("wanted fn for call but got %v", a)
	}
	next := listise(a)
	frame := NewRegularFrame(next)
	vm.callStack.Push(frame)
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

	next := listise(branch)
	vm.Load(next)
	return nil
}

func listise(val stackoval.StackoVal) []stackoval.StackoVal {
	switch val.StackoType {
	case stackoval.StackoFn, stackoval.StackoList:
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
