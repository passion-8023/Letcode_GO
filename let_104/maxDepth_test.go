package let_104

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxDepth(tt.args.root); got != tt.want {
				t.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
