package o51

import "testing"

func TestReversePair(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{nums: []int{7, 5, 6, 4}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReversePair(tt.args.nums); got != tt.want {
				t.Errorf("ReversePair() = %v, want %v", got, tt.want)
			}
		})
	}
}
