package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args[0]

	lastSlashIndex := -1
	for i, r := range argument {
		if r == '/' {
			lastSlashIndex = i
		}
	}

	executableName := argument
	if lastSlashIndex != -1 {
		executableName = argument[lastSlashIndex+1:]
	}

	for _, r := range executableName {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
