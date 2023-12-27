package let_876

import (
	"github.com/passion-8023/letcodego/data"
)

// 第一种：先求链表长度，然后取中
func middleNode(head *data.ListNode) *data.ListNode {
	curr := head
	size := 0
	for curr != nil {
		size++
		curr = curr.Next
	}

	tmpCurr := head
	i := 0
	for tmpCurr != nil {
		if i == size/2 {
			break
		}
		tmpCurr = tmpCurr.Next
		i++
	}
	return tmpCurr
}

// 方法二：快慢指针
func middleNode2(head *data.ListNode) *data.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
