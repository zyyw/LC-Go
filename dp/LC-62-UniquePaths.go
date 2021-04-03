package dp

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	} else if m == 1 || n == 1 {
		return 1
	}

	// 1. dp[i][j]
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 2. initialization
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}

	// 3. dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
		}
	}

	// 4. return
	return dp[m - 1][n - 1]
}

///// 优化空间: O(2n)
func uniquePaths2(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	} else if m == 1 || n == 1 {
		return 1
	}

	prev := make([]int, n)
	cur := make([]int, n)
	for i := 0; i < n; i++ {
		prev[i] = 1
		cur[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			cur[j] = cur[j - 1] + prev[j]
		}
		prev, cur = cur, prev
	}

	return prev[n-1]
}