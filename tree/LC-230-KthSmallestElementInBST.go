package tree

import "math"

// 思路：中序遍历，迭代法
func kthSmallest(root *TreeNode, k int) int {
	// stack
	stk := make([]*TreeNode, 0)
	for root != nil || len(stk) > 0 {
		if root != nil {
			stk = append(stk, root) // stk.push()
			root = root.Left
		} else {
			root = stk[len(stk)-1] // stk.peek()
			stk = stk[:len(stk)-1] // stk.pop()
			if k == 1 {
				return root.Val
			}
			k--
			root = root.Right
		}
	}

	return math.MaxInt32
}
