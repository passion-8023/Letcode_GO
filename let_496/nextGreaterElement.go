package let_496

//[1,3,5,2,4]
//[6,5,4,3,2,1,7]
//
//[4,1,2]
//[1,3,4,2]

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	for i, v := range nums1 {
		j := 0
		for j < len(nums2) && nums2[j] != v {
			j++
		}
		k := j + 1
		for k < len(nums2) && nums2[k] < nums2[j] {
			k++
		}
		if k < len(nums2) {
			nums1[i] = nums2[k]
		} else {
			nums1[i] = -1
		}
	}
	return nums1
}


//单调栈
func nextGreaterElement_stack(nums1 []int, nums2 []int) []int {
	stack := []int{}
	mp := map[int]int{}
	for i := len(nums2)-1; i >= 0; i-- {
		num := nums2[i]
		for len(stack) > 0 && num >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			mp[num] = stack[len(stack)-1]
		} else {
			mp[num] = -1
		}
		stack = append(stack, num)
	}
	for i, num := range nums1 {
		nums1[i] = mp[num]
	}
	return nums1

}
