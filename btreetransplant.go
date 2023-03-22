package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	if node == root {
		return rplc
	}

	parent := node.Parent
	if parent.Left == node {
		parent.Left = rplc
	} else {
		parent.Right = rplc
	}

	if rplc != nil {
		rplc.Parent = parent
	}

	return root
}
