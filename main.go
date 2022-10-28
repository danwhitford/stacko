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
			log.Printf("error while tokenising: %s\n", err)
			continue
		}

		prser = parser.NewParser(tokens)
		vals, err := prser.Parse()
		if err != nil {
			log.Printf("error while parsing: %s\n", err)
			continue
		}

		vm.Load(vals)
		err = vm.Execute()
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
