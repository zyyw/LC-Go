package tree

// 思路；层序遍历
func rightSideView(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}

	// queue
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]  // queue.peek()
			queue = queue[1:] // queue.pop()
			if i == size-1 {
				ret = append(ret, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left) // queue.push()
			}
			if node.Right != nil {
				queue = append(queue, node.Right) // queue.push()
			}
		}
	}

	return ret
}
