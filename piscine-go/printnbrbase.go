package piscine

import (
	"github.com/01-edu/z01"
)

// actual func, if base is invalid, prints "NV".
func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		printStr("NV")
		return
	}

	if nbr < 0 {
		z01.PrintRune('-')
		convertAndPrint(-nbr, base)
	} else {
		convertAndPrint(nbr, base)
	}
}

// checks if base is valid according to the given rules.
func isValidBase(base string) bool {
	if len(base) < 2 || contains(base, '+') || contains(base, '-') {
		return false
	}

	for i := range base {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}

	return true
}

// checks if string s contains the character char.
func contains(s string, char rune) bool {
	for _, c := range s {
		if c == char {
			return true
		}
	}
	return false
}

// convertAndPrint converts the number to the base and prints it, using recursion.
func convertAndPrint(n int, base string) {
	if n >= len(base) {
		convertAndPrint(n/len(base), base)
	}
	z01.PrintRune(rune(base[n%len(base)]))
}

// printStr prints a string, character by character.
func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
}
