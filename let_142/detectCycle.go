package let_142

import "github.com/passion-8023/letcodego/data"

// 哈希表
func detectCycle(head *data.ListNode) *data.ListNode {
	tmpHas := map[*data.ListNode]struct{}{}

	for head != nil {
		if _, ok := tmpHas[head]; ok {
			return head
		}

		tmpHas[head] = struct{}{}

		head = head.Next
	}

	return nil
}
