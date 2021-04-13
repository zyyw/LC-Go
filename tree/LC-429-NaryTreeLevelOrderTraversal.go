package tree

func levelOrderOfNaryTree(root *Node) [][]int {
	ret := make([][]int, 0)
	if root == nil {
		return ret
	}

	// queue
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		level := make([]int, 0)
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]  // queue.peek()
			queue = queue[1:] // queue.pop()
			level = append(level, node.Val)
			for _, child := range node.Children {
				if child != nil {
					queue = append(queue, child) // queue.push()
				}
			}
		}
		ret = append(ret, level)
	}

	return ret
}
