package let_24

import "github.com/passion-8023/letcodego/data"

type ListNode = data.ListNode

//递归实现
func SwapPairsByCallback(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	head.Next = SwapPairsByCallback(next.Next)
	next.Next = head
	return next
}

//迭代解法
func SwapPairs(head *ListNode) *ListNode {
	//哨兵模式
	dummyHead := &ListNode{0, head}
	temp := dummyHead

	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		temp = node1
	}
	return dummyHead.Next
}
