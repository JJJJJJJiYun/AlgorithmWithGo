package t111

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 给定一个二叉树，找出其最小深度。
// 最小深度是从根节点到最近叶子节点的最短路径上的节点数量。

// BFS
func minDepth(root *TreeNode) int {
	// 注意这里的判空
	if root == nil {
		return 0
	}
	roots := make([]*TreeNode, 0)
	roots = append(roots, root)
	depth := 1
	for len(roots) != 0 {
		rootsLen := len(roots)
		for i := 0; i < rootsLen; i++ {
			cur := roots[0]
			roots = roots[1:]
			if cur.Left == nil && cur.Right == nil {
				return depth
			}
			if cur.Left != nil {
				roots = append(roots, cur.Left)
			}
			if cur.Right != nil {
				roots = append(roots, cur.Right)
			}
		}
		depth++
	}
	return 0
}
