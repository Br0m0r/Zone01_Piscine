package piscine

import "github.com/01-edu/z01"

// Prints True if int is Negative/Prints False if positive
func IsNegative(nb int) {
	if nb < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}
