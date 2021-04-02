package microsoft

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		if head.Next == nil {
			head.Next = pre
			break
		}
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return head
}
