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
		reader.Scan()
		line = reader.Text()
		if len(line) < 1 {
			continue
		}

		words := strings.Split(line, " ")
		for _, word := range words {
			switch word {
			case "+":
				{
					a := stack.Pop()
					b := stack.Pop()
					stack.Push(a + b)
				}
			case "-":
				{
					a := stack.Pop()
					b := stack.Pop()
					stack.Push(a - b)
				}
			case "*":
				{
					a := stack.Pop()
					b := stack.Pop()
					stack.Push(a * b)
				}
			case "/":
				{
					a := stack.Pop()
					b := stack.Pop()
					stack.Push(a / b)
				}
			case "dup":
				{
					a := stack.Pop()
					stack.Push(a)
					stack.Push(a)
				}
			case "drop":
				{
					stack.Pop()
				}
			case "swap":
				{
					a := stack.Pop()
					b := stack.Pop()
					stack.Push(a)
					stack.Push(b)
				}
			case "over":
				{
					a := stack.Pop()
					b := stack.Pop()
					stack.Push(b)
					stack.Push(a)
					stack.Push(b)
				}
			case "rot":
				{
					a := stack.Pop()
					b := stack.Pop()
					c := stack.Pop()
					stack.Push(b)
					stack.Push(a)
					stack.Push(c)
				}
			case ".":
				{
					fmt.Printf("%v\n", stack.Pop())
				}
			case "emit":
				{
					fmt.Printf("%c", stack.Pop())
				}
			case "cr": {
				fmt.Println()
			}
			default:
				{
					val, err := strconv.Atoi(word)
					if err != nil {
						fmt.Printf("failed to convert '%v' to int\n", err)
						continue
					}
					stack.Push(val)
				}
			}
		}
		fmt.Println(stack)
	}
}
