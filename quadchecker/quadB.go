package main

import (
	"fmt"
	"os"
	"strconv"
)

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			switch {
			case j == 0 && i == 0: // Top left corner
				fmt.Print("/")
			case j == 0 && i == x-1: // Top right corner
				fmt.Print("\\")
			case j == y-1 && i == 0: // Bottom left corner
				fmt.Print("\\")
			case j == y-1 && i == x-1: // Bottom right corner
				fmt.Print("/")
			case j == 0 || j == y-1: // Top and bottom edges
				fmt.Print("*")
			case i == 0 || i == x-1: // Left and right edges
				fmt.Print("*")
			default: // if none of the above, print blank spaces
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <width> <height>")
		return
	}
	width, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid width:", os.Args[1])
		return
	}
	height, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid height:", os.Args[2])
		return
	}
	QuadB(width, height)
}
