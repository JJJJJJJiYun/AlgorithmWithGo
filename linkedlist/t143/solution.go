package t143

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1 2 3 4 5
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	newHead := reverseList(slow.Next)
	slow.Next = nil
	for newHead != nil {
		temp := newHead.Next
		newHead.Next = head.Next
		head.Next = newHead
		head = newHead.Next
		newHead = temp
	}
}

func reverseList(head *ListNode) *ListNode {
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
