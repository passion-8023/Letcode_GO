package let_interview_0805

import "testing"

func Test_multiply(t *testing.T) {
	type args struct {
		A int
		B int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"测试用例1",
			args{
				A: 1,
				B: 10,
			},
			10,
		},
		{
			"测试用例2",
			args{
				A: 3,
				B: 4,
			},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply2(tt.args.A, tt.args.B); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
