package let_498

import "fmt"

func findDiagonalOrder(mat [][]int) []int {
	result := []int{mat[0][0]}

	m := len(mat)
	n := len(mat[0])

	if m == 1 && n == 1 {
		return result
	}

	// 定义上下左右方向
	var (
		up    = -1
		down  = 1
		left  = -1
		right = 1
	)

	flag := "right_up" // right_up left_down

	var (
		i = 0
		j = 0
	)

	for {
		fmt.Println("i: ", i, " j: ", j, " flag: ", flag)
		if i == (m-1) && j == (n-1) {
			break
		}

		// 先移动
		if flag == "right_up" {
			i += up
			j += right
		}

		if flag == "left_down" {
			i += down
			j += left
		}

		fmt.Println("i: ", i, " j: ", j, " flag: ", flag)

		if i < 0 {
			flag = "left_down"
			if j == n {
				// 向下移动两个 向左移动一个
				i += down * 2
				j += left
			} else {
				// 向下移动一个
				i += down
			}
		}

		if i == m {
			flag = "right_up"
			// 向右移动两个 向上移动一个
			i += up
			j += right * 2
		}

		if j < 0 {
			flag = "right_up"
			if i == m {
				// 向右移动两个 向上移动一个
				i += up
				j += right * 2
			} else {
				// 向右移动一个
				j += right
			}
		}

		if j == n {
			flag = "left_down"
			// 向下移动两个 向左移动一个
			i += down * 2
			j += left
		}

		fmt.Println("=========")
		result = append(result, mat[i][j])
	}

	return result
}
