package let_interview_0806

func hanota(A []int, B []int, C []int) []int {
	var move func(int, *[]int, *[]int, *[]int)
	move = func(num int, from, mid, to *[]int) {

		if num == 1 {
			*to = append(*to, (*from)[len(*from)-1])
			*from = (*from)[:len(*from)-1]
			return
		}

		move(num-1, from, to, mid)
		move(1, from, mid, to)
		move(num-1, mid, from, to)
	}
	move(len(A), &A, &B, &C)
	return C

}
