package main

import (
	"fmt"
	"os"
	"strconv"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for j := 0; j < y; j++ {
		for i := 0; i < x; i++ {
			switch {
			case j == 0 && i == 0:
				fmt.Print("o")
			case j == 0 && i == x-1:
				fmt.Print("o")
			case j == y-1 && i == 0:
				fmt.Print("o")
			case j == y-1 && i == x-1:
				fmt.Print("o")
			case j == 0 || j == y-1:
				fmt.Print("-")
			case i == 0 || i == x-1:
				fmt.Print("|")
			default:
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
	QuadA(width, height)
}
