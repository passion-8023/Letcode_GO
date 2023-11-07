package let_153

import "sort"

func findMin(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	if nums[0] < nums[1] && nums[0] < nums[n-1] {
		return nums[0]
	}

	rotateTimes := 0 // 旋转次数
	for i := n - 1; i >= 0; i-- {
		if nums[0] > nums[i] {
			rotateTimes++
		} else {
			break
		}
	}

	return nums[n-rotateTimes]
}

func findMin2(nums []int) int {
	low, high := 0, len(nums)-1

	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else {
			low = pivot + 1
		}

	}

	return nums[low]
}

func findMin3(nums []int) int {
	sort.Ints(nums)
	return nums[0]
}
