package vm

//go:generate go run builtin_gen.go

import (
	"fmt"

	"github.com/danwhitford/stacko/stackoval"
)

func (vm *VM) PrintTop() error {
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	fmt.Println(a)
	return nil
}

func (vm *VM) PrintStack() error {
	for i := len(vm.stack) - 1; i >= 0; i-- {
		fmt.Println(vm.stack[i])
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
	case "if":
		return true, vm.If()
	default:
		return false, nil
	}
}

func listise(val stackoval.StackoVal) []stackoval.StackoVal {	
	switch val.StackoType {
	case stackoval.StackoList:
		casted := val.Val.([]stackoval.StackoVal)		
		return casted
	case stackoval.StackoSymbol:
		return []stackoval.StackoVal{{StackoType: stackoval.StackoWord, Val: val.Val}}
	default:
		return []stackoval.StackoVal{val}
	}
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

func (vm *VM) Dbg() error {
	fmt.Printf("%+v\n", *vm)
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
