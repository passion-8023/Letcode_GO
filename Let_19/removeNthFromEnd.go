package Let_19


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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{
		Val:  0,
		Next: head,
	}
	q := newHead
	p := newHead
	for i := 0; i < n+1; i++ {
		
	}
}
