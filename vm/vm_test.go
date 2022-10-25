package vm

import (
	"testing"

	"github.com/danwhitford/stacko/stack"
)

func TestAdd(t *testing.T) {
	vm := NewVM()
	vm.Load([]stack.StackoVal{
		{StackoType: stack.StackoInt, Val: 1}, 
		{StackoType: stack.StackoInt, Val: 2},
		{StackoType: stack.StackoWord, Val: "+"}})
	vm.Execute()

	if vm.stack[0].Val != 3 {
		t.Fatal("can't add up")
	}
}
