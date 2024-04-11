/*package main

import "piscine"
       "github.com/01-edu/z01"

func main() {
	piscine.PrintComb2()
}



package piscine



func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for l := '0'; l <= '9'; l++ {
					if i != '9' || j != '9' || k != '9' || l != '9' {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					} else {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(l)

					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
EWW i should have 2 for loops and each var go from 0 to 99
*/

package piscine

import "github.com/01-edu/z01"

// prints in ascending order and on a single line: all possible combinations of two different two-digit numbers.These combinations are separated by a comma and a space.
func PrintComb2() {
	for i := 0; i <= 99; i++ {
		for j := i + 1; j <= 99; j++ {
			z01.PrintRune(rune(i/10 + '0')) // gets the first one by dividing by 10
			z01.PrintRune(rune(i%10 + '0')) // get the second one doing modulo essentialy getting the ipoloipo
			z01.PrintRune(' ')
			z01.PrintRune(rune(j/10 + '0'))
			z01.PrintRune(rune(j%10 + '0'))
			if i != 98 || j != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
