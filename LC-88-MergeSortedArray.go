package LC_Go

func mergeSortedArray(nums1 []int, m int, nums2 []int, n int)  {
	if len(nums2) == 0 {
		return
	}

	k := m + n - 1
	for m > 0 && n > 0 {
		if nums1[m - 1] < nums2[n - 1] {
			nums1[k] = nums2[n - 1]
			n--
		} else {
			nums1[k] = nums1[m - 1]
			m--
		}
		k--
	}
	for n > 0 {
		nums1[k] = nums2[n - 1]
		n--
		k--
	}
}
