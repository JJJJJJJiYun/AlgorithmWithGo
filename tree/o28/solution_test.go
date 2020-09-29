package o28

import (
	"LeetCode/utils"
	"testing"
)

func TestIsMirrorTree(t *testing.T) {
	type args struct {
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
						Val: 6,
						Left: &utils.TreeNode{
							Val: 7,
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
			if got := IsMirrorTree(tt.args.root); got != tt.want {
				t.Errorf("IsMirrorTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
