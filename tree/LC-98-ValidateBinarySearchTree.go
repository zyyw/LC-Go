package tree

// 方法一：迭代法
func isValidBST(root *TreeNode) bool {
	var prev *TreeNode

	// stack
	stk := make([]*TreeNode, 0)
	for root != nil || len(stk) > 0 {
		if root != nil {
			stk = append(stk, root) // stk.push()
			root = root.Left
		} else {
			root = stk[len(stk)-1] // stk.peek()
			stk = stk[:len(stk)-1] // stk.pop()
			if prev != nil && prev.Val >= root.Val {
				return false
			}
			prev = root
			root = root.Right
		}
	}
	return true
}

// 方法二：递归法
func isValidBST2(root *TreeNode) bool {
	return isValidBSTHelper(root, nil, nil)
}

func isValidBSTHelper(root, minNode, maxNode *TreeNode) bool {
	if root == nil {
		return true
	}

	if minNode != nil && minNode.Val >= root.Val {
		return false
	}
	if maxNode != nil && root.Val >= maxNode.Val {
		return false
	}

	return isValidBSTHelper(root.Left, minNode, root) && isValidBSTHelper(root.Right, root, maxNode)
}
