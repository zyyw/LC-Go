package LC_Go

import "sort"

func threeSum(nums []int) [][]int {
	ret := make([][]int, 0)
	if len(nums) < 3 {
		return ret
	}

	// sort the slice
	sort.Ints(nums)

	// for... start, end
	m := make(map[[3]int]bool)
	for i := 0; i < len(nums); i++ {
		targetSum := 0 - nums[i]
		start, end := i+1, len(nums)-1
		for start < end {
			sum := nums[start] + nums[end]
			if sum == targetSum {
				arr := [3]int{nums[i], nums[start], nums[end]}
				m[arr] = true
				start++
				end--
			} else if sum < targetSum {
				start++
			} else {
				end--
			}
		}
	}

	// for map
	for key, _ := range m {
		ret = append(ret, []int{key[0], key[1], key[2]})
	}
	return ret
}
