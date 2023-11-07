package let_26

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return n
	}

	j := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j-1] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

func removeDuplicates1(nums []int) int {
	i := 0

	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
