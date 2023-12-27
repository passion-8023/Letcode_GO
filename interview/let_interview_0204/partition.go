package let_interview_0204

import "github.com/passion-8023/letcodego/data"

// 题目要求说 不需要保留每个分区中各节点的初始相对位置
// 所以我们把小的往前放，把大于等于的往后放
func partition(head *data.ListNode, x int) *data.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var result *data.ListNode       // 最终结果链表
	var resultEndNod *data.ListNode // 结果链表的最后一个节点

	curr := head
	for curr != nil {
		node := &data.ListNode{Val: curr.Val}
		if result == nil {
			result = node
			resultEndNod = result
		} else {
			if curr.Val < x {
				// 头插
				node.Next = result
				result = node
			} else {
				// 尾插
				resultEndNod.Next = node
				resultEndNod = resultEndNod.Next
			}
		}
		curr = curr.Next
	}

	return result
}
