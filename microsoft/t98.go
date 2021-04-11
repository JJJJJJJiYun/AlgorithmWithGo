package microsoft

import "math"

func isValidBST(root *TreeNode) bool {
	pre := math.MinInt64
	return travel(root, &pre)
}

func travel(root *TreeNode, pre *int) bool {
	if root == nil {
		return true
	}
	b1 := travel(root.Left, pre)
	if root.Val <= *pre {
		return false
	}
	*pre = root.Val
	b2 := travel(root.Right, pre)
	return b1 && b2
}
