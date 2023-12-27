package Let_19

import "github.com/passion-8023/letcodego/data"

type ListNode = data.ListNode
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//双指针方法
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{
		Val:  0,
		Next: head,
	}
	q := newHead
	p := newHead
	for i := 0; i < n+1; i++ {
		q = q.Next
	}
	for q != nil {
		p = p.Next
		q = q.Next
	}
	p.Next = p.Next.Next
	return newHead.Next
}
