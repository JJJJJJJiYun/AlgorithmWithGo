package o6

import "fmt"

type Node struct {
	val  int
	next *Node
}

// 反转打印链表
func ReversePrintLinkedList(node *Node) {
	stack := make([]int, 0)
	for node != nil {
		stack = append(stack, node.val)
		node = node.next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Printf("%d ", stack[i])
	}
}
