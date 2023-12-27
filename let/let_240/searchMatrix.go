package let_240

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	var i int = m - 1
	var j int = n - 1
	for i >= 0 && j >= 0 {
		if target == matrix[i][j] {
			return true
		} else if target < matrix[i][j] {
			i = i/2
			j = j/2
		} else {

		}
	}
	return false
}
