package LC_Go

func twoSum(nums []int, target int) []int {
	ret := make([]int, 2)
	if len(nums) < 2 {
		return ret
	}

	// map to store (num_value, index)
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		index, ok := m[target - nums[i]]
		if ok {
			// found (index, i) pair whose sum is target
			ret[0] = index
			ret[1] = i
			return ret
		} else {
			// not found
			m[nums[i]] = i
		}
	}

	return ret
}
