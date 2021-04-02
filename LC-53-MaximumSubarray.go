package LC_Go

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 1. 状态定义：dp[i] - 以 nums[i] 结尾的连续子数组的最大和
	dp := make([]int, len(nums))

	// 2. 初始化
	dp[0] = nums[0]

	// 3. 递推公式：dp[i] = (dp[i-1] > 0 ? dp[i-1] : 0) + nums[i]
	ret := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > ret {
			ret = dp[i]
		}
	}

	// 4. return MAX{dp[i]}
	return ret
}
