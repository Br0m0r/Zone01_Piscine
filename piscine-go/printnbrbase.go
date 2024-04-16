package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 || contains(base, '+') || contains(base, '-') || !isUnique(base) {
		printStr("NV") // Base is invalid
		return
	}
	if nbr < 0 {
		z01.PrintRune('-') // Handle negative numbers
		nbr = -nbr
	}
	convertAndPrint(nbr, base) // Convert and print the number in the given base
}

func contains(s string, char rune) bool {
	for _, c := range s {
		if c == char {
			return true // Character found
		}
	}
	return false // Character not found
}

func isUnique(s string) bool {
	for i, c := range s {
		if contains(s[i+1:], c) {
			return false // Duplicate character found
		}
	}
	return true
}

func convertAndPrint(n int, base string) {
	baseLen := len(base)
	if n < 0 {
		if n == -9223372036854775808 {
			convertAndPrint(n/-baseLen, base) // Adjust n here to avoid the negative index
			z01.PrintRune(rune(base[(-n-1)%baseLen+1]))
			return
		}
		n = -n // Make n positive if it's negative and not the edge case
	}
	if n >= baseLen {
		convertAndPrint(n/baseLen, base) // Recursive call for each digit
	}
	z01.PrintRune(rune(base[n%baseLen])) // Print the character for the current digit
}

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c) // Print each character
	}
}
