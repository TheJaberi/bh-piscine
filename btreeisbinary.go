package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSubtreeLesser(root.Left, root.Data) &&
		isSubtreeGreater(root.Right, root.Data) &&
		BTreeIsBinary(root.Left) &&
		BTreeIsBinary(root.Right)
}

func isSubtreeLesser(node *TreeNode, val string) bool {
	if node == nil {
		return true
	}
	if node.Data <= val &&
		isSubtreeLesser(node.Left, val) &&
		isSubtreeLesser(node.Right, val) {
		return true
	}
	return false
}

func isSubtreeGreater(node *TreeNode, val string) bool {
	if node == nil {
		return true
	}
	if node.Data > val &&
		isSubtreeGreater(node.Left, val) &&
		isSubtreeGreater(node.Right, val) {
		return true
	}
	return false
}
