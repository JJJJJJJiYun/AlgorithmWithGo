package o55

import (
	"LeetCode/utils"
	"testing"
)

func TestDepthOfTree(t *testing.T) {
	type args struct {
		root *utils.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{
				root: &utils.TreeNode{
					Val: 1,
					Left: &utils.TreeNode{
						Val: 2,
						Left: &utils.TreeNode{
							Val: 4,
						},
						Right: &utils.TreeNode{
							Val: 5,
							Left: &utils.TreeNode{
								Val: 7,
							},
						},
					},
					Right: &utils.TreeNode{
						Val: 3,
						Right: &utils.TreeNode{
							Val: 6,
						},
					},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DepthOfTree(tt.args.root); got != tt.want {
				t.Errorf("DepthOfTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
