package tree

// 方法一：递归法
func maxDepthOfNaryTree(root *Node) int {
	if root == nil {
		return 0
	}

	depth := 0
	for _, child := range root.Children {
		res := maxDepthOfNaryTree(child)
		if res > depth {
			depth = res
		}
	}

	return depth + 1
}

// 方法二：层序遍历法
func maxDepthOfNaryTree2(root *Node) int {
	if root == nil {
		return 0
	}

	level := 0
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		level++
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]  // queue.peek()
			queue = queue[1:] // queue.pop()
			for _, child := range node.Children {
				if child != nil {
					queue = append(queue, child) // queue.push()
				}
			}
		}
	}

	return level
}
