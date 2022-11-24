//go:build ignore

package main

import (
	"log"
	"os"
	"text/template"
)

const tmpl string = `
package vm

import 	"github.com/danwhitford/stacko/stackoval"

{{ range . }}
func (vm *VM) {{ .Name }}() error {
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
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoInt, Val: b.Val.(int) {{ .Op }} a.Val.(int)})
{{ if not .IntOnly }}
	case a.StackoType == stackoval.StackoInt && b.StackoType == stackoval.StackoFloat:
		aa := float64(a.Val.(int))
		bb := b.Val.(float64)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb {{ .Op }} aa})
	case a.StackoType == stackoval.StackoFloat && b.StackoType == stackoval.StackoInt:
		aa := a.Val.(float64)
		bb := float64(b.Val.(int))
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb {{ .Op }} aa})
	case a.StackoType == stackoval.StackoFloat && b.StackoType == stackoval.StackoFloat:
		aa := a.Val.(float64)
		bb := b.Val.(float64)
		vm.stack.Push(stackoval.StackoVal{StackoType: stackoval.StackoFloat, Val: bb {{ .Op }} aa})
{{ end }}
	}
	return nil
}
{{end}}
`

func main() {
	var operators = []struct {
		Name    string
		Op      string
		IntOnly bool
	}{
		{
			Name: "Add",
			Op:   "+",
		},
		{
			Name: "Sub",
			Op:   "-",
		},
		{
			Name: "Mult",
			Op:   "*",
		},
		{
			Name: "Div",
			Op:   "/",
		},
		{
			Name:    "Mod",
			Op:      "%",
			IntOnly: true,
		},
		{
			Name: "Eq",
			Op:   "==",
		},
	}

	t, err := template.New("letter").Parse(tmpl)
	if err != nil {
		log.Fatal("error executing template:", err)
	}

	f, err := os.Create("maths_builtins.go")
	if err != nil {
		log.Fatal("error executing template:", err)
	}
	defer f.Close()
	err = t.Execute(f, operators)
	if err != nil {
		log.Fatal("error executing template:", err)
	}
}
