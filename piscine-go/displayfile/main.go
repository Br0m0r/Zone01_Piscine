package main

import (
	"fmt"
	"os"
)

func main() {
	if len((os.Args)) != 2 {
		if len(os.Args) < 2 {
			fmt.Println("File name missing")
		} else {
			fmt.Println("Too many arguments")
		}
		return

	}
	filename := os.Args[1]

	// os to handle the opening,error handling,reading the file
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Almost there!!")
		return
	}

	fmt.Print(string(data)) // dont forget to convert
}
