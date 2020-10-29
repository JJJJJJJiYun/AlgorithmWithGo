package t129

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, preSum int) int {
	if root == nil {
		return 0
	}
	sum := preSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	} else {
		return helper(root.Left, sum) + helper(root.Right, sum)
	}
}
