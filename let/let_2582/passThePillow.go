package let_2582

func passThePillow(n int, time int) int {
	if time < n {
		return time + 1
	}

	if (time/(n-1))%2 == 0 {
		return time%(n-1) + 1
	} else {
		return n - time%(n-1)
	}
}
