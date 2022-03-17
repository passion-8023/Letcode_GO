package let_21

import (
	"github.com/passion-8023/letcodego/data"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := data.ListNodeBySlice([]int{1,2,4})
	list2 := data.ListNodeBySlice([]int{1,3,4})
	lists := MergeTwoLists(list1, list2)
	t.Log(data.ShowListNode(lists))
}
