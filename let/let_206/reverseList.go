package let_206

import (
	"github.com/passion-8023/letcodego/data"
)

type ListNode = data.ListNode

func reverseList(head *data.ListNode) *data.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var node1, node2 *data.ListNode = nil, head

	for node2 != nil {
		next := node2.Next
		node2.Next = node1
		node1 = node2
		node2 = next
	}

	return node1
}

// 递归
func reverseList2(head *ListNode) *ListNode {
	var resultList *ListNode
	var currList *ListNode

	var callbackFunc func(node *ListNode) *ListNode

	callbackFunc = func(node *ListNode) *ListNode {
		if node == nil {
			return node
		}

		callbackFunc(node.Next)

		if resultList == nil {
			resultList = &ListNode{Val: node.Val}
			currList = resultList
		} else {
			currList.Next = &ListNode{Val: node.Val}
			currList = currList.Next
		}
		return resultList
	}

	return callbackFunc(head)
}

func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	resultList := reverseList3(head.Next)

	head.Next.Next = head

	head.Next = nil

	return resultList
}
