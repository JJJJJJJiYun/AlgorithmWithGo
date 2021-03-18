package microsoft

func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	flag := true
	cur := []*TreeNode{root}
	var next []*TreeNode
	for len(cur) != 0 {
		temp := make([]int, len(cur))
		for i, node := range cur {
			if flag {
				temp[i] = node.Val
			} else {
				temp[len(cur)-1-i] = node.Val
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		res = append(res, temp)
		cur = next
		next = make([]*TreeNode, 0)
		flag = !flag
	}
	return res
}
