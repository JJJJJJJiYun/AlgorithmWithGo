package t236

import (
	"reflect"
	"testing"
)

func Test_lowestCommonAncestor(t *testing.T) {
	t1 := &TreeNode{Val: 3}
	t2 := &TreeNode{Val: 5}
	t3 := &TreeNode{Val: 1}
	t4 := &TreeNode{Val: 6}
	t5 := &TreeNode{Val: 2}
	t6 := &TreeNode{Val: 0}
	t7 := &TreeNode{Val: 8}
	t8 := &TreeNode{Val: 7}
	t9 := &TreeNode{Val: 4}
	t1.Left = t2
	t1.Right = t3
	t2.Left = t4
	t2.Right = t5
	t5.Left = t8
	t5.Right = t9
	t3.Left = t6
	t3.Right = t7
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
				root: t1,
				p:    t2,
				q:    t9,
			},
			want: t2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
