package let_779

import "testing"

//func Test_callback(t *testing.T) {
//	type args struct {
//		n int
//	}
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		// TODO: Add test cases.
//		{
//			"示例1",
//			args{2},
//			"01",
//		},
//		{
//			"示例2",
//			args{3},
//			"0110",
//		},
//		{
//			"示例3",
//			args{30},
//			"01101001",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := callback(tt.args.n); got != tt.want {
//				t.Errorf("callback() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_kthGrammar(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"示例1",
			args{2, 1},
			0,
		},
		{
			"示例2",
			args{3, 2},
			1,
		},
		{
			"示例3",
			args{30, 434991989},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthGrammar(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("kthGrammar() = %v, want %v", got, tt.want)
			}
		})
	}
}
