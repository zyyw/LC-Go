package LC_Go

import "sort"

func merge(intervals [][]int) [][]int {
	ret := make([][]int, 0)
	if len(intervals) == 0 {
		return ret
	}
	sort.Sort(Interval(intervals))

	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if end < intervals[i][0] {
			// add (start, end) to ret
			interval := make([]int, 2)
			interval[0], interval[1] = start, end
			ret = append(ret, interval)
			// update start, end
			start, end = intervals[i][0], intervals[i][1]
		} else {
			// intervals[i][0] <= end
			if end <= intervals[i][1] {
				// update end if end <= intervals[i][1]
				end = intervals[i][1]
			}
		}
	}
	interval := make([]int, 2)
	interval[0], interval[1] = start, end
	ret = append(ret, interval)

	return ret
}

type Interval [][]int

func (interval Interval) Len() int {
	return len(interval)
}

func (interval Interval) Swap(i, j int)  {
	interval[i], interval[j] = interval[j], interval[i]
}

func (interval Interval) Less(i, j int) bool {
	return interval[i][0] < interval[j][0]
}