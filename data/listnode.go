package data

type ListNode struct {
	Val int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
	}
}

//头插法
func (l *ListNode) HeadInsert(val int) *ListNode {
	node := &ListNode{Val: val}
	if l == nil {
		return node
	}
	node.Next = l
	l = node
	return node
}

//翻转链表
func (l *ListNode) ReverseList() *ListNode {
	if l == nil || l.Next == nil {
		return l
	}
	node1 := l
	node2 := l.Next
	for node2 != nil {
		next := node2.Next
		node2.Next = node1
		node1 = node2
		node2 = next
	}
	l.Next = nil
	l = node1
	return node1
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