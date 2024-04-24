package piscine

func ListSize(l *List) int {
	counter := 0

	for l.Head != nil {
		counter += 1
		l.Head = l.Head.Next
	}
	return counter
}
