package let_1646

func getMaximumGenerated(n int) int {
	n = n + 1
	if n <= 2 {
		return 1
	}
	num := make([]int, n)
	num[0],num[1] = 1, 1
	maxNum := 1
	for i := 1; i < n; i++ {
		if 2*i == n || 2*i+1 == n {
			return maxNum
		}
		num[2*i] = num[i]
		if (num[2*i] > maxNum) {
			maxNum = num[2*i]
		}
		num[2*i+1] = num[i] + num[i+1]
		if (num[2*i+1] > maxNum) {
			maxNum = num[2*i+1]
		}
	}
	return maxNum
}
