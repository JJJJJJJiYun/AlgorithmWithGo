package o23

import (
	"LeetCode/utils"
	"fmt"
	"testing"
)

func TestFindFirstNodeInLoop(t *testing.T) {
	first := &utils.LinkedNode{Val: 1}
	second := &utils.LinkedNode{Val: 2}
	third := &utils.LinkedNode{Val: 3}
	fourth := &utils.LinkedNode{Val: 4}
	fifth := &utils.LinkedNode{Val: 5}
	first.Next = second
	second.Next = third
	third.Next = fourth
	fourth.Next = fifth
	fifth.Next = second
	type args struct {
		first *utils.LinkedNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{first: first},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindFirstNodeInLoop(tt.args.first); got != nil {
				fmt.Println(got.Val)
			} else {
				fmt.Printf("nil")
			}
		})
	}
}
