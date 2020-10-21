package utils

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeWithParent struct {
	Val    int
	Left   *TreeNodeWithParent
	Right  *TreeNodeWithParent
	Parent *TreeNodeWithParent
}

type LinkedNode struct {
	Val  int
	Next *LinkedNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type ComplexLinkedNode struct {
	Val     int
	Next    *ComplexLinkedNode
	Sibling *ComplexLinkedNode
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func HierarchicalTraversalWithBreak(root *TreeNode) {
	curs := make([]*TreeNode, 0)
	nexts := make([]*TreeNode, 0)
	curs = append(curs, root)
	for len(curs) > 0 {
		temp := curs[0]
		curs = curs[1:]
		fmt.Printf("%v ", temp.Val)
		if temp.Left != nil {
			nexts = append(nexts, temp.Left)
		}
		if temp.Right != nil {
			nexts = append(nexts, temp.Right)
		}
		if len(curs) == 0 {
			curs = nexts
			nexts = make([]*TreeNode, 0)
			fmt.Println()
		}
	}
}

func PrintLinkedList(first *LinkedNode) {
	for first != nil {
		fmt.Printf("%v ", first.Val)
		first = first.Next
	}
}

func PrintComplexLinkedList(first *ComplexLinkedNode) {
	for first != nil {
		fmt.Printf("%v ", first.Val)
		first = first.Next
	}
}
