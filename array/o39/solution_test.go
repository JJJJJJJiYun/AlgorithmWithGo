package o39

import "testing"

func TestMoreThanHalfNum(t *testing.T) {
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
			args: args{nums: []int{1, 2, 2, 3, 2, 2, 4, 2}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MoreThanHalfNum(tt.args.nums); got != tt.want {
				t.Errorf("MoreThanHalfNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
