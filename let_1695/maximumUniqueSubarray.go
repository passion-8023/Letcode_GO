package let_1695

import "fmt"

func maximumUniqueSubarray(nums []int) int {
	sum := nums[0]
	numsMap := map[int]int{
		nums[0]: 0,
	}

	maxSum := sum
	curr := 1
	//i := 1
	for curr < len(nums) {
		//fmt.Println(i)
		//i++
		val, ok := numsMap[nums[curr]]
		if !ok {
			sum += nums[curr]
			numsMap[nums[curr]] = curr
			curr++
		} else {
			maxSum = max(maxSum, sum)
			tmp := val + 1
			fmt.Println(tmp)
			sum = nums[tmp]
			numsMap[nums[curr]] = curr
			curr = tmp
		}
	}

	return max(maxSum, sum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
