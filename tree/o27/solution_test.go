package o27

import (
	"LeetCode/utils"
	"testing"
)

func TestMirrorTree(t *testing.T) {
	type args struct {
		root *utils.TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				root: &utils.TreeNode{
					Val: 8,
					Left: &utils.TreeNode{
						Val: 6,
						Left: &utils.TreeNode{
							Val: 5,
						},
						Right: &utils.TreeNode{
							Val: 7,
						},
					},
					Right: &utils.TreeNode{
						Val: 10,
						Left: &utils.TreeNode{
							Val: 9,
						},
						Right: &utils.TreeNode{
							Val: 11,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MirrorTree(tt.args.root)
			utils.HierarchicalTraversalWithBreak(got)
		})
	}
}
