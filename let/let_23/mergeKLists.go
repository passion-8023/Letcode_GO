package let_23

import (
	"github.com/passion-8023/letcodego/data"
	"github.com/passion-8023/letcodego/let/let_21"
)

type ListNode = data.ListNode

// 顺序合并
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	ans := lists[0]
	for i := 1; i < len(lists); i++ {
		ans = let_21.MergeTwoLists(ans, lists[i])
	}
	return ans
}

// 分治
func MergeKListsByPartition(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) >> 1
	return let_21.MergeTwoLists(merge(lists, l, mid), merge(lists, mid+1, r))
}
