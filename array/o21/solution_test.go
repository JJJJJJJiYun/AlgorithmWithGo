package o21

import (
	"fmt"
	"testing"
)

func TestReorderOddAndEven(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{nums: []int{9, 7, 8, 4, 3, 9, 6, 8, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReorderOddAndEven(tt.args.nums)
			fmt.Println(got)
		})
	}
}
