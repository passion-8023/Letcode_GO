package let_interview_0206

import (
	"github.com/passion-8023/letcodego/data"
)

// 方法一：写入数组
func isPalindrome(head *data.ListNode) bool {
	var listSlice []int

	curr := head
	for curr != nil {
		listSlice = append(listSlice, curr.Val)
		curr = curr.Next
	}

	for left, right := 0, len(listSlice)-1; left < right; left, right = left+1, right-1 {
		if listSlice[left] != listSlice[right] {
			return false
		}
	}

	return true
}

// 方法二：翻转后半部分，然后再比较
func isPalindrome2(head *data.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 找到中间节点
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 翻转后半部分
	list := reverseList(slow.Next)

	curr1, curr2 := head, list
	for curr2 != nil {
		if curr1.Val != curr2.Val {
			return false
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	return true
}

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

// 方法三：
func isPalindrome3(head *data.ListNode) bool {
	curr := head
	var callbackFunc func(node *data.ListNode) bool
	callbackFunc = func(node *data.ListNode) bool {
		if node != nil {
			if !callbackFunc(node.Next) {
				return false
			}

			if curr.Val != node.Val {
				return false
			}
			curr = curr.Next
		}
		return true
	}

	return callbackFunc(head)
}
