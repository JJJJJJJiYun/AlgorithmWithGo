package microsoft

type NodeWithRandom struct {
	Val    int
	Next   *NodeWithRandom
	Random *NodeWithRandom
}

type NodeWithNext struct {
	Val   int
	Left  *NodeWithNext
	Right *NodeWithNext
	Next  *NodeWithNext
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
