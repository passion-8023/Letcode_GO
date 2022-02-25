package let_53

//动态规划方程
//f(i) = max{f(i-1)+nums[i], nums[i]}

func maxSubArray(nums []int) int {
	maxNums := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] + nums[i] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxNums {
			maxNums = nums[i]
		}
	}
	return maxNums
}
