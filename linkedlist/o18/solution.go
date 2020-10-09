package o18

import "AlgorithWithGo/utils"

// 删除链表中的节点，时间复杂度O(1)

// 默认确保了node在链表中
func DeleteNode(first, node *utils.LinkedNode) *utils.LinkedNode {
	// 删除头节点
	if first == node {
		return first.Next
	}
	// 删除尾节点，这里不能直接将node赋成nil，需要从头找起
	// 和一些可以直接管理内存的语言不同
	if node.Next == nil {
		temp := first
		for temp.Next.Next != nil {
			temp = temp.Next
		}
		temp.Next = nil
		return first
	}
	node.Val = node.Next.Val
	node.Next = node.Next.Next
	return first
}

// 删除链表中的重复节点

func DeleteDuplicateNode(first *utils.LinkedNode) *utils.LinkedNode {
	head := &utils.LinkedNode{
		Val:  first.Val - 1,
		Next: first,
	}
	p, q := head, head.Next
	for q != nil {
		temp := q
		for temp.Next != nil && temp.Next.Val == temp.Val {
			temp = temp.Next
		}
		if temp != q {
			p.Next = temp.Next
			q = temp.Next
		} else {
			p = q
			q = temp.Next
		}
	}
	return head.Next
}
