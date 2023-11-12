package let_1695

//func maximumUniqueSubarray(nums []int) int {
//	sum := 0                     // 当前子串的和
//	sumMax := 0                  // 子串的最大和
//	numsMap := make(map[int]int) // key=数组值 value=数组下标
//
//	start, end := 0, 0 // 子串的开始下标和结束下标
//
//	for end < len(nums) {
//		currentValue := nums[end]
//		// 判断当前下标下数组值是否在哈希map里
//		val, ok := numsMap[currentValue]
//		if !ok {
//			// 不存在
//			numsMap[currentValue] = end
//			// 加和
//			sum += currentValue
//			// 继续下一个值
//			end++
//		} else {
//			// 存在
//			numsMap = map[int]int{}
//			// 重置指针位置
//			start = val + 1
//			end = start
//			// 比较当前子串和与最大和，存储最大和
//			sumMax = max(sumMax, sum)
//			sum = 0
//		}
//	}
//
//	return max(sumMax, sum)
//}

func maximumUniqueSubarray2(nums []int) int {
	sum := 0                      // 当前子串的和
	sumMax := 0                   // 子串的最大和
	numsMap := make(map[int]bool) // key=数组值 value=数组下标

	start, end := 0, 0 // 子串的开始下标和结束下标

	for end < len(nums) {
		currentValue := nums[end]

		// 删除重复元素之前的值
		for numsMap[currentValue] {
			sum -= nums[start]
			delete(numsMap, nums[start])
			start++
		}
		// 加和
		sum += currentValue
		sumMax = max(sumMax, sum)
		numsMap[currentValue] = true
		end++
	}

	return max(sumMax, sum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
