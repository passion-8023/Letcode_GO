package let_795

import (
	"fmt"
	"testing"
)

func Test_numSubarrayBoundedMax(t *testing.T) {
	nums := []int{2, 9, 2, 5, 6}
	left := 2
	right := 8

	got := numSubarrayBoundedMax(nums, left, right)
	fmt.Println(got)
}
