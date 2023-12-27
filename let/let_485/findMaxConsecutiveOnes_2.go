package let_485

func findMaxConsecutiveOnes_2(nums []int) int {
	var zeroIndex []int

	for i, num := range nums {
		if num == 0 {
			zeroIndex = append(zeroIndex, i)
		}
	}

	zeroIndex = append(zeroIndex, len(nums))

	if len(zeroIndex) == 0 {
		return len(nums)
	}

	var res int
	for i, index := range zeroIndex {
		if i == 0 {
			res = max(res, index)
		} else {
			res = max(res, index-zeroIndex[i-1]-1)
		}
	}
	return res
}
