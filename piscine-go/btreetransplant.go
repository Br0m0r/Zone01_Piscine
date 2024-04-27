package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root // If node is nil, there's nothing to replace, return the original root.
	}
	if rplc != nil {
		rplc.Parent = node.Parent // Set rplc's parent to node's parent
	}
	if node.Parent == nil {
		root = rplc // If node is the root, rplc becomes the new root.
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc // If node is a left child, replace it with rplc in its parent's left pointer.
	} else {
		node.Parent.Right = rplc // If node is a right child, replace it with rplc in its parent's right pointer.
	}
	return root // Return the new root of the tree (might be unchanged if node was not the root).
}
