package o27

import "AlgorithWithGo/utils"

// 得到镜像树

func MirrorTree(root *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}
	newRoot := &utils.TreeNode{Val: root.Val}
	if root.Left != nil {
		newRoot.Right = MirrorTree(root.Left)
	}
	if root.Right != nil {
		newRoot.Left = MirrorTree(root.Right)
	}
	return newRoot
}
