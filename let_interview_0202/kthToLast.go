package let_interview_0202

import "github.com/passion-8023/letcodego/data"

func kthToLast(head *data.ListNode, k int) int {
	size := 0

	curr := head
	for curr != nil {
		size++
		curr = curr.Next
	}

	result := 0

	for i := size - k; i >= 0 && head != nil; i-- {
		if i == 0 {
			result = head.Val
			break
		}
		head = head.Next
	}

	return result
}
