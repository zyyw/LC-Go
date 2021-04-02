package LC_Go

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	// 1. dp[i][j] 表示 s[i:j] 是否为 palindrome
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	// 2. 初始化
	// dp[i][i] = true; dp[i - 1][i] = (s[i-1] == s[i])
	start, end := 0, 0
	dp[0][0] = true
	for i := 1; i < len(s); i++ {
		dp[i][i] = true
		dp[i-1][i] = s[i-1] == s[i]
		if dp[i-1][i] {
			start, end = i-1, i
		}
	}

	// 3. 递推公式
	// dp[i][j] = (dp[i + 1][j - 1] && s[i] == s[j])
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 2; j < len(s); j++ {
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			if dp[i][j] && j-i > end-start {
				start, end = i, j
			}
		}
	}

	return s[start : end+1]
}