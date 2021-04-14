package microsoft

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for pre.Next != nil && pre.Next.Next != nil {
		node1, node2 := pre.Next, pre.Next.Next
		pre.Next = node2
		node1.Next = node2.Next
		node2.Next = node1
		pre = node1
	}
	return dummy.Next
}
