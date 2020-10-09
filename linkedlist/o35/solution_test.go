package o35

import (
	"AlgorithWithGo/utils"
	"testing"
)

func TestCopyComplexLinkedList(t *testing.T) {
	first := &utils.ComplexLinkedNode{Val: 1}
	second := &utils.ComplexLinkedNode{Val: 2}
	third := &utils.ComplexLinkedNode{Val: 3}
	fourth := &utils.ComplexLinkedNode{Val: 4}
	fifth := &utils.ComplexLinkedNode{Val: 5}
	first.Next = second
	second.Next = third
	third.Next = fourth
	fourth.Next = fifth
	first.Sibling = third
	second.Sibling = fifth
	fourth.Sibling = second
	type args struct {
		first *utils.ComplexLinkedNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				first: first,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CopyComplexLinkedList(tt.args.first)
			utils.PrintComplexLinkedList(got)
		})
	}
}
