package let_interview_0205

import (
	"github.com/passion-8023/letcodego/data"
)

func addTwoNumbers(l1 *data.ListNode, l2 *data.ListNode) *data.ListNode {
	var icon int

	resultNode := &data.ListNode{}
	curr := resultNode
	for l1 != nil || l2 != nil || icon > 0 {
		tmpNode := &data.ListNode{}

		var val1 int
		var val2 int

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		tmpVal := val1 + val2 + icon

		tmpNode.Val = tmpVal % 10

		icon = tmpVal / 10

		curr.Next = tmpNode
		curr = curr.Next

	}

	return resultNode.Next
}
