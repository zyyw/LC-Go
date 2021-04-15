package tree

func sortedArrayToBST(nums []int) *TreeNode {
	return convertHelper(nums, 0, len(nums)-1)
}

func convertHelper(nums []int, start, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	root := &TreeNode{
		Val:   nums[mid],
		Left:  nil,
		Right: nil,
	}
	root.Left = convertHelper(nums, start, mid-1)
	root.Right = convertHelper(nums, mid+1, end)

	return root
}
