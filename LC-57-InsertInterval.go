package LC_Go

func insert(intervals [][]int, newInterval []int) [][]int {
	ret := make([][]int, 0)

	i := 0

	// 1. [i].end < start
	//    追加 ([i].start, [i].end)
	for ; i < len(intervals) && intervals[i][1] < newInterval[0]; i++ {
		interval := make([]int, 2)
		interval[0], interval[1] = intervals[i][0], intervals[i][1]
		ret = append(ret, interval)
	}
	/// NOW: start <= [i].end

	// 2. [i].start <= end
	//    只要 [i].start <= end, 就要一直往右试探找到右边界 Max([i].end, end)
	start, end := newInterval[0], newInterval[1]
	if i < len(intervals) && intervals[i][0] < start {
		start = intervals[i][0]
	}
	// [i].start <= end
	for ; i < len(intervals) && intervals[i][0] <= newInterval[1]; i++ {
		if newInterval[1] < intervals[i][1] {
			// end < [i].end
			end = intervals[i][1]
		}
	}
	// add (start, end) to ret
	interval := make([]int, 2)
	interval[0], interval[1] = start, end
	ret = append(ret, interval)
	/// NOW: end < [i].start

	// 3. end < [i].start
	//    既然 end < [i].start, 追加 ([i].start, [i].end)
	for ; i < len(intervals); i++ {
		interval := make([]int, 2)
		interval[0], interval[1] = intervals[i][0], intervals[i][1]
		ret = append(ret, interval)
	}

	return ret
}
