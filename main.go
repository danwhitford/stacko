package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/danwhitford/stacko/stack"
)

func main() {
	fmt.Println("==== stacko ===")

	var stack stack.Stack
	reader := bufio.NewScanner(os.Stdin)
	var line string
	for {
		fmt.Print(">>> ")
		stuff := reader.Scan()
		if !stuff {
			break
		}
		line = reader.Text()
		if len(line) < 1 {
			continue
		}

		words := strings.Split(line, " ")
		var err error
		for _, word := range words {
			switch word {
			case "+":
				{
					var a, b int
					a, err = stack.Pop()
					if err != nil {
						break
					}
					b, err = stack.Pop()
					if err != nil {
						break
					}
					stack.Push(a + b)
				}
			case "-":
				{
					var a, b int
					a, err = stack.Pop()
					if err != nil {
						break
					}
					b, err = stack.Pop()
					if err != nil {
						break
					}
					stack.Push(a - b)
				}
			case "*":
				{
					var a, b int
					a, err = stack.Pop()
					if err != nil {
						break
					}
					b, err = stack.Pop()
					if err != nil {
						break
					}
					stack.Push(a * b)
				}
			case "/":
				{
					var a, b int
					a, err = stack.Pop()
					if err != nil {
						break
					}
					b, err = stack.Pop()
					if err != nil {
						break
					}
					stack.Push(a / b)
				}
			case "dup":
				{
					var a int
					a, err = stack.Pop()
					if err != nil {
						break
					}
					stack.Push(a)
					stack.Push(a)
				}
			case "drop":
				{
					stack.Pop()
				}
			case "swap":
				{
					var a, b int
					a, err = stack.Pop()
					if err != nil {
						break
					}
					b, err = stack.Pop()
					if err != nil {
						break
					}
					stack.Push(a)
					stack.Push(b)
				}
			case "over":
				{
					var a, b int
					a, err = stack.Pop()
					if err != nil {
						break
					}
					b, err = stack.Pop()
					if err != nil {
						break
					}
					stack.Push(b)
					stack.Push(a)
					stack.Push(b)
				}
			case "rot":
				{
					var a, b, c int
					a, err = stack.Pop()
					if err != nil {
						break
					}
					b, err = stack.Pop()
					if err != nil {
						break
					}
					c, err = stack.Pop()
					if err != nil {
						break
					}
					stack.Push(b)
					stack.Push(a)
					stack.Push(c)
				}
			case ".":
				{
					var a int
					a, err = stack.Pop()
					if err != nil {
						break
					}
					fmt.Printf("%v\n", a)
				}
			case "emit":
				{
					var a int
					a, err = stack.Pop()
					if err != nil {
						break
					}
					fmt.Printf("%c", a)
				}
			case "cr":
				{
					fmt.Println()
				}
			case "stack":
				{
					fmt.Println(stack)
				}
			default:
				{
					var val int
					val, err = strconv.Atoi(word)
					if err != nil {
						err = fmt.Errorf("could not parse '%v', not a known word or integer", word)
						break
					}
					stack.Push(val)
				}
			}
			if err != nil {
				break
			}
		}
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("ok")
		}
	}
}
