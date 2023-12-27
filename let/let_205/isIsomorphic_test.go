package let_205

import "testing"

func TestIsIsomorphic(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"第一种", args{"egg", "add"}, true},
		{"第二种", args{"foo", "bar"}, false},
		{"第三种", args{"paper", "title"}, true},
		{"第四种", args{"bbbaaaba", "aaabbbba"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIsomorphic(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("IsIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
