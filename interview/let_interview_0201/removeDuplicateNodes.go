package let_interview_0201

import (
	"github.com/passion-8023/letcodego/data"
)

// 哨兵机制+哈希表
func removeDuplicateNodes(head *data.ListNode) *data.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	nodeValMap := make(map[int]struct{})

	curr := &data.ListNode{
		Val:  -1,
		Next: head,
	}

	for curr.Next != nil {
		if _, ok := nodeValMap[curr.Next.Val]; ok {
			tmpNode := curr.Next.Next
			curr.Next = tmpNode
		} else {
			nodeValMap[curr.Next.Val] = struct{}{}
			curr = curr.Next
		}
	}

	return head
}
