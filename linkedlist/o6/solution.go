package o6

import (
	"LeetCode/utils"
	"fmt"
)

// 反转打印链表

func ReversePrintLinkedList(node *utils.LinkedNode) {
	stack := make([]int, 0)
	for node != nil {
		stack = append(stack, node.Val)
		node = node.Next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		fmt.Printf("%d ", stack[i])
	}
}
