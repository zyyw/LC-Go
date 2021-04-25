package arrays

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	// [0, left) 全为 0
	// (right, n-1] 全为 2
	left, right := 0, len(nums)-1
	for i := 0; i <= right; i++ {
		if nums[i] == 0 {
			// swap(nums, left, i)
			nums[left], nums[i] = nums[i], nums[left]
			left++
		} else if nums[i] == 2 {
			// swap(nums, i, right)
			nums[i], nums[right] = nums[right], nums[i]
			right--
			i--
		}
	}
}
