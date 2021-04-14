package tree

/*
	判断一棵树是否是平衡的二叉树
*/

func isBalanced(root *TreeNode) bool {
	ret := [1]bool{true}
	isBalancedHelper(root, &ret)
	return ret[0]
}

func isBalancedHelper(root *TreeNode, ret *[1]bool) int {
	if (*ret)[0] == false {
		return 0
	}
	if root == nil {
		return 0
	}

	leftHeight := isBalancedHelper(root.Left, ret)
	rightHeight := isBalancedHelper(root.Right, ret)
	if leftHeight-rightHeight > 1 || rightHeight-leftHeight > 1 {
		(*ret)[0] = false
	}

	if leftHeight > rightHeight {
		return 1 + leftHeight
	}
	return 1 + rightHeight
}
