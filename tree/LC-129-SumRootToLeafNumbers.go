package tree

func sumNumbers(root *TreeNode) int {
	return sumNumbersHelper(root, 0)
}

func sumNumbersHelper(root *TreeNode, prefix int) int {
	if root == nil {
		return 0
	}

	sum := prefix*10 + root.Val

	// check leaf node
	if root.Left == nil && root.Right == nil {
		return sum
	}

	return sumNumbersHelper(root.Left, sum) + sumNumbersHelper(root.Right, sum)
}
