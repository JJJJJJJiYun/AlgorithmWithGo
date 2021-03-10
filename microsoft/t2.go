package microsoft

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	flag := 0
	for l1 != nil || l2 != nil {
		num1, num2 := 0, 0
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}
		val := num1 + num2 + flag
		if val > 9 {
			val -= 10
			flag = 1
		} else {
			flag = 0
		}
		node := &ListNode{
			Val: val,
		}
		p.Next = node
		p = p.Next
	}
	if flag == 1 {
		p.Next = &ListNode{
			Val: 1,
		}
	}
	return dummy.Next
}
