package tree

// 方法一：递归法
func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	postorderHelper(root, &ret)
	return ret
}

func postorderHelper(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	postorderHelper(root.Left, ret)
	postorderHelper(root.Right, ret)
	*ret = append(*ret, root.Val)
}

// 方法二：迭代法
func postorderTraversal2(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}

	// stack
	stk := make([]*TreeNode, 0)
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		if root != nil {
			stk = append(stk, root) // stk.push()
			root = root.Left
		} else {
			root = stk[len(stk)-1] // stk.peek()
			if root.Right == nil || root.Right == prev {
				// 没有右节点 或者 右子树已经遍历过
				stk = stk[:len(stk)-1] // stk.pop()
				ret = append(ret, root.Val)
				prev = root
				root = nil
			} else {
				// 有右子树，并且右子树没有遍历过
				root = root.Right
			}
		}
	}

	return ret
}
