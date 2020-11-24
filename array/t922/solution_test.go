package t922

import (
	"fmt"
	"testing"
)

func Test_sortArrayByParityII(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "case1",
			args: args{nums: []int{4, 2, 5, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortArrayByParityII(tt.args.nums)
			fmt.Println(got)
		})
	}
}
