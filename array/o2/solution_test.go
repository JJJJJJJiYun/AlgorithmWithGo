package o2

import "testing"

func TestFindNumInTwoDimensionArray(t *testing.T) {
	type args struct {
		nums   [][]int
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				nums: [][]int{
					{1, 2, 8, 9},
					{2, 4, 9, 12},
					{4, 7, 10, 13},
					{6, 8, 11, 15},
				},
				target: 7,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindNumInTwoDimensionArray(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("FindNumInTwoDimensionArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
