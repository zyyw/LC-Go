package tree

func diameterOfBinaryTree(root *TreeNode) int {
	ret := [1]int{0}
	height(root, &ret)
	return ret[0]
}

/*
	height 返回子树的树高
*/
func height(root *TreeNode, ret *[1]int) int {
	if root == nil {
		return 0
	}

	retLeft := height(root.Left, ret)
	retRight := height(root.Right, ret)

	if retLeft+retRight > (*ret)[0] {
		(*ret)[0] = retLeft + retRight
	}

	if retLeft > retRight {
		return retLeft + 1
	}
	return retRight + 1
}
