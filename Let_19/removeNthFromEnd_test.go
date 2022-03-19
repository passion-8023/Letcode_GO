package Let_19

import (
	"github.com/passion-8023/letcodego/data"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	listNode := data.ListNodeBySlice([]int{1, 2, 3, 4, 5})
	result := RemoveNthFromEnd(listNode, 2)
	t.Log(data.ShowListNode(result))
}
