package tree

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	//queue
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		level := make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]  // queue.peek()
			queue = queue[1:] // queue.pop()
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left) // queue.push()
			}
			if node.Right != nil {
				queue = append(queue, node.Right) // queue.push()
			}
		}
		ret = append(ret, level)
	}

	return ret
}
