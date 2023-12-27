package let_interview_0204

import (
	"fmt"
	"github.com/passion-8023/letcodego/data"
	"testing"
)

func Test_partition(t *testing.T) {
	listNode := data.ListNodeBySlice([]int{1, 4, 3, 2, 5, 2})

	node := partition(listNode, 3)

	fmt.Println(data.ShowListNode(node))
}
