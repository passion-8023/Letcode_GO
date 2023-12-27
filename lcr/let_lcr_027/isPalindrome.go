package let_lcr_027

import (
	"github.com/passion-8023/letcodego/data"
)

// 遍历写入切片，然后双指针判断是否回文
func isPalindrome(head *data.ListNode) bool {
	var tmpSlice []int

	for head != nil {
		tmpSlice = append(tmpSlice, head.Val)
		head = head.Next
	}

	// 从两端往中间
	//for left, right := 0, len(tmpSlice)-1; left < right; left, right = left+1, right-1 {
	//	if tmpSlice[left] != tmpSlice[right] {
	//		return false
	//	}
	//}

	// 从中间往两端
	for i := len(tmpSlice)/2 - 1; i >= 0; i-- {
		opp := len(tmpSlice) - 1 - i
		if tmpSlice[i] != tmpSlice[opp] {
			return false
		}
	}

	return true
}
