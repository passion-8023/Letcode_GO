package let_lcr_123

import (
	"fmt"
	"github.com/passion-8023/letcodego/data"
	"testing"
)

func Test_reverseBookList(t *testing.T) {
	slice := data.ListNodeBySlice([]int{6, 3, 4, 1})

	list := reverseBookList(slice)

	fmt.Println(list)
}
