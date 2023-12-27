package let_147

import (
	"github.com/passion-8023/letcodego/data"
	"sort"
)

// 超出内容限制了，难受
func insertionSortList(head *data.ListNode) *data.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var tmpSlice []int

	curr := head
	for curr != nil {
		tmpSlice = append(tmpSlice, curr.Val)
	}

	sort.Ints(tmpSlice)

	newListNode := &data.ListNode{}
	for _, val := range tmpSlice {
		newListNode.Next = &data.ListNode{
			Val:  val,
			Next: nil,
		}
	}

	return newListNode.Next
}

func insertionSortList2(head *data.ListNode) *data.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newListNode := &data.ListNode{
		Val:  -1,
		Next: head,
	}

	start, curr := head, head.Next
	for curr != nil {
		if start.Val <= curr.Val {
			start = start.Next
		} else {
			tmpNewListCurr := newListNode
			for tmpNewListCurr.Next.Val <= curr.Val {
				tmpNewListCurr = tmpNewListCurr.Next
			}

			start.Next = curr.Next
			curr.Next = tmpNewListCurr.Next
			tmpNewListCurr.Next = curr
		}

		curr = start.Next
	}

	return newListNode.Next
}
