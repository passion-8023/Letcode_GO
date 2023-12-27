package let_lcr_088

func minCostClimbingStairs(cost []int) int {
	size := len(cost)

	dp := make([]int, size+1)

	for i := 2; i <= size; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[size]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
