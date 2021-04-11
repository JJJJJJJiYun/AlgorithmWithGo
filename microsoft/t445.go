package microsoft

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1, stack2 []*ListNode
	for l1 != nil {
		stack1 = append([]*ListNode{l1}, stack1...)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append([]*ListNode{l2}, stack2...)
		l2 = l2.Next
	}
	var i, j, carry int
	var res *ListNode
	for i < len(stack1) || j < len(stack2) || carry != 0 {
		var num1, num2 int
		if i < len(stack1) {
			num1 = stack1[i].Val
			i++
		}
		if j < len(stack2) {
			num2 = stack2[j].Val
			j++
		}
		node := &ListNode{Val: (num1 + num2 + carry) % 10}
		carry = (num1 + num2 + carry) / 10
		node.Next = res
		res = node
	}
	return res
}
