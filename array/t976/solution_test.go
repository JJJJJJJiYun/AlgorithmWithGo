package t976

import "testing"

func Test_largestPerimeter(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name: "case1",
		//	args: args{[]int{3, 6, 2, 3}},
		//	want: 8,
		//},
		//{
		//	name: "case2",
		//	args: args{[]int{3, 2, 3, 4}},
		//	want: 10,
		//},
		{
			name: "case3",
			args: args{[]int{3, 4, 15, 2, 9, 4}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestPerimeter(tt.args.arr); got != tt.want {
				t.Errorf("largestPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}
