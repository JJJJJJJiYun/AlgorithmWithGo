package microsoft

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := root.Val
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(dfs(node.Left), 0)
		right := max(dfs(node.Right), 0)
		this := left + right + node.Val
		res = max(res, this)
		return node.Val + max(left, right)
	}
	dfs(root)
	return res
}
