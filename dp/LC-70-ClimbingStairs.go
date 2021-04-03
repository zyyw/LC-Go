package dp

func climbStairs(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	x, y := 1, 2
	for i := 3; i <= n; i++ {
		x, y = y, x+y
	}

	return y
}
