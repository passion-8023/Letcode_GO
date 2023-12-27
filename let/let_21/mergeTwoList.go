package let_21

import "github.com/passion-8023/letcodego/data"

func MergeTwoLists(list1 *data.ListNode, list2 *data.ListNode) *data.ListNode {
	head := &data.ListNode{}
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

func MergeTwoLists2(list1 *data.ListNode, list2 *data.ListNode) *data.ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = MergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists2(list1, list2.Next)
		return list2
	}
}
