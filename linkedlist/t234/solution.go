package t234

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	// 找到链表中点
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 翻转链表的后半部分
	newHead := reverseList(slow.Next)
	slow.Next = nil
	for newHead != nil && head.Val == newHead.Val {
		head = head.Next
		newHead = newHead.Next
	}
	if newHead != nil {
		return false
	}
	return true
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
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
