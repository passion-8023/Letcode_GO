package Let_2

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//定义一个头节点
	head := &ListNode{Val: 0}
	current := head
	isAdd := 0 //是否需要进位

	for l1 != nil || l2 != nil || isAdd != 0 {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}
		value := v1 + v2 + isAdd
		value, isAdd = value%10, value/10
		if current.Next == nil {
			current.Next = &ListNode{Val: value}
			current = current.Next
		}
	}
	return head.Next
}