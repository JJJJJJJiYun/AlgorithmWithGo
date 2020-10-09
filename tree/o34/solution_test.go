package o34

import (
	"AlgorithWithGo/utils"
	"testing"
)

func TestTreePathOfNum(t *testing.T) {
	type args struct {
		root   *utils.TreeNode
		target int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				root: &utils.TreeNode{
					Val: 10,
					Left: &utils.TreeNode{
						Val: 5,
						Left: &utils.TreeNode{
							Val: 4,
						},
						Right: &utils.TreeNode{
							Val: 7,
						},
					},
					Right: &utils.TreeNode{
						Val: 12,
					},
				},
				target: 22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TreePathOfNum(tt.args.root, tt.args.target)
		})
	}
}
