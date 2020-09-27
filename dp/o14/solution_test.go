package o14

import "testing"

func TestCutRope(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				n: 8,
			},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CutRope(tt.args.n); got != tt.want {
				t.Errorf("CutRope() = %v, want %v", got, tt.want)
			}
		})
	}
}
