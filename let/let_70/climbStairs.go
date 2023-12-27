package let_70

// 动态规划 + 滚动数组
// dp方程 F(x) = F(x-1) + F(x-2)
func climbStairs(n int) int {
	q, p, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		q = p
		p = r
		r = q + p
	}
	return r
}

// 递归
func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}

	return climbStairs2(n-1) + climbStairs(n-2)
}

// 递归+去重项
func climbStairs3(n int) int {
	tmpResultMap := make(map[int]int)

	return climbStairs3Handler(n, tmpResultMap)
}

func climbStairs3Handler(n int, resultMap map[int]int) int {
	if val, ok := resultMap[n]; ok {
		return val
	}

	var result int

	if n == 1 {
		result = 1
	}

	if n == 2 {
		result = 2
	}

	if n > 2 {
		result = climbStairs3Handler(n-1, resultMap) + climbStairs3Handler(n-2, resultMap)
	}

	resultMap[n] = result
	return result
}

// 动态规划
func climbStairs4(n int) int {
	dp := []int{1, 2, 3}

	if n <= 3 {
		return dp[n-1]
	}

	for i := 4; i < n; i++ {
		curr := dp[1] + dp[2]
		dp[0], dp[1] = dp[1], dp[2]
		dp[2] = curr
	}

	return dp[2]
}
