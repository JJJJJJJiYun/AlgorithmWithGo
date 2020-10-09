package o36

import "AlgorithWithGo/utils"

func ConvertTreeToLinkedList(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}
	var last *utils.TreeNode
	var helper func(root *utils.TreeNode)
	helper = func(root *utils.TreeNode) {
		if root.Left != nil {
			helper(root.Left)
		}
		root.Left = last
		if last != nil {
			last.Right = root
		}
		last = root
		if root.Right != nil {
			helper(root.Right)
		}
	}
	helper(root)
	for root.Left != nil {
		root = root.Left
	}
	return root
}
