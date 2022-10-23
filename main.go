package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/danwhitford/stacko/tokeniser"
	"github.com/danwhitford/stacko/parser"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var tknsr tokeniser.Tokeniser
	var prser parser.Parser

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

		fmt.Printf("%+v\n", vals)
	}
}
