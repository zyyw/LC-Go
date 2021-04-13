package tree

// 方法一：递归法
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth < rightDepth {
		return 1 + rightDepth
	}
	return 1 + leftDepth
}

// 方法二：层序遍历
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 0
	for len(queue) != 0 {
		level++
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]  // queue.peek
			queue = queue[1:] // queue.pop
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return level
}
