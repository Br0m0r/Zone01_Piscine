package piscine

func ListReverse(l *List) {
	current := l.Head
	var previous *NodeL
	var next *NodeL

	for current != nil {
		next = current.Next
		current.Next = previous
		previous = current
		current = next
	}
	l.Head, l.Tail = l.Tail, l.Head
}
