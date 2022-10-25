package main

import (
	"bufio"
	"log"
	"os"

	"github.com/danwhitford/stacko/parser"
	"github.com/danwhitford/stacko/tokeniser"
	"github.com/danwhitford/stacko/vm"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var tknsr tokeniser.Tokeniser
	var prser parser.Parser
	vm := vm.NewVM()

	for scanner.Scan() {
		t := scanner.Text()
		tknsr = tokeniser.NewTokeniser(t)
		tokens, err := tknsr.Tokenise()
		if err != nil {
			log.Println(err)
			continue
		}

		prser = parser.NewParser(tokens)
		vals, err := prser.Parse()
		if err != nil {
			log.Println(err)
			continue
		}

		vm.Load(vals)
		vm.Execute()
	}
}
