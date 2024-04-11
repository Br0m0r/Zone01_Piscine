package piscine

import "fmt"

// Describe what the function does
func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for j := 0; j < y; j++ { // loop one to create a dimension for x
		for i := 0; i < x; i++ { // loop two to create a 2nd dimension for y
			switch {
			case j == 0 && i == 0: // Top left corner
				fmt.Print("o")
			case j == 0 && i == x-1: // Top right corner
				fmt.Print("o")
			case j == y-1 && i == 0: // Bottom left corner
				fmt.Print("o")
			case j == y-1 && i == x-1: // Bottom right corner
				fmt.Print("o")
			case j == 0 || j == y-1: // Top and bottom edges
				fmt.Print("-")
			case i == 0 || i == x-1: // Left and right edges
				fmt.Print("|")
			default: // if none of the above,print blank spaces
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}