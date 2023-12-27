package let_lcr_027

import (
	"fmt"
	"github.com/passion-8023/letcodego/data"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	slice := data.ListNodeBySlice([]int{1, 1, 2, 1})

	list := demo(slice)

	fmt.Println(data.ShowListNode(list))
}
