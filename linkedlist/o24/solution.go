package o24

import "LeetCode/utils"

func ReverseLinkedList(first *utils.LinkedNode) *utils.LinkedNode {
	if first == nil || first.Next == nil {
		return first
	}
	p, q := first, first.Next
	p.Next = nil
	for q != nil {
		temp := q.Next
		q.Next = p
		p = q
		q = temp
	}
	return p
}
