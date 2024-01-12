package let_50

import "fmt"

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	fmt.Println(x, n)
	if n < 0 {
		return 1 / (x * myPow(x, n+1))
	} else {
		return x * myPow(x, n-1)
	}
}
