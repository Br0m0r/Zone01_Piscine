package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	// Check if the number of arguments is not 9
	if len(args) != 9 {
		fmt.Println("Error")
		return
	}

	// Generating the 9x9 grid
	var grid [9][9]int
	for i, arg := range args {
		if len(arg) != 9 {
			fmt.Println("Error")
			return
		}
		for j, char := range arg {
			if char == '.' {
				grid[i][j] = 0
			} else if char >= '1' && char <= '9' {
				number := int(char - '0')
				grid[i][j] = number
			} else {
				fmt.Println("Error")
				return
			}
		}
	}

	// Run the backtracking solver
	if backtracking(&grid) {
		printGrid(grid)
	} else {
		fmt.Println("Error")
	}
}

func backtracking(grid *[9][9]int) bool {
	if gridComplete(*grid) {
		return true
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid[i][j] == 0 {
				candidates := getCandidates(*grid, i, j)
				for _, value := range candidates {
					grid[i][j] = value
					if backtracking(grid) {
						return true
					}
					grid[i][j] = 0
				}
				return false
			}
		}
	}
	return false
}

func getCandidates(grid [9][9]int, rowIndex int, columnIndex int) []int {
	candidates := []int{}
	for num := 1; num <= 9; num++ {
		if !contains(grid[rowIndex], num) && !isColumnDuplicate(grid, columnIndex, num) && !isSubgridDuplicate(grid, rowIndex, columnIndex, num) {
			candidates = append(candidates, num)
		}
	}
	return candidates
}

func contains(row [9]int, number int) bool {
	for _, value := range row {
		if value == number {
			return true
		}
	}
	return false
}

func isColumnDuplicate(grid [9][9]int, columnIndex int, number int) bool {
	for i := 0; i < 9; i++ {
		if grid[i][columnIndex] == number {
			return true
		}
	}
	return false
}

func isSubgridDuplicate(grid [9][9]int, rowIndex int, columnIndex int, number int) bool {
	startRow := 3 * (rowIndex / 3)
	startCol := 3 * (columnIndex / 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[startRow+i][startCol+j] == number {
				return true
			}
		}
	}
	return false
}

func printGrid(grid [9][9]int) {
	for _, row := range grid {
		for j, number := range row {
			fmt.Print(number)
			if j != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println(" ")
	}
}

func gridComplete(grid [9][9]int) bool {
	for _, row := range grid {
		for _, number := range row {
			if number == 0 {
				return false
			}
		}
	}
	return true
}
