package piscine

import (
	"github.com/01-edu/z01"
)

// prints an int passed in parameter.All possible values of type int go through
func PrintNbr(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
	}
	printIndirect(n)
}

func printIndirect(n int) {
	if n == 0 {
		return
	}
	digit := n % 10
	if digit < 0 {
		digit = -digit
	}
	printIndirect(n / 10)
	z01.PrintRune(rune(digit) + '0')
}
