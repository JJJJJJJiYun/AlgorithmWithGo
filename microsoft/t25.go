package microsoft

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for {
		start := pre.Next
		end := start
		for i := 0; end != nil && i < k-1; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}
		next := end.Next
		end.Next = nil
		reverseListNode(start)
		pre.Next = end
		start.Next = next
		pre = start
	}
	return dummy.Next
}

func reverseListNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p, q := head, head.Next
	for q != nil {
		temp := q.Next
		q.Next = p
		p = q
		q = temp
	}
	head.Next = nil
	return p
}
