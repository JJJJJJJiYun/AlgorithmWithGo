package microsoft

// 这个left和right，记录的是当前节点的子树中是否包含了要求的节点
func lowestCommonAncestorOfBinaryTree(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lowestCommonAncestorOfBinaryTree(root.Left, p, q)
	right := lowestCommonAncestorOfBinaryTree(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}

func lowestCommonAncestorOfBinaryTree2(root, p, q *TreeNode) *TreeNode {
	var father *TreeNode
	fatherMap := make(map[*TreeNode]*TreeNode)
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		fatherMap[root] = father
		father = root
		dfs(root.Left)
		father = root
		dfs(root.Right)
	}
	dfs(root)
	pFathers := make(map[*TreeNode]bool)
	for p != nil {
		pFathers[p] = true
		p = fatherMap[p]
	}
	for q != nil {
		if pFathers[q] {
			return q
		}
		q = fatherMap[q]
	}
	return nil
}
