package main

import (
	"io"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printContent(os.Stdin) // for case hello hello ^C
	} else {
		for _, filename := range args {
			file, err := os.Open(filename)
			if err != nil {
				printString("ERROR: open " + filename + ": no such file or directory\n")
				os.Exit(1)
				continue
			}

			printContent(file)
			file.Close()
		}
	}
}

// reads all data from io.Reader and prints it
func printContent(src io.Reader) {
	content, err := io.ReadAll(src)
	if err != nil {
		printString("ERROR: reading from input\n")
		os.Exit(1)
	}
	printString(string(content))
}

func printString(s string) {
	for _, rune := range s {
		z01.PrintRune(rune)
	}
}
