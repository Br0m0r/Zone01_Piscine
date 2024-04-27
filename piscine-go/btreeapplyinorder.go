package piscine

func BTreeApplyInorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	// Traverse the left subtree
	BTreeApplyInorder(root.Left, f)

	// Apply the function to the current node's data
	f(root.Data)

	// Traverse the right subtree
	BTreeApplyInorder(root.Right, f)
}
