package let_485

func findMaxConsecutiveOnes(nums []int) int {
	left, right := 0, 0
	res := 0

	for right < len(nums) {
		if nums[right] == 0 {
			right++
			left = right
			continue
		}
		res = max(res, right-left+1)
		right++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
