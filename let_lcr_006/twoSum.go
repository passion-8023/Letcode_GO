package let_lcr_006

func twoSum(numbers []int, target int) []int {
	start, end := 0, len(numbers)-1

	for start < end {
		tmp := numbers[start] + numbers[end]
		if tmp > target {
			end--
		} else if tmp < target {
			start++
		} else {
			return []int{start, end}
		}
	}
	return []int{}
}
