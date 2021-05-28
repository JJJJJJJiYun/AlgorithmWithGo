package microsoft

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestorOfBinaryTree(t *testing.T) {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node1.Left = node2
	node2.Left = node3
	node1.Right = node4
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "case1",
			args: args{
				root: node1,
				p:    node3,
				q:    node4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestorOfBinaryTree2(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestorOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
