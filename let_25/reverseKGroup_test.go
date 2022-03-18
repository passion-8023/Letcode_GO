package let_25

import (
	"github.com/passion-8023/letcodego/data"
	"testing"
)

func TestReverse(t *testing.T) {
	listNode := data.ListNodeBySlice([]int{1,2,3,4,5,6,7,8})
	group := ReverseKGroup(listNode, 3)
	t.Log(data.ShowListNode(group))
}
