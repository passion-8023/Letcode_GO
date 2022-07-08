package let_21

func MergeTwoLists(list1 *data.ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	cur := head
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			cur.Next = list2
			list2 = list2.Next
		} else {
			cur.Next = list1
			list1 = list1.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return head.Next
}
