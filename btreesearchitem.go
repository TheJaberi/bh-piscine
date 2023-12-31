package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Data == elem {
		return root
	}
	if res := BTreeSearchItem(root.Left, elem); res != nil {
		return res
	}
	return BTreeSearchItem(root.Right, elem)
}
