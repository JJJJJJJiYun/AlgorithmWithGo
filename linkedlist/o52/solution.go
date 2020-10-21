package o52

import "AlgorithWithGo/utils"

func FirstDuplicateNode(node1, node2 *utils.LinkedNode) *utils.LinkedNode {
	p := node1
	count1, count2 := 0, 0
	for p != nil {
		count1++
		p = p.Next
	}
	p = node2
	for p != nil {
		count2++
		p = p.Next
	}
	if count1 > count2 {
		for i := 0; i < count1-count2; i++ {
			node1 = node1.Next
		}
		for {
			if node1 == node2 {
				return node1
			}
			node1 = node1.Next
			node2 = node2.Next
		}
	} else {
		for i := 0; i < count2-count1; i++ {
			node2 = node2.Next
		}
		for {
			if node1 == node2 {
				return node2
			}
			node1 = node1.Next
			node2 = node2.Next
		}
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1->2->3
// 4->5->6
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = headB
		}
		if q != nil {
			q = q.Next
		} else {
			q = headA
		}
	}
	return p
}
