package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 9; i >= 0; i-- {
		for k := 9; k >= 0; k-- {
			for l := 9; l >= 0; l-- {
				for m := 9; m >= 0; m-- {
					if i > l || (i == l && k > m) {
						z01.PrintRune(rune(i + '0'))
						z01.PrintRune(rune(k + '0'))
						z01.PrintRune(' ')
						z01.PrintRune(rune(l + '0'))
						z01.PrintRune(rune(m + '0'))
						if !(i == 0 && k == 1 && l == 0 && m == 0) {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}
