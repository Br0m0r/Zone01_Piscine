package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i >= 0; i-- {
		// Print each rune  of the argument
		for _, r := range os.Args[i] {
			if r != '/' {
				z01.PrintRune(r)
			} else {
				break
			}
		}
		// print a newline character
		if i != 0 {
			z01.PrintRune('\n')
		}
	}
}
