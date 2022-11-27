//go:build ignore

package main

import (
	"io"
	"log"
	"os"
	"text/template"
)

func mathsBuiltins() {
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
	}

	runTemplate(operators, "templates/maths_builtins.go.tmpl", "maths_builtins.go")
}

func booleanBuiltins() {
	var operators = []struct {
		Name string
		Op   string
	}{
		{
			Name: "Eq",
			Op:   "==",
		},
		{
			Name: "Gt",
			Op:   ">",
		},
		{
			Name: "Lt",
			Op:   "<",
		},
	}

	runTemplate(operators, "templates/boolean_builtins.go.tmpl", "boolean_builtins.go")
}

func runTemplate(data interface{}, tmplName, outName string) {
	tmplf, err := os.Open(tmplName)
	if err != nil {
		panic(err)
	}
	tmpl, err := io.ReadAll(tmplf)
	if err != nil {
		panic(err)
	}
	t, err := template.New("letter").Parse(string(tmpl))
	if err != nil {
		log.Fatal("error executing template:", err)
	}

	f, err := os.Create(outName)
	if err != nil {
		log.Fatal("error executing template:", err)
	}
	defer f.Close()
	err = t.Execute(f, data)
	if err != nil {
		log.Fatal("error executing template:", err)
	}
}

func main() {
	mathsBuiltins()
	booleanBuiltins()
}
