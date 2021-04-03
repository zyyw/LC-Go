package LC_Go

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	totalProfit := 0
	for i := 1; i < len(prices); i++ {
		curProfit := prices[i] - prices[i - 1]
		if curProfit > 0 {
			totalProfit = totalProfit + curProfit
		}
	}

	return totalProfit
}
