package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		currLevelSize := len(queue)

		for i := 0; i < currLevelSize; i++ {
			currNode := queue[0]
			queue = queue[1:]

			// Call the function f with the current node's data
			_, err := f(currNode.Data)
			if err != nil {
				return
			}

			if currNode.Left != nil {
				queue = append(queue, currNode.Left)
			}

			if currNode.Right != nil {
				queue = append(queue, currNode.Right)
			}
		}
	}
}
