package t993

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	recorder := make(map[int]struct {
		depth  int
		parent *TreeNode
	})
	preOrder(root, nil, x, y, 0, recorder)
	if recorder[x].depth == recorder[y].depth && recorder[x].parent != recorder[y].parent {
		return true
	}
	return false
}

func preOrder(root, parent *TreeNode, x, y, depth int, recorder map[int]struct {
	depth  int
	parent *TreeNode
}) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		recorder[root.Val] = struct {
			depth  int
			parent *TreeNode
		}{depth: depth, parent: parent}
	}
	preOrder(root.Left, root, x, y, depth+1, recorder)
	preOrder(root.Right, root, x, y, depth+1, recorder)
}
