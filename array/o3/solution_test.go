package o3

import "testing"

func TestFindDuplicateNum(t *testing.T) {
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
			args: args{nums: []int{2, 3, 1, 0, 2, 5, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicateNum(tt.args.nums); got != tt.want {
				t.Errorf("FindDuplicateNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDuplicateNum2(t *testing.T) {
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
			args: args{nums: []int{2, 3, 1, 0, 2, 5, 3}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDuplicateNum2(tt.args.nums); got != tt.want {
				t.Errorf("FindDuplicateNum2() = %v, want %v", got, tt.want)
			}
		})
	}
}
