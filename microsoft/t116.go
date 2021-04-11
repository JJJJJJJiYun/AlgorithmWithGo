package microsoft

func connect(root *NodeWithNext) *NodeWithNext {
	travelConnect(root)
	return root
}

func travelConnect(root *NodeWithNext) {
	if root == nil {
		return
	}
	if root.Left != nil {
		root.Left.Next = root.Right
	}
	if root.Next != nil && root.Right != nil {
		root.Right.Next = root.Next.Left
	}
	travelConnect(root.Left)
	travelConnect(root.Right)
}
