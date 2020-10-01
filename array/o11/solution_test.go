package o11

import "testing"

func TestFindMinInRotatedArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name: "case1",
		//	args: args{nums: []int{3, 4, 5, 1, 2}},
		//	want: 1,
		//},
		//{
		//	name: "case2",
		//	args: args{nums: []int{3, 4, 5, 6, 7, 1, 2}},
		//	want: 1,
		//},
		//{
		//	name: "case3",
		//	args: args{nums: []int{5, 6, 7, 1, 2, 3, 4}},
		//	want: 1,
		//},
		//{
		//	name: "case4",
		//	args: args{nums: []int{1, 2, 3, 4, 5}},
		//	want: 1,
		//},
		{
			name: "case5",
			args: args{nums: []int{1, 0, 1, 1, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMinInRotatedArray(tt.args.nums); got != tt.want {
				t.Errorf("FindMinInRotatedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
