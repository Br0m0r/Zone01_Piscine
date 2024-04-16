package main

import "github.com/01-edu/z01"

type smthng struct {
	x int
	y int
}

func setPoint(ptr *smthng) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &smthng{}
	setPoint(points)
	str := "x = 42, y = 21"
	for _, r := range str {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
