package main

import (
	"fmt"
)

// ftPutChar is a function that takes a rune (a Unicode code point) and prints it as a string.
// This function is used to print individual characters to the console.
func ftPutChar(c rune) {
	fmt.Print(string(c)) // Converts the rune to a string and prints it.
}

// ftPrint is a function that prints a pattern of characters. It takes four runes:
// - len: the length of the pattern to print.
// - beginChar: the character to print at the beginning of the pattern.
// - midChar: the character to print in the middle of the pattern.
// - endChar: the character to print at the end of the pattern.
// This function uses a loop to print the pattern, with special handling for the first and last characters.
func ftPrint(len int, beginChar, midChar, endChar rune) {
	for i := 1; i <= len; i++ { // Looping from 1 to len to print the pattern.
		switch i { // Using a switch statement to handle the first, last, and middle characters differently.
		case 1: // For the first character, print the beginChar.
			ftPutChar(beginChar)
		case len: // For the last character, print the endChar.
			ftPutChar(endChar)
		default: // For all other characters, print the midChar.
			ftPutChar(midChar)
		}
	}
	ftPutChar('\n') // After printing the pattern, print a newline character to end the line.
}