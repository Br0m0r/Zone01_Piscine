package main

import "github.com/01-edu/z01"

type Door struct {
	state string
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func OpenDoor(ptr *Door) bool {
	PrintStr("Door Opening...")
	ptr.state = "OPEN"
	return true
}

func CloseDoor(ptrDoor *Door) bool {
	PrintStr("Door Closing...")
	ptrDoor.state = "CLOSE"
	return true
}

func IsDoorOpen(ptrDoor *Door) bool {
	PrintStr("is the Door opened ?")
	return ptrDoor.state == "OPEN"
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return ptrDoor.state == "CLOSE"
}

func main() {
	door := &Door{
		state: "CLOSE",
	}
	OpenDoor(door)
	z01.PrintRune('\n')
	if IsDoorClose(door) {
		OpenDoor(door)
	}
	z01.PrintRune('\n')
	if IsDoorOpen(door) {
		z01.PrintRune('\n')
		CloseDoor(door)
	}
	if door.state == "OPEN" {
		CloseDoor(door)
	}
	z01.PrintRune('\n')
}
