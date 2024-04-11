package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, arg := range os.Args[1:] {
		// Print each rune (character) of the argument
		for _, r := range arg {
			z01.PrintRune(r)
		}
		// After printing each argument, print a newline character
		z01.PrintRune('\n')
	}
}
