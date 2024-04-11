package piscine

import (
	"github.com/01-edu/z01"
)

// function which prints the digits of an int passed in parameter in ascending order. All possible values of type int have to go through, excluding negative numbers.
func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(('0'))
	}

	digitcount := [10]int{} // use the indexes to represent 0-9
	for n > 0 {
		digit := n % 10     // gets the last digit
		digitcount[digit]++ // increment the position of index
		n /= 10             // remove the digit and move on to the next
	}

	for i, count := range digitcount { // loop through the slice
		for count > 0 {
			z01.PrintRune(rune(i) + '0') // print the index,essentialy in how many times the number was found,was 'saved' in the slice index
			count--                      // decrement to go to the next one

		}
	}
}
