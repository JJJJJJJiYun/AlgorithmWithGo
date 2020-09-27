package o22

import "LeetCode/utils"

// 链表的倒数第k个节点

func LastKNode(first *utils.LinkedNode, k int) *utils.LinkedNode {
	if first == nil || k <= 0 {
		return nil
	}
	q := first
	for i := 1; i < k; i++ {
		if q.Next == nil {
			return nil
		}
		q = q.Next
	}
	p := first
	for q.Next != nil {
		p = p.Next
		q = q.Next
	}
	return p
}
