package piscine

import "github.com/01-edu/z01"

// prints one by one the characters of a string on the screen.
func PrintStr(s string) {
	astringChanged := []byte(s)
	for i := 0; i < len(s); i++ {
		A := astringChanged[i]
		B := rune(A)
		z01.PrintRune(B)

	}
}
