package o25

import "AlgorithWithGo/utils"

func MergeTwoSortedLinkedList(first1, first2 *utils.LinkedNode) *utils.LinkedNode {
	head := &utils.LinkedNode{}
	p := head
	for first1 != nil && first2 != nil {
		if first1.Val < first2.Val {
			p.Next = first1
			first1 = first1.Next
		} else {
			p.Next = first2
			first2 = first2.Next
		}
		p = p.Next
	}
	for first1 != nil {
		p.Next = first1
		p = p.Next
		first1 = first1.Next
	}
	for first2 != nil {
		p.Next = first2
		p = p.Next
		first2 = first2.Next
	}
	return head.Next
}
