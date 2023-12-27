package let_561

import "sort"

func arrayPairSum(nums []int) int {
	sort.Ints(nums)

	var ans int
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}

	return ans
}
