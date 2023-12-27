package let_141

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}

/**
哈希表
 */
func hasCycle(head *ListNode) bool {
	tmp := map[*ListNode]struct{}{}
	for head != nil {
		if _, ok := tmp[head]; ok {
			return true
		}
		tmp[head] = struct{}{}
		head = head.Next
	}
	return false
}

/**
快慢指针
 */
func hasCycle_1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	low, fast := head, head.Next
	for low != fast {
		if fast == nil || fast.Next == nil {
			return false
		}
		low = low.Next
		fast = fast.Next.Next
	}
	return true
}
