package let_209

func minSubArrayLen(target int, nums []int) int {
	left, right, sum := 0, 0, 0
	res := len(nums) + 1
	for right < len(nums) {
		sum += nums[right]
		if sum >= target {
			for sum-nums[left] >= target {
				sum -= nums[left]
				left++
			}
			res = min(res, right-left+1)
		}
		right++
	}

	if res == len(nums)+1 {
		return 0
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
