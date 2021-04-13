package tree

// 方法一：递归法
func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	inorderHelper(root, &ret)
	return ret
}

func inorderHelper(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}
	inorderHelper(root.Left, ret)
	*ret = append(*ret, root.Val)
	inorderHelper(root.Right, ret)
}

// 方法二：迭代法
func inorderTraversal2(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}

	// stack
	stk := make([]*TreeNode, 0)
	for root != nil || len(stk) > 0 {
		if root != nil {
			stk = append(stk, root) // stk.push()
			root = root.Left
		} else {
			root = stk[len(stk)-1] // stk.peek()
			stk = stk[:len(stk)-1] // stk.pop()
			ret = append(ret, root.Val)
			root = root.Right
		}
	}

	return ret
}
