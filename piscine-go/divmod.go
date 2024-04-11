package piscine

// Insert your comment here
func DivMod(a int, b int, div *int, mod *int) {
	Dres := a / b
	Mres := a % b
	*div = Dres
	*mod = Mres
}
