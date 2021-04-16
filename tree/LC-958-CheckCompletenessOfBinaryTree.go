package tree

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// queue
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	hasNilChild := false
	for len(queue) > 0 {
		cur := queue[0]   // queue.peek()
		queue = queue[1:] // queue.pop()
		if cur.Left != nil {
			if hasNilChild {
				return false
			}
			queue = append(queue, cur.Left) // queue.push()
		} else {
			hasNilChild = true
		}
		if cur.Right != nil {
			if hasNilChild {
				return false
			}
			queue = append(queue, cur.Right) // queue.push()
		} else {
			hasNilChild = true
		}
	}

	return true
}
