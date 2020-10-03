package o42

import "testing"

func TestMaxSubArraySum(t *testing.T) {
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
			args: args{nums: []int{1, 2, -1, -2, 2, 1, -2, 1, 4, -5, 4}},
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxSubArraySum(tt.args.nums); got != tt.want {
				t.Errorf("MaxSubArraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
