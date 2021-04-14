package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	// queue
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)

	// while !queue.isEmpty()
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, size)
		index := len(ret)
		for i := 0; i < size; i++ {
			node := queue[0]  // queue.peek()
			queue = queue[1:] // queue.pop()
			if index%2 == 0 {
				// 从左到右记录到 level 切片中
				level[i] = node.Val
			} else {
				// 从右到左记录到 level 切片中
				level[size-1-i] = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, level)
	}

	return ret
}
