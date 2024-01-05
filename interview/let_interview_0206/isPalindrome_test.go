package let_interview_0206

import (
	"github.com/passion-8023/letcodego/data"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *data.ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			"测试示例1",
			args{
				head: data.ListNodeBySlice([]int{1, 2, 2, 1}),
			},
			true,
		},
		{
			"测试示例2",
			args{
				head: data.ListNodeBySlice([]int{1, 2}),
			},
			false,
		},
		{
			"测试示例3",
			args{
				head: data.ListNodeBySlice([]int{1, 2, 3, 4, 5, 4, 3, 2, 1}),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome3(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
