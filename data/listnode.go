package data

type ListNode struct {
	Val int
	Next *ListNode
}

func ListNodeBySlice(nums []int) *ListNode {
	head := &ListNode{}
	current := head

	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head.Next
}

func ShowListNode(head *ListNode) []int {
	current := head
	var nums []int
	for current != nil {
		nums = append(nums, current.Val)
		current = current.Next
	}
	return nums
}