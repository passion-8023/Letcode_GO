package let_offer_58

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "示例1",
			args: args{
				s: "I am wang xin",
			},
			want: "xin wang am I",
		},
		{
			name: "示例2",
			args: args{
				s: "   I am     wang xin   ",
			},
			want: "xin wang am I",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
