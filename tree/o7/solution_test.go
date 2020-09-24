package o7

import (
	"LeetCode/utils"
	"testing"
)

func TestReconstructTreeWithPreAndInOrderTraversal(t *testing.T) {
	type args struct {
		pre []int
		in  []int
	}
	tests := []struct {
		name string
		args args
		want *utils.TreeNode
	}{
		{
			name: "case1",
			args: args{
				pre: []int{1, 2, 4, 7, 3, 5, 6, 8},
				in:  []int{4, 7, 2, 1, 5, 3, 8, 6},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReconstructTreeWithPreAndInOrderTraversal(tt.args.pre, tt.args.in)
			utils.HierarchicalTraversalWithBreak(got)
		})
	}
}
