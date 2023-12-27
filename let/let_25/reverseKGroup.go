package let_25

import "github.com/passion-8023/letcodego/data"

type ListNode = data.ListNode

func ReverseKGroup(head *ListNode, k int) *ListNode {
	//哨兵模式
	node := &ListNode{-1, head}
	pre := node
	end := node
	for end.Next != nil {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		start := pre.Next
		next := end.Next
		end.Next = nil
		pre.Next = reverse(start)
		start.Next = next
		pre = start
		end = pre
	}
	return node.Next
}

func reverse(head *ListNode) *ListNode {
	cur := head
	node := &ListNode{}
	for cur != nil {
		next := cur.Next
		cur.Next = node
		node = cur
		cur = next
	}
	return node
}

