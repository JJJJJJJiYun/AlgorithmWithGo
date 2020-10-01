package o35

import "LeetCode/utils"

func CopyComplexLinkedList(first *utils.ComplexLinkedNode) *utils.ComplexLinkedNode {
	p := first
	for p != nil {
		q := p.Next
		tmp := &utils.ComplexLinkedNode{
			Val:  p.Val,
			Next: q,
		}
		p.Next = tmp
		p = q
	}
	p = first
	q := first.Next
	res := first.Next
	for p != nil {
		if p.Sibling != nil {
			q.Sibling = p.Sibling.Next
		}
		p = q.Next
		if p == nil {
			break
		}
		q.Next = q.Next.Next
		q = q.Next
	}
	return res
}
