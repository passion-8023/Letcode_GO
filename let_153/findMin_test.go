package let_153

import "testing"

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"测试用例一",
			args{nums: []int{3, 4, 5, 1, 2}},
			1,
		},
		{
			"测试用例二",
			args{nums: []int{4, 5, 6, 7, 0, 1, 2}},
			0,
		},
		{
			"测试用例三",
			args{nums: []int{11, 13, 15, 17}},
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			min := findMin(tt.args.nums)
			t.Log(min)
			//if got := findMin(tt.args.nums); got != tt.want {
			//	t.Errorf("findMin() = %v, want %v", got, tt.want)
			//}
		})
	}
}
