package let_interview_0201

import (
	"fmt"
	"github.com/passion-8023/letcodego/data"
	"testing"
)

func Test_removeDuplicateNodes(t *testing.T) {
	listNode := data.ListNodeBySlice([]int{1, 2, 3, 3, 2, 1})

	nodes := removeDuplicateNodes(listNode)

	node := data.ShowListNode(nodes)
	fmt.Println(node)
}
