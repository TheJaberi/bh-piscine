package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	node := &TreeNode{Data: data}
	if root == nil {
		return node
	}
	if data < root.Data {
		if root.Left == nil {
			root.Left = node
			node.Parent = root
		} else {
			BTreeInsertData(root.Left, data)
		}
	} else if data > root.Data {
		if root.Right == nil {
			root.Right = node
			node.Parent = root
		} else {
			BTreeInsertData(root.Right, data)
		}
	}
	return root
}
