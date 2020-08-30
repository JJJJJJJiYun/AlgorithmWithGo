package t337

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val:  3,
			Left: nil,
			Right: &TreeNode{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(rob(&root))
}
