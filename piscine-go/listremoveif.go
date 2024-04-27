package piscine

/*
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}
*/

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return // If the list is empty, there's nothing to remove
	}

	// Handle the case where the Head node needs to be removed
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next // Move the Head to the next node
	}

	current := l.Head // Start traversal from the Head

	// Traverse the list and remove nodes with matching data
	for current != nil && current.Next != nil {
		if current.Next.Data == data_ref {
			// If the next node's data matches `data_ref`, remove it
			current.Next = current.Next.Next

			// If the removed node was the Tail, update the Tail pointer
			if current.Next == nil {
				l.Tail = current
			}
		} else {
			// Move to the next node
			current = current.Next
		}
	}
}
