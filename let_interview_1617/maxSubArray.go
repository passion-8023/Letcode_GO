package let_interview_1617

//	func maxSubArray(nums []int) int {
//		maxNum := nums[0]
//		sumSlice := []int{maxNum}
//		for i := 1; i < len(nums); i++ {
//			var currSum int
//			if sumSlice[i-1] < 0 {
//				currSum = nums[i]
//			} else {
//				currSum = nums[i] + sumSlice[i-1]
//			}
//
//			sumSlice = append(sumSlice, currSum)
//
//			maxNum = max(maxNum, currSum)
//
//		}
//
//		return maxNum
//	}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	result := nums[0]

	maxSubArrayHandler(nums, &result)

	return result
}

func maxSubArrayHandler(nums []int, result *int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	tmpNum := maxSubArrayHandler(nums[:l-1], result)
	if tmpNum > *result {
		result = &tmpNum
	}

	if tmpNum+nums[l-1] > nums[l-1] {
		return tmpNum + nums[l-1]
	} else {
		return nums[l-1]
	}
}
