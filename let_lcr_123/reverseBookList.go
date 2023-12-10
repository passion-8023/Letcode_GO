package let_lcr_123

import (
	"github.com/passion-8023/letcodego/data"
)

func reverseBookList(head *data.ListNode) []int {
	var result []int

	curr := head

	for curr != nil {
		result = append([]int{curr.Val}, result...)
		curr = curr.Next
	}

	return result
}
