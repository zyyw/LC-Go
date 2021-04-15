package tree

/*
	root2leaf, 找出所有 path sum == target sum 的路径
*/

func pathSum(root *TreeNode, targetSum int) [][]int {
	ret := make([][]int, 0)
	path := make([]int, 0)

	pathSumHelper(root, &ret, path, targetSum)
	return ret
}

func pathSumHelper(root *TreeNode, ret *[][]int, path []int, sum int) {
	if root == nil {
		return
	}

	path = append(path, root.Val)
	sum -= root.Val

	// check leaf node
	if root.Left == nil && root.Right == nil {
		if sum == 0 {
			onePath := make([]int, len(path))
			copy(onePath, path)
			*ret = append(*ret, onePath)
		}
		return
	}

	pathSumHelper(root.Left, ret, path, sum)
	pathSumHelper(root.Right, ret, path, sum)
}
