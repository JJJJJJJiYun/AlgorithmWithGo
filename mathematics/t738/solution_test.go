package t738

import "testing"

func Test_monotoneIncreasingDigits(t *testing.T) {
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
			args: args{n: 10},
			want: 9,
		},
		{
			name: "case2",
			args: args{n: 1234},
			want: 1234,
		},
		{
			name: "case3",
			args: args{n: 321},
			want: 299,
		},
		{
			name: "case4",
			args: args{n: 1333322},
			want: 1299999,
		},
		{
			name: "case5",
			args: args{n: 668841},
			want: 667999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := monotoneIncreasingDigits(tt.args.n); got != tt.want {
				t.Errorf("monotoneIncreasingDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
