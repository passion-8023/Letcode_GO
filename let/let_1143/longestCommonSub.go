package let_1143

//回溯算法思想
func LongestCommonSubsequence(text1 string, text2 string) int {
	maxNum := 0
	var handlerFun func(str1 string, str2 string, i, j, res int) int
	handlerFun = func(str1 string, str2 string, i, j, res int) int {
			if i == len(str1) || j == len(str2) {
				if res > maxNum {
					maxNum = res
				}
				return maxNum
			}

			if str1[i] == str2[j] {
				handlerFun(str1, str2, i+1, j+1, res+1)
			} else {
				handlerFun(str1, str2, i+1, j, res)
				handlerFun(str1, str2, i, j+1, res)
			}
			return maxNum
	}

	res := handlerFun(text1, text2, 0, 0, 0)
	return res
}

//动态规划思想
func LongestCommonSubsequenceByDp(text1 string, text2 string) int {
	n, m := len(text1), len(text2)
	var maxlcs = make([][]int, n)
	for i := 0; i < n; i++ {
		column := make([]int, m)
		maxlcs[i] = column
	}
	for j := 0; j < m; j++ {
		if text1[0] == text2[j] {
			maxlcs[0][j] = 1
		} else if j != 0 {
			maxlcs[0][j] = maxlcs[0][j-1]
		} else {
			maxlcs[0][j] = 0
		}
	}
	for i := 0; i < n; i++ {
		if text1[i] == text2[0] {
			maxlcs[i][0] = 1
		} else if i != 0 {
			maxlcs[i][0] = maxlcs[i-1][0]
		} else {
			maxlcs[i][0] = 0
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if text1[i] == text2[j] {
				maxlcs[i][j] = maxlcs[i-1][j-1] + 1
			} else {
				maxlcs[i][j] = max(maxlcs[i-1][j], maxlcs[i][j-1])
			}
		}
	}
	return maxlcs[n-1][m-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
