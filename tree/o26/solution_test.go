package o26

import (
	"LeetCode/utils"
	"testing"
)

func TestPartOfTree(t *testing.T) {
	type args struct {
		part *utils.TreeNode
		root *utils.TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case1",
			args: args{
				part: &utils.TreeNode{
					Val: 3,
					Left: &utils.TreeNode{
						Val: 4,
					},
					Right: &utils.TreeNode{
						Val: 5,
					},
				},
				root: &utils.TreeNode{
					Val: 1,
					Left: &utils.TreeNode{
						Val:  2,
						Left: &utils.TreeNode{Val: 4},
					},
					Right: &utils.TreeNode{
						Val: 3,
						Left: &utils.TreeNode{
							Val: 4,
						},
						Right: &utils.TreeNode{
							Val: 5,
						},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PartOfTree(tt.args.part, tt.args.root); got != tt.want {
				t.Errorf("PartOfTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
