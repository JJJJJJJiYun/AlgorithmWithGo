package microsoft

func buildTree(preorder []int, inorder []int) *TreeNode {
	return buildHelper(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func buildHelper(preorder, inorder []int, ps, pe, is, ie int) *TreeNode {
	if ps > pe || is > ie {
		return nil
	}
	root := &TreeNode{Val: preorder[ps]}
	index := -1
	for i := is; i <= ie; i++ {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}
	root.Left = buildHelper(preorder, inorder, ps+1, ps+index-is, is, index-1)
	root.Right = buildHelper(preorder, inorder, ps+index-is+1, pe, index+1, ie)
	return root
}
