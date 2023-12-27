package let_interview_0108

//3 * 4
//[1,1,1,1],
//[0,0,1,1],
//[1,1,1,1]

func setZeroes(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	var zeroSlice [][]int

	for i := 0; i < m; i++ {
		// 判断每一行是否存在 0
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroSlice = append(zeroSlice, []int{i, j})
			}
		}
	}

	for _, item := range zeroSlice {
		// 行
		for i := 0; i < n; i++ {
			matrix[item[0]][i] = 0
		}

		// 列
		for j := 0; j < m; j++ {
			matrix[j][item[1]] = 0
		}
	}
}
