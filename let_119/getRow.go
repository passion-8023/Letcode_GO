package let_119

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}

	res := make([]int, rowIndex+1)
	res[0], res[1], res[rowIndex-1], res[rowIndex] = 1, rowIndex, rowIndex, 1

	for i, val := range res {
		if i <= rowIndex/2 && val == 0 {
			res[i] = res[i-1] * (rowIndex - i + 1) / i
		}
		if i > rowIndex/2 && val == 0 {
			res[i] = res[rowIndex-i]
		}
	}
	return res
}

func getRow_2(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}

	res := make([]int, rowIndex+1)
	res[0], res[1], res[rowIndex-1], res[rowIndex] = 1, rowIndex, rowIndex, 1

	for i := 2; i <= rowIndex/2; i++ {
		res[i] = res[i-1] * (rowIndex - i + 1) / i
		if rowIndex-i == i {
			continue
		}
		res[rowIndex-i] = res[i]
	}
	return res
}
