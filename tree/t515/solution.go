package t515

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	cur := []*TreeNode{root}
	next := make([]*TreeNode, 0)
	res := make([]int, 0)
	max := math.MinInt32
	for len(cur) != 0 {
		temp := cur[0]
		cur = cur[1:]
		if temp.Left != nil {
			next = append(next, temp.Left)
		}
		if temp.Right != nil {
			next = append(next, temp.Right)
		}
		if temp.Val > max {
			max = temp.Val
		}
		if len(cur) == 0 {
			cur = next
			next = make([]*TreeNode, 0)
			res = append(res, max)
			max = math.MinInt32
		}
	}
	return res
}
