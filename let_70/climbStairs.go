package let_70

//动态规划 + 滚动数组
//dp方程 F(x) = F(x-1) + F(x-2)
func climbStairs(n int) int {
	q, p, r := 0, 0, 1
	for i := 1; i <= n; i++ {
		q = p
		p = r
		r = q + p
	}
	return r
}
