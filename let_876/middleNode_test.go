package let_876

import (
	"fmt"
	"github.com/passion-8023/letcodego/data"
	"testing"
)

func Test_middleNode(t *testing.T) {
	listNode := data.ListNodeBySlice([]int{1, 2, 3, 4, 5})

	node := middleNode(listNode)

	showListNode := data.ShowListNode(node)

	fmt.Println(showListNode)
}
