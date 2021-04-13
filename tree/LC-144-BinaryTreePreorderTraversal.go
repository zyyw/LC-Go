package tree

// 方法一：递归法
func preorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	preorderHelper(root, &ret)

	return ret
}

func preorderHelper(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	*ret = append(*ret, root.Val)
	preorderHelper(root.Left, ret)
	preorderHelper(root.Right, ret)
}

// 方法二：迭代法
func preorderTraversal2(root *TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}

	// stack
	stk := make([]*TreeNode, 0)
	for root != nil || len(stk) > 0 {
		if root != nil {
			ret = append(ret, root.Val)
			stk = append(stk, root) // stk.push()
			root = root.Left
		} else {
			root = stk[len(stk)-1] // stk.peek()
			stk = stk[:len(stk)-1] // stk.pop()
			root = root.Right
		}
	}

	return ret
}
