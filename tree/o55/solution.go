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

func IsBalanceTree(root *utils.TreeNode) bool {
	res, _ := isBalanceTree(root)
	return res
}

func isBalanceTree(root *utils.TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	l, ld := isBalanceTree(root.Left)
	r, rd := isBalanceTree(root.Right)
	diff := ld - rd
	if !l || !r || (diff > 1 || diff < -1) {
		return false, 0
	}
	if ld > rd {
		return true, 1 + ld
	}
	return true, 1 + rd
}
