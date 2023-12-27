package let_392

import "testing"

func TestIsSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	const s = "axc"
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"第一种", args{"abc", "ahbgdc"}, true},
		{"第二种", args{"axc", "ahbgdc"}, false},
		{"第三种", args{"", "ahbgdc"}, true},
		{"第四种", args{"ss", ""}, false},
		{"第五种", args{"bb", "ahbgdc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
