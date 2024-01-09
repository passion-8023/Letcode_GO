package let_206

import (
	"github.com/passion-8023/letcodego/data"
	"reflect"
	"testing"
)

func Test_reverseList2(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			"测试1",
			args{head: data.ListNodeBySlice([]int{1, 2, 3, 4, 5})},
			data.ListNodeBySlice([]int{5, 4, 3, 2, 1}),
		},
		{
			"测试2",
			args{head: data.ListNodeBySlice([]int{1, 2})},
			data.ListNodeBySlice([]int{2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList3(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList2() = %v, want %v", got, tt.want)
			}
		})
	}
}
