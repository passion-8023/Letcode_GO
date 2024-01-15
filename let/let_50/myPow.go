package let_50

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}

	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		return 1 / (x * myPow(x, -n-1))
	} else {
		return x * myPow(x, n-1)
	}
}

func myPow1(x float64, n int) float64 {
	var callbackFunc func(x float64, n int) float64
	callbackFunc = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}

		val := callbackFunc(x, n/2)

		if n%2 == 0 {
			return val * val
		}
		return x * val * val
	}

	if n < 0 {
		return 1.0 / callbackFunc(x, -n)
	}

	return callbackFunc(x, n)
}
