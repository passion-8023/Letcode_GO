package let_interview_1617

import (
	"fmt"
	"testing"
)

func Test_maxSubArray(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	result := maxSubArray(nums)

	fmt.Println(result)
}
