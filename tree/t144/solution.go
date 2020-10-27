package t144

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for root != nil || len(stack) != 0 {
		if root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}
