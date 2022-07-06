package let_724

import "testing"

func TestPivotIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"第一种", args{nums: []int{1, 7, 3, 6, 5, 6}}, 3},
		{"第二种", args{nums: []int{1, 2, 3}}, -1},
		{"第三种", args{nums: []int{2, 1, -1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PivotIndex(tt.args.nums); got != tt.want {
				t.Errorf("PivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
