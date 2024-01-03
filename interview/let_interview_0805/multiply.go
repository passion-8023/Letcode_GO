package let_interview_0805

// 递归加法
func multiply(A int, B int) int {
	if A == 0 || B == 0 {
		return 0
	}

	if A < B {
		A--
		B += multiply(A, B)
		return B
	} else {
		B--
		A += multiply(A, B)
		return A
	}
}

// 递归位运算
func multiply2(A int, B int) int {
	if B == 1 {
		return A
	}

	if B&1 == 1 {
		// B是奇数
		return A + multiply2(A<<1, B>>1)
	} else {
		return multiply2(A<<1, B>>1)
	}
}
