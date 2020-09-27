package o22

import (
	"LeetCode/utils"
	"fmt"
	"testing"
)

func TestLastKNode(t *testing.T) {
	type args struct {
		first *utils.LinkedNode
		k     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				first: &utils.LinkedNode{
					Val: 1,
					Next: &utils.LinkedNode{
						Val: 2,
						Next: &utils.LinkedNode{
							Val: 3,
							Next: &utils.LinkedNode{
								Val: 4,
								Next: &utils.LinkedNode{
									Val: 5,
								},
							},
						},
					},
				},
				k: 6,
			},
		},
		{
			name: "case1",
			args: args{
				first: &utils.LinkedNode{
					Val: 1,
				},
				k: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastKNode(tt.args.first, tt.args.k); got != nil {
				fmt.Println(got.Val)
			} else {
				fmt.Printf("nil")
			}
		})
	}
}
