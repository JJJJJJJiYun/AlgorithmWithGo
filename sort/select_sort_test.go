package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{nums: []int{7, 17, 1, 9, 9, 8, 0, 7, 1, 7}},
		},
		{
			name: "case2",
			args: args{nums: []int{}},
		},
		{
			name: "case3",
			args: args{nums: []int{7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SelectSort(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
