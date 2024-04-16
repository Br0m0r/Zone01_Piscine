package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lowercaseAlphabet := "abcdefghijklmnopqrstuvwxyz"
	uppercaseAlphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	isUpper := false
	startIndex := 1
	if len(os.Args) > 1 && os.Args[1] == "--upper" {
		isUpper = true
		startIndex++
	}

	// Track if any valid input (resulting in output) has been processed.
	anyValidInput := false

	for _, arg := range os.Args[startIndex:] {
		if num := Atoi(arg); num >= 1 && num <= 26 {
			anyValidInput = true // Mark that we've processed a valid input.
			if isUpper {
				z01.PrintRune(rune(uppercaseAlphabet[num-1]))
			} else {
				z01.PrintRune(rune(lowercaseAlphabet[num-1]))
			}
		} else {
			// Only print a space (and mark valid input) if it's explicitly an invalid number, not for missing arguments.
			if arg != "" {
				z01.PrintRune(' ')
				anyValidInput = true
			}
		}
	}

	// Print a newline only if there was any valid input.
	if anyValidInput {
		z01.PrintRune('\n')
	}
}

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	result := 0
	sign := 1
	startIndex := 0

	if s[0] == '-' || s[0] == '+' {
		if s[0] == '-' {
			sign = -1
		}
		startIndex = 1
	}

	for i := startIndex; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}
	return result * sign
}
