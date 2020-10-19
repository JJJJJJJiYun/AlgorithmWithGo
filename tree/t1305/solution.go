package t1305

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 对两个排好序的数组排序，用归并！
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	order1 := inOrderTraversal(root1)
	order2 := inOrderTraversal(root2)
	res := make([]int, 0)
	i, j := 0, 0
	for i < len(order1) && j < len(order2) {
		if order1[i] < order2[j] {
			res = append(res, order1[i])
			i++
		} else {
			res = append(res, order2[j])
			j++
		}
	}
	if i < len(order1) {
		res = append(res, order1[i:]...)
	}
	if j < len(order2) {
		res = append(res, order2[j:]...)
	}
	return res
}

func inOrderTraversal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for root != nil || len(stack) != 0 {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			res = append(res, root.Val)
			stack = stack[:len(stack)-1]
			root = root.Right
		}
	}
	return res
}
