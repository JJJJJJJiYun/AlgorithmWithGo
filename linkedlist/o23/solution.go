package o23

import (
	"AlgorithWithGo/utils"
)

// 链表中环的起始节点

func FindFirstNodeInLoop(first *utils.LinkedNode) *utils.LinkedNode {
	if first == nil || first.Next == nil {
		return nil
	}
	// 证明有环，一快一慢两个指针，一个每个走两步一个每个走一步，
	// 如果相遇则有环，如果快指针走到了nil，说明无环
	fast, slow := first.Next.Next, first.Next
	for fast != slow {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	// 计算环内元素个数n
	count := 1
	fast = fast.Next
	for fast != slow {
		fast = fast.Next
		count++
	}
	// 让快指针先走n步，然后同速走，相遇的点就是环的入口
	fast, slow = first, first
	for i := 0; i < count; i++ {
		fast = fast.Next
	}
	for fast != slow {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
