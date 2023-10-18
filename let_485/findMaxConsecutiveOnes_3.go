package let_485

func findMaxConsecutiveOnes_3(nums []int) int {
	res := 0
	current := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			current++
			res = max(res, current)
		} else {
			res = max(res, current)
			current = 0
		}
	}
	return res
}
