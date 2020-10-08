package sort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{nums: []int{77, 1, 9, 9, 8, 0, 7, 1, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HeapSort(tt.args.nums)
			fmt.Println(tt.args.nums)
		})
	}
}
