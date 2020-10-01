package o34

import (
	"LeetCode/utils"
	"fmt"
)

func TreePathOfNum(root *utils.TreeNode, target int) {
	if root == nil {
		return
	}
	nums := make([]int, 0)
	helper(root, nums, 0, target)
}

func helper(root *utils.TreeNode, nums []int, current, target int) {
	current += root.Val
	nums = append(nums, root.Val)
	if current == target {
		fmt.Println(nums)
		return
	}
	if root.Left != nil {
		helper(root.Left, nums, current, target)
	}
	if root.Right != nil {
		helper(root.Right, nums, current, target)
	}
}
