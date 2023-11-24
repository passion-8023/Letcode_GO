package let_118

func genetate(numRows int) [][]int {
	res := make([][]int, numRows)
	res[0] = []int{1}
	for i := 1; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1
		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j] + res[i-1][j-1]
		}
	}
	return res
}

func genetate2(numRows int) [][]int {
	res := make([][]int, numRows)
	if numRows == 1 {
		res[0] = []int{1}
		return res
	}

	if numRows == 2 {
		res[0] = []int{1}
		res[1] = []int{1, 1}
		return res
	}

	res = genetate2(numRows - 1)

	l := len(res)

	var currSlice []int
	for i := 0; i < len(res[l-1])-1; i++ {
		currSlice = append(currSlice, res[l-1][i]+res[l-1][i+1])
	}

	if len(currSlice) > 0 {
		currSlice = append(currSlice, 1)
		tmpCurrSlice := []int{1}
		tmpCurrSlice = append(tmpCurrSlice, currSlice...)
		res = append(res, tmpCurrSlice)
	}

	return res
}
