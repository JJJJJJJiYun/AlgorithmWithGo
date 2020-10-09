package o28

import "AlgorithWithGo/utils"

// 是否是镜像树

func IsMirrorTree(root *utils.TreeNode) bool {
	var helper func(left, right *utils.TreeNode) bool
	helper = func(left, right *utils.TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left != nil && right != nil && left.Val == right.Val {
			return helper(left.Left, right.Right) && helper(left.Right, right.Left)
		}
		return false
	}
	return helper(root, root)
}
