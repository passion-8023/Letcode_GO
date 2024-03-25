package let_222

import (
	"testing"
)

func Test_countNodes(t *testing.T) {

	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"测试用例1",
			args{root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
					Right: &TreeNode{
						Val: 5,
					},
				},
				Right: &TreeNode{
					Val: 3,
					Left: &TreeNode{
						Val: 6,
					},
					Right: nil,
				},
			}},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countNodes_2(tt.args.root); got != tt.want {
				t.Errorf("countNodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
