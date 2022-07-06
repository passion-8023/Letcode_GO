package let_1480

import "testing"

func TestRunningSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	result := RunningSum(nums)
	t.Log(result)
}
