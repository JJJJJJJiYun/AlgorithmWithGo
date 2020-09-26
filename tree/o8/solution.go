package o8

import "LeetCode/utils"

// 中序遍历的下一个节点

func GetNextNodeOfInOrderTraversal(node *utils.TreeNodeWithParent) *utils.TreeNodeWithParent {
	if node == nil {
		return nil
	}
	if node.Right != nil {
		temp := node.Right
		for temp.Left != nil {
			temp = temp.Left
		}
		return temp
	} else {
		for node.Parent != nil && node.Parent.Right == node {
			node = node.Parent
		}
		return node.Parent
	}
}
