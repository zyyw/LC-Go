package LC_Go

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	ret := 0
	start, end := 0, len(height)-1
	for start < end {
		water := 0
		if height[start] < height[end] {
			water = height[start] * (end - start)
			start++
		} else {
			water = height[end] * (end - start)
			end--
		}
		if water > ret {
			ret = water
		}
	}

	return ret
}