package LC_Go

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	start, end := 0, m*n-1

	// binary search between [start, end]
	for start <= end {
		mi := (start + end) / 2
		row, col := mi/n, mi%n
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] < target {
			start = mi + 1
		} else {
			end = mi - 1
		}
	}
	return false
}