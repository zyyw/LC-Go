package binary_search

func mySqrt(x int) int {
	ret := -1
	start, end := 0, x

	for start <= end {
		mi := (start + end) / 2
		if mi * mi <= x {
			ret = mi
			start = mi + 1
		} else {
			end = mi - 1
		}
	}

	return ret
}
