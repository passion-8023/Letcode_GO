package let_24

import (
	"github.com/passion-8023/letcodego/data"
	"testing"
)

func TestSwapPairsByCallback(t *testing.T) {
	listNode := data.ListNodeBySlice([]int{1, 2, 3, 4})
	node := SwapPairsByCallback(listNode)
	t.Log(data.ShowListNode(node))
}



func TestSwapPairs(t *testing.T) {
	listNode := data.ListNodeBySlice([]int{1, 2, 3, 4})
	node := SwapPairs(listNode)
	t.Log(data.ShowListNode(node))
}
