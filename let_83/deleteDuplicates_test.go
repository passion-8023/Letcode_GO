package let_83

import (
	"testing"
)

/**
 * Definition for singly-linked list.
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head
	for current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	return head
}

func TestDeleteDuplicates(t *testing.T)  {
	head  := ListNode{
		Val:  1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 1,
				Next: nil,
			},
		},
	}

	node := deleteDuplicates(&head)
	for node != nil {
		t.Log(node.Val)
		node = node.Next
	}
}
