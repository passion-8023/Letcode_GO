package let_11

func maxArea(height []int) int {
	result := 0
	left := 0
	right := len(height) - 1

	for left != right {
		if height[left] < height[right] {
			result = max(result, height[left]*(right-left))
			left++
		} else {
			result = max(result, height[right]*(right-left))
			right--
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
