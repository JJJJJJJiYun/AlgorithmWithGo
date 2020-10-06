package o52

import (
	"LeetCode/utils"
	"reflect"
	"testing"
)

func TestFirstDuplicateNode(t *testing.T) {
	n1 := &utils.LinkedNode{Val: 1}
	n2 := &utils.LinkedNode{Val: 2}
	n3 := &utils.LinkedNode{Val: 3}
	n4 := &utils.LinkedNode{Val: 4}
	n5 := &utils.LinkedNode{Val: 5}
	n6 := &utils.LinkedNode{Val: 6}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n5.Next = n6
	n6.Next = n4
	type args struct {
		node1 *utils.LinkedNode
		node2 *utils.LinkedNode
	}
	tests := []struct {
		name string
		args args
		want *utils.LinkedNode
	}{
		{
			name: "case1",
			args: args{
				node1: n1,
				node2: n5,
			},
			want: n4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstDuplicateNode(tt.args.node1, tt.args.node2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstDuplicateNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
