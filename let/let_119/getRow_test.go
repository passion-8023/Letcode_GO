package let_119

import (
	"testing"
)

func Test_getRow(t *testing.T) {
	type args struct {
		rowIndex int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.\
		{
			name: "用例一",
			args: args{rowIndex: 2},
			want: []int{1, 2, 1},
		},
		{
			name: "用例二",
			args: args{rowIndex: 5},
			want: []int{1, 5, 10, 10, 5, 1},
		},
		{
			name: "用例三",
			args: args{rowIndex: 6},
			want: []int{1, 6, 15, 20, 15, 6, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRow(tt.args.rowIndex)
			t.Log(got)
			//if got := getRow(tt.args.rowIndex); !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("getRow() = %v, want %v", got, tt.want)
			//}
		})
	}
}
