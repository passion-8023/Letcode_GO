package let_5

// 动态规划方法
func longestPalindrome_2(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}

	maxLen := 1
	begin := 0

	var dp [][]bool
	for i := 0; i < l; i++ {
		var tmpDp []bool
		for j := 0; j < l; j++ {
			if i == j {
				tmpDp = append(tmpDp, true)
			} else {
				tmpDp = append(tmpDp, false)
			}
		}
		dp = append(dp, tmpDp)
	}

	for m := 2; m <= l; m++ {
		for i := 0; i < l; i++ {
			j := m + i - 1
			if j >= l {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				begin = i
			}
		}
	}

	return s[begin : begin+maxLen]
}
