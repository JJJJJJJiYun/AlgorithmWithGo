package t222

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftLvl := countLevel(root.Left)
	rightLvl := countLevel(root.Right)
	if leftLvl == rightLvl {
		return countNodes(root.Right) + 1<<leftLvl
	} else {
		return countNodes(root.Left) + 1<<rightLvl
	}
}

func countLevel(root *TreeNode) int {
	lvl := 0
	for root != nil {
		root = root.Left
		lvl++
	}
	return lvl
}
