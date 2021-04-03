package LC_Go

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	lowestPrice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		curProfit := prices[i] - lowestPrice
		if curProfit > maxProfit {
			maxProfit = curProfit
		}
		if prices[i] < lowestPrice {
			lowestPrice = prices[i]
		}
	}

	return maxProfit
}
