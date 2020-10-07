package o55

import "LeetCode/utils"

func DepthOfTree(root *utils.TreeNode) int {
	if root == nil {
		return 0
	}
	left := DepthOfTree(root.Left)
	right := DepthOfTree(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
