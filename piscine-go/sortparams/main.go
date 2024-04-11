package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:] // Get all arguments except the program name

	// Sort the arguments array using Bubble Sort
	for i := 0; i < len(args)-1; i++ {
		for j := 0; j < len(args)-i-1; j++ {
			if args[j] > args[j+1] {
				// Swap the arguments
				args[j], args[j+1] = args[j+1], args[j]
			}
		}
	}

	// Print each argument in the sorted order
	for _, arg := range args {
		// Print the argument
		for _, r := range arg {
			z01.PrintRune(r)
		}
		// Print a newline character except after the last argument

		z01.PrintRune('\n')

	}
}
