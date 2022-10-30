package vm

import (
	"github.com/danwhitford/stacko/stackoval"
)

func (vm *VM) Add() error {
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	b, err := vm.stack.Pop()
	if err != nil {
		return err
	}

	switch {
	case a.StackoType == stackoval.StackoInt && b.StackoType == stackoval.StackoInt:
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoInt, Val: b.Val.(int) + a.Val.(int)})

	case a.StackoType == stackoval.StackoInt && b.StackoType == stackoval.StackoFloat:
		aa := float64(a.Val.(int))
		bb := b.Val.(float64)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb + aa})
	case a.StackoType == stackoval.StackoFloat && b.StackoType == stackoval.StackoInt:
		aa := a.Val.(float64)
		bb := float64(b.Val.(int))
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb + aa})
	case a.StackoType == stackoval.StackoFloat && b.StackoType == stackoval.StackoFloat:
		aa := a.Val.(float64)
		bb := b.Val.(float64)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb + aa})

	}
	return nil
}

func (vm *VM) Sub() error {
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	b, err := vm.stack.Pop()
	if err != nil {
		return err
	}

	switch {
	case a.StackoType == stackoval.StackoInt && b.StackoType == stackoval.StackoInt:
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoInt, Val: b.Val.(int) - a.Val.(int)})

	case a.StackoType == stackoval.StackoInt && b.StackoType == stackoval.StackoFloat:
		aa := float64(a.Val.(int))
		bb := b.Val.(float64)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb - aa})
	case a.StackoType == stackoval.StackoFloat && b.StackoType == stackoval.StackoInt:
		aa := a.Val.(float64)
		bb := float64(b.Val.(int))
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb - aa})
	case a.StackoType == stackoval.StackoFloat && b.StackoType == stackoval.StackoFloat:
		aa := a.Val.(float64)
		bb := b.Val.(float64)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb - aa})

	}
	return nil
}

func (vm *VM) Mult() error {
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	b, err := vm.stack.Pop()
	if err != nil {
		return err
	}

	switch {
	case a.StackoType == stackoval.StackoInt && b.StackoType == stackoval.StackoInt:
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoInt, Val: b.Val.(int) * a.Val.(int)})

	case a.StackoType == stackoval.StackoInt && b.StackoType == stackoval.StackoFloat:
		aa := float64(a.Val.(int))
		bb := b.Val.(float64)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb * aa})
	case a.StackoType == stackoval.StackoFloat && b.StackoType == stackoval.StackoInt:
		aa := a.Val.(float64)
		bb := float64(b.Val.(int))
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb * aa})
	case a.StackoType == stackoval.StackoFloat && b.StackoType == stackoval.StackoFloat:
		aa := a.Val.(float64)
		bb := b.Val.(float64)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb * aa})

	}
	return nil
}

func (vm *VM) Div() error {
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	b, err := vm.stack.Pop()
	if err != nil {
		return err
	}

	switch {
	case a.StackoType == stackoval.StackoInt && b.StackoType == stackoval.StackoInt:
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoInt, Val: b.Val.(int) / a.Val.(int)})

	case a.StackoType == stackoval.StackoInt && b.StackoType == stackoval.StackoFloat:
		aa := float64(a.Val.(int))
		bb := b.Val.(float64)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb / aa})
	case a.StackoType == stackoval.StackoFloat && b.StackoType == stackoval.StackoInt:
		aa := a.Val.(float64)
		bb := float64(b.Val.(int))
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb / aa})
	case a.StackoType == stackoval.StackoFloat && b.StackoType == stackoval.StackoFloat:
		aa := a.Val.(float64)
		bb := b.Val.(float64)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb / aa})

	}
	return nil
}

func (vm *VM) Mod() error {
	a, err := vm.stack.Pop()
	if err != nil {
		return err
	}
	b, err := vm.stack.Pop()
	if err != nil {
		return err
	}

	switch {
	case a.StackoType == stackoval.StackoInt && b.StackoType == stackoval.StackoInt:
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoInt, Val: b.Val.(int) % a.Val.(int)})

	}
	return nil
}
