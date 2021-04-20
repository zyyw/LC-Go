package arrays

/*
problem statement:
给定一个整数数组，把所有的 0 都移到数组的右边。 非 0 数字 之间的相对顺序要保持不变。
==> 把所有非 0 数字都移到数组的左边，而且非 0 数字之间的相对顺序要保持不变。
*/

func moveZero(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	// 3, 0, 1, 0, 4
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] != 0 {
			// swap(nums, left, right)
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

	return nums
}
