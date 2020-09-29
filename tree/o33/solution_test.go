package o33

import "testing"

func TestIsPostOrderTraversal(t *testing.T) {
	type args struct {
		t []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{t: []int{3, 7, 13, 12, 5, 20, 16, 15}},
			want: true,
		},
		{
			name: "case2",
			args: args{t: []int{7, 4, 6, 5}},
			want: false,
		},
		{
			name: "case3",
			args: args{t: []int{7, 4}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPostOrderTraversal(tt.args.t); got != tt.want {
				t.Errorf("IsPostOrderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
