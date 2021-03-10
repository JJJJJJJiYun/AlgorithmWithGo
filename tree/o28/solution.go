package o28

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isSymmetricHelper(root, root)
}

func isSymmetricHelper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return false
	}
	if left != nil && right != nil && left.Val == right.Val {
		return isSymmetricHelper(left.Left, right.Right) && isSymmetricHelper(left.Right, right.Left)
	}
	return false
}
