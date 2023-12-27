package let_23

import (
	"github.com/passion-8023/letcodego/data"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	list1 := data.ListNodeBySlice([]int{1,4,5})
	list2 := data.ListNodeBySlice([]int{1,3,4})
	list3 := data.ListNodeBySlice([]int{2,6})
	lists := MergeKLists([]*ListNode{list1, list2, list3})
	t.Log(data.ShowListNode(lists))
}

func TestMergeKListsByPartition(t *testing.T) {
	list1 := data.ListNodeBySlice([]int{1,4,5})
	list2 := data.ListNodeBySlice([]int{1,3,4})
	list3 := data.ListNodeBySlice([]int{2,6})
	lists := MergeKListsByPartition([]*ListNode{list1, list2, list3})
	t.Log(data.ShowListNode(lists))
}
