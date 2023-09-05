package let_48

import (
	"fmt"
	"math"
)

func rotate(matrix [][]int) {
	n := len(matrix)
	fmt.Println("n:", n)

	num := int(math.Ceil(float64(n) / 2))

	fmt.Println(num)

	m := 0
	for i := 0; i < num; i++ {
		fmt.Println("i==:", i)

		if i == 0 {
			m = n - 1
		}

		fmt.Println("m: ", m)

		for j := i; j < m; j++ {
			fmt.Println("j:", j)
			t := 4 - i
			if t < 3 {
				t = 3
			}

			if j == m-1 {
				m = j
			}
			callFunc(t, n, i, j, matrix[i][j], matrix)
		}
	}
}

func callFunc(num, n, i, j, val int, matrix [][]int) int {
	if num < 0 {
		return num
	}
	x, y := new_x_y(n, i, j)

	if x == i && y == j {
		return num
	}

	tmpVal := matrix[x][y]
	matrix[x][y] = val

	num--
	return callFunc(num, n, x, y, tmpVal, matrix)
}

func new_x_y(n, i, j int) (int, int) {
	newi := j
	newj := n - 1 - i
	return newi, newj
}
