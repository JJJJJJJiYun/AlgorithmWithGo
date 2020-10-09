package t236

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	fathers  []*TreeNode
	pFathers []*TreeNode
	qFathers []*TreeNode
)

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, p, q)
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	i := 0
	helper(root, p, q)
	for i < len(pFathers) && i < len(qFathers) && pFathers[i] == qFathers[i] {
		i++
	}
	i--
	return pFathers[i]
}

func helper(root, p, q *TreeNode) {
	fathers = append(fathers, root)
	if root == p {
		pFathers = append(pFathers, fathers...)
	}
	if root == q {
		qFathers = append(qFathers, fathers...)
	}
	if root.Left != nil {
		helper(root.Left, p, q)
	}
	if root.Right != nil {
		helper(root.Right, p, q)
	}
	fathers = fathers[:len(fathers)-1]
}
