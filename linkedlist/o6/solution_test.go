package o6

import (
	"LeetCode/utils"
	"testing"
)

func TestReversePrintLinkedList(t *testing.T) {
	type args struct {
		node *utils.LinkedNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{node: &utils.LinkedNode{
				Val: 1,
				Next: &utils.LinkedNode{
					Val: 2,
					Next: &utils.LinkedNode{
						Val:  3,
						Next: nil,
					},
				},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReversePrintLinkedList(tt.args.node)
		})
	}
}
