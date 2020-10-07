package o54

import (
	"LeetCode/utils"
	"fmt"
)

// 第k大的节点

func KthNode(root *utils.TreeNode, k int) int {
	stack := make([]*utils.TreeNode, 0)
	cnt := 0
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			cnt++
			if cnt == k {
				return root.Val
			}
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return -1
}

// 二叉树中序遍历（非递归）
func PreOrderTraversal(root *utils.TreeNode) {
	stack := make([]*utils.TreeNode, 0)
	for root != nil || len(stack) != 0 {
		if root != nil {
			// do something
			fmt.Println(root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
}

// 二叉树前序遍历（非递归）
func InOrderTraversal(root *utils.TreeNode) {
	stack := make([]*utils.TreeNode, 0)
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			// do something
			fmt.Println(root.Val)
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
}

// 二叉树后序遍历（非递归）
func PostOrderTraversal(root *utils.TreeNode) {
	stack := make([]*utils.TreeNode, 0)
	helper := make([]*utils.TreeNode, 0)
	for root != nil || len(stack) != 0 {
		if root != nil {
			// do something
			helper = append(helper, root)
			stack = append(stack, root)
			root = root.Right
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Left
		}
	}
	for i := len(helper) - 1; i >= 0; i-- {
		fmt.Println(helper[i].Val)
	}
}
