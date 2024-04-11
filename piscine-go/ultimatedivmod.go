package piscine

// Insert your comment here
func UltimateDivMod(a *int, b *int) {
	Dres := *a / *b
	Mres := *a % *b
	*a = Dres
	*b = Mres
}
