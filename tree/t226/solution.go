package t226

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left = right
	root.Right = left
	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	newRoot := &TreeNode{Val: root.Val}
	helper(root, newRoot)
	return newRoot
}

// helper helper
func helper(root, newRoot *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		newRoot.Right = &TreeNode{Val: root.Left.Val}
		helper(root.Left, newRoot.Right)
	}
	if root.Right != nil {
		newRoot.Left = &TreeNode{Val: root.Right.Val}
		helper(root.Right, newRoot.Left)
	}
}
