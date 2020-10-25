package t845

import "testing"

func Test_longestMountain(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name: "case1",
		//	args: args{array: []int{2, 1, 4, 7, 3, 2, 5}},
		//	want: 5,
		//},
		//{
		//	name: "case2",
		//	args: args{array: []int{2, 2, 2}},
		//	want: 0,
		//},
		{
			name: "case3",
			args: args{array: []int{1, 3, 2, 4, 5, 4, 3}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestMountain(tt.args.array); got != tt.want {
				t.Errorf("longestMountain() = %v, want %v", got, tt.want)
			}
		})
	}
}
