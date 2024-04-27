package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	current := root

	for current != nil {

		if current.Data == elem {
			return current
		}
		if elem < current.Data {
			current = current.Left
		}
		if elem > current.Data {
			current = current.Right
		}
	}
	return nil
}
