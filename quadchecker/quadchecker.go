package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// executeQuadFunc executes a quad function with given dimensions and captures its output.
func executeQuadFunc(executablePath, x, y string) (string, error) {
	cmd := exec.Command(executablePath, x, y)
	output, err := cmd.Output()
	if err != nil {
		return "", err // Return the error to handle it properly
	}
	return string(output), nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No command-line arguments provided.")
		return
	}
	// Now, we're sure at least one argument is present.
	inputStr := os.Args[1]

	lines := strings.Split(inputStr, "\n")
	rows := len(lines) - 1 // the final \n after the shape
	cols := 0

	for _, line := range lines {
		if len(line) > cols {
			cols = len(line)
		}
	}

	xRef, yRef := strconv.Itoa(cols), strconv.Itoa(rows)
	// Assuming quad executables are in the same directory as the program
	execDir, _ := filepath.Abs(filepath.Dir("./"))
	quadExecutables := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}

	var matches []string
	for _, quad := range quadExecutables {
		quadPath := filepath.Join(execDir, quad)
		output, err := executeQuadFunc(quadPath, xRef, yRef)
		if err != nil {
			fmt.Printf("Error executing quad function '%s': %s\n", quad, err)
			continue // Skip this quad function on error
		}

		if output == inputStr {
			matches = append(matches, fmt.Sprintf("[%s] [%s] [%s]", quad, xRef, yRef))
		}
	}
	if len(matches) == 0 {
		fmt.Println("Not a quad function.")
	} else {
		sort.Strings(matches)
		joinedstring := strings.Join(matches, " || ") + "\n"
		fmt.Print(joinedstring)
	}
}
