package o36

import (
	"AlgorithWithGo/utils"
	"fmt"
	"testing"
)

func TestConvertTreeToLinkedList(t *testing.T) {
	type args struct {
		root *utils.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *utils.TreeNode
	}{
		{
			name: "case1",
			args: args{
				root: &utils.TreeNode{
					Val: 15,
					Left: &utils.TreeNode{
						Val: 5,
						Left: &utils.TreeNode{
							Val: 3,
						},
						Right: &utils.TreeNode{
							Val: 12,
							Left: &utils.TreeNode{
								Val: 7,
							},
							Right: &utils.TreeNode{
								Val: 13,
							},
						},
					},
					Right: &utils.TreeNode{
						Val: 16,
						Right: &utils.TreeNode{
							Val: 20,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertTreeToLinkedList(tt.args.root)
			for got != nil {
				fmt.Printf("%v ", got.Val)
				got = got.Right
			}
		})
	}
}
