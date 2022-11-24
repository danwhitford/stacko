package main

import (
	"bufio"
	"fmt"
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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	runner := Runner{
		vm: vm.NewVM(os.Stdout),
	}

	for scanner.Scan() {
		t := scanner.Text()
		err := runner.doLine(t)
		if err != nil {
			log.Println(err)
		}
	}
}
