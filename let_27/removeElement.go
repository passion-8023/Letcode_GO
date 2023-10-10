package let_27

func removeElement(nums []int, val int) int {
	left, right := 0, len(nums)-1

	index := 0
	for left <= right {
		if nums[left] != val {
			left++
			index++
		} else {
			if nums[right] == val {
				right--
			} else {
				nums[left], nums[right] = nums[right], nums[left]
				left++
				right--
				index++
			}
		}
	}
	return index
}
