package tree

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	// queue
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		cnt := 0
		for i := 0; i < size; i++ {
			prevCnt := cnt
			cur := queue[0]   // queue.peek()
			queue = queue[1:] // queue.pop()
			if cur.Left != nil {
				queue = append(queue, cur.Left) // queue.push()
				if cur.Left.Val == x || cur.Left.Val == y {
					cnt++
				}
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right) // queue.push()
				if cur.Right.Val == x || cur.Right.Val == y {
					cnt++
				}
			}
			if cnt == 2 {
				// cnt == 2 表示两个节点都在本层中。
				// 如果 prevCnt == 0，表示两个节点共父节点，非 cousins 节点
				// 如果 prevCnt == 1，表示两个节点不共父节点，为 cousins 节点
				return prevCnt != 0
			}
		}
		if cnt == 1 {
			// 这层找到一个节点 x 或 y，那么它们两个节点不再同一层，肯定非 cousins 节点
			return false
		}
	}
	return false
}
