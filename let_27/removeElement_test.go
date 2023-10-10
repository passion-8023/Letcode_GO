package let_27

import "testing"

func Test_removeElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	element := removeElement(nums, 3)
	t.Logf("数组：%+v  数量：%d", nums, element)
}
