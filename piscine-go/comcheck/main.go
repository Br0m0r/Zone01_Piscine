package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	for _, r := range arguments {
		if r == "01" || r == "galaxy" || r == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
