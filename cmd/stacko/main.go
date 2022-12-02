package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/danwhitford/stacko/parser"
	"github.com/danwhitford/stacko/tokeniser"
	"github.com/danwhitford/stacko/vm"
)

type Runner struct {
	tknsr tokeniser.Tokeniser
	prser parser.Parser
	vm    vm.VM
}

func (runner *Runner) doLine(s string) error {
	runner.tknsr = tokeniser.NewTokeniser(s)
	tokens, err := runner.tknsr.Tokenise()
	if err != nil {
		return fmt.Errorf("error while tokenising: %w", err)
	}
	if len(tokens) < 1 {
		return nil
	}

	runner.prser = parser.NewParser(tokens)
	vals, err := runner.prser.Parse()
	if err != nil {
		return fmt.Errorf("error while parsing: %w", err)
	}

	runner.vm.Load(vals)
	err = runner.vm.Execute()
	if err != nil {
		return fmt.Errorf("error while executing: %w", err)
	}

	return nil
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	machine, err := vm.NewVM(os.Stdout)
	if err != nil {
		panic(err)
	}
	runner := Runner{
		vm: machine,
	}

	for scanner.Scan() {
		t := scanner.Text()
		err := runner.doLine(t)
		if err != nil {
			log.Println(err)
		}
	}
}

func runFile(fname string, outf io.Writer) {
	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	machine, err := vm.NewVM(os.Stdout)
	if err != nil {
		panic(err)
	}
	runner := Runner{
		vm: machine,
	}
	err = runner.doLine(string(b))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		repl()
	} else {
		fname := os.Args[1]
		runFile(fname, os.Stdout)
	}
}
