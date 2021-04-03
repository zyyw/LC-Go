package dp


func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if m == 0 || n == 0 {
		return 0
	}

	// 1. dp[i][j]
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 2. initialization
	if obstacleGrid[0][0] == 0 {
		dp[0][0] = 1
	} else {
		dp[0][0] = 0
	}
	// 初始化 第一行
	for j := 1; j < n; j++ {
		if dp[0][j - 1] == 1 && obstacleGrid[0][j] == 0 {
			dp[0][j] = 1
		} else {
			dp[0][j] = 0
		}
	}
	// 初始化 第一列
	for i := 1; i < m; i++ {
		if dp[i - 1][0] == 1 && obstacleGrid[i][0] == 0 {
			dp[i][0] = 1
		} else {
			dp[i][0] = 0
		}
	}

	// 3. if obstacleGrid[i][j] == 0: dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
	//    else: dp[i][j] = 0
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
			} else {
				dp[i][j] = 0
			}
		}
	}

	// 4. return
	return dp[m - 1][n - 1]
}
