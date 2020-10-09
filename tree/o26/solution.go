package o26

import (
	"AlgorithWithGo/utils"
)

// 树的子结构

func PartOfTree(part, root *utils.TreeNode) bool {
	var helper func(part, root *utils.TreeNode) bool
	helper = func(part, root *utils.TreeNode) bool {
		if root == nil && part == nil {
			return true
		}
		if (root == nil && part != nil) || (root != nil && part == nil) || part.Val != root.Val {
			return false
		}
		return helper(part.Left, root.Left) && helper(part.Right, root.Right)
	}
	if helper(part, root) {
		return true
	}
	if root.Left != nil {
		if PartOfTree(part, root.Left) {
			return true
		}
	}
	if root.Right != nil {
		if PartOfTree(part, root.Right) {
			return true
		}
	}
	return false
}
