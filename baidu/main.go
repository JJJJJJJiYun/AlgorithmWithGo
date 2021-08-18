package baidu

import (
	"AlgorithWithGo/utils"
	"fmt"
)

// getIndex getIndex
func getIndex(nums []int, left, right int) int {
	pivot := nums[left]
	for left < right {
		for left < right && nums[right] >= pivot {
			right--
		}
		nums[left] = nums[right]
		for left < right && nums[left] <= pivot {
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = pivot
	return left
}

// merge merge
func merge(nums, helper []int, left, mid, right int) {
	i, j, k := left, mid+1, left
	for i <= mid && j <= right {
		if nums[i] < nums[j] {
			helper[k] = nums[i]
			i++
		} else {
			helper[k] = nums[j]
			j++
		}
		k++
	}
	for i <= mid {
		helper[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		helper[k] = nums[j]
		j++
		k++
	}
	for left <= right {
		nums[left] = helper[left]
		left++
	}
}

// inorder inorder
func inorder(root *utils.TreeNode) {
	stack := make([]*utils.TreeNode, 0)
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Println(root.Val)
			root = root.Right
		}
	}
}
