package tree

import (
	"LeetCode/utils"
	"fmt"
)

type Node struct {
	val    int
	time   int
	height int
	left   *Node
	right  *Node
}

// Insert 插入
func Insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{
			val: val,
		}
	}
	if root.val > val {
		root.left = Insert(root.left, val)
	} else if root.val == val {
		root.time++
	} else {
		root.right = Insert(root.right, val)
	}
	return balanceTree(root)
}

// balanceTree 针对不同情况进行平衡操作
func balanceTree(root *Node) *Node {
	if utils.Abs(diffHeightL2R(root)) < 2 {
		updateHeight(root)
	} else if diffHeightL2R(root) == 2 { // 左偏
		if diffHeightL2R(root.left) == -1 { // LR
			root.left = rotateR2L(root.left) // 先左旋
			root = rotateL2R(root)           // 再右旋
		} else { // LL
			root = rotateL2R(root) // 右旋
		}
	} else if diffHeightL2R(root) == -2 { // 右偏
		if diffHeightL2R(root.right) == 1 { // RL
			root.right = rotateL2R(root.right) // 先右旋
			root = rotateR2L(root)             // 再左旋
		} else { // RR
			root = rotateR2L(root) // 左旋
		}
	}
	return root
}

// diffHeightL2R 左右子树高度差
func diffHeightL2R(root *Node) int {
	if root.left != nil && root.right != nil {
		return root.left.height - root.right.height
	} else if root.left != nil {
		return root.left.height + 1
	} else if root.right != nil {
		return -(root.right.height + 1)
	} else {
		return 0
	}
}

// rotateL2R 右旋
func rotateL2R(root *Node) *Node {
	leftChild := root.left
	leftChild.right, root.left = root, leftChild.right
	// 注意：旋转实际上只改变了root和leftChild的高度，其他高度是不变的
	// 注意：必须先更新原来的root高度，因为他处于下层
	updateHeight(root)
	updateHeight(leftChild)
	return leftChild
}

// rotateR2L 左旋
func rotateR2L(root *Node) *Node {
	rightChild := root.right
	rightChild.left, root.right = root, rightChild.left
	updateHeight(root)
	updateHeight(rightChild)
	return rightChild
}

// updateHeight 更新节点高度
func updateHeight(root *Node) {
	if root.left != nil && root.right != nil {
		root.height = utils.Max(root.left.height, root.right.height) + 1
	} else if root.left != nil {
		root.height = root.left.height + 1
	} else if root.right != nil {
		root.height = root.right.height + 1
	} else {
		root.height = 0
	}
}

func HierarchicalTraversalWithBreak(root *Node) {
	curs := make([]*Node, 0)
	nexts := make([]*Node, 0)
	curs = append(curs, root)
	for len(curs) > 0 {
		temp := curs[0]
		curs = curs[1:]
		fmt.Printf("%v ", temp.val)
		if temp.left != nil {
			nexts = append(nexts, temp.left)
		}
		if temp.right != nil {
			nexts = append(nexts, temp.right)
		}
		if len(curs) == 0 {
			curs = nexts
			nexts = make([]*Node, 0)
			fmt.Println()
		}
	}
}
