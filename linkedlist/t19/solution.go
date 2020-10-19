package t19

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1->2->3->4->5
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 创建一个临时节点，防止删除的是头节点引发问题
	temp := &ListNode{Next: head}
	// 找到倒数第n+1个节点，并删除第n个
	p := temp
	for i := 1; i < n+1; i++ {
		p = p.Next
	}
	q := temp
	for p.Next != nil {
		p = p.Next
		q = q.Next
	}
	q.Next = q.Next.Next
	return temp.Next
}
