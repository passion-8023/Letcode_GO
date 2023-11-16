package let_795

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	start, end := 0, 0
	var sum int
	var smallIsContinuous int //记录小于left的数的连续个数
	for end < len(nums) {
		if left <= nums[end] && nums[end] <= right {
			sum += end - start + 1
			end++
			smallIsContinuous = 0
		} else if nums[end] < left {
			smallIsContinuous++
			if smallIsContinuous > 1 {
				sum += end - smallIsContinuous + 1 - start
			} else {
				sum += end - start
			}
			end++
		} else {
			end++
			start = end
			smallIsContinuous = 0
		}
	}
	return sum
}
