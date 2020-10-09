package o54

import (
	"AlgorithWithGo/utils"
	"testing"
)

func TestKthNode(t *testing.T) {
	type args struct {
		root *utils.TreeNode
		k    int
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
					Val: 5,
					Left: &utils.TreeNode{
						Val: 3,
						Left: &utils.TreeNode{
							Val: 2,
						},
						Right: &utils.TreeNode{
							Val: 4,
						},
					},
					Right: &utils.TreeNode{
						Val: 7,
						Left: &utils.TreeNode{
							Val: 6,
						},
						Right: &utils.TreeNode{
							Val: 8,
						},
					},
				},
				k: 7,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KthNode(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("KthNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPreOrderTraversal(t *testing.T) {
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
					Val: 5,
					Left: &utils.TreeNode{
						Val: 3,
						Left: &utils.TreeNode{
							Val: 2,
						},
						Right: &utils.TreeNode{
							Val: 4,
						},
					},
					Right: &utils.TreeNode{
						Val: 7,
						Left: &utils.TreeNode{
							Val: 6,
						},
						Right: &utils.TreeNode{
							Val: 8,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PreOrderTraversal(tt.args.root)
		})
	}
}

func TestInOrderTraversal(t *testing.T) {
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
					Val: 5,
					Left: &utils.TreeNode{
						Val: 3,
						Left: &utils.TreeNode{
							Val: 2,
						},
						Right: &utils.TreeNode{
							Val: 4,
						},
					},
					Right: &utils.TreeNode{
						Val: 7,
						Left: &utils.TreeNode{
							Val: 6,
						},
						Right: &utils.TreeNode{
							Val: 8,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InOrderTraversal(tt.args.root)
		})
	}
}

func TestPostOrderTraversal(t *testing.T) {
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
					Val: 5,
					Left: &utils.TreeNode{
						Val: 3,
						Left: &utils.TreeNode{
							Val: 2,
						},
						Right: &utils.TreeNode{
							Val: 4,
						},
					},
					Right: &utils.TreeNode{
						Val: 7,
						Left: &utils.TreeNode{
							Val: 6,
						},
						Right: &utils.TreeNode{
							Val: 8,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PostOrderTraversal(tt.args.root)
		})
	}
}
