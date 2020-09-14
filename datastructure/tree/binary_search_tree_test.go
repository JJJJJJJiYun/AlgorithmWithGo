package tree

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	node1 := &Node{
		val:  15,
		time: 1,
	}
	node2 := &Node{
		val:  5,
		time: 1,
	}
	node3 := &Node{
		val:  16,
		time: 1,
	}
	node4 := &Node{
		val:  3,
		time: 1,
	}
	node5 := &Node{
		val:  12,
		time: 1,
	}
	node6 := &Node{
		val:  20,
		time: 1,
	}
	node7 := &Node{
		val:  16,
		time: 1,
	}
	root := Insert(nil, node1)
	Insert(root, node2)
	Insert(root, node3)
	Insert(root, node4)
	Insert(root, node5)
	Insert(root, node6)
	Insert(root, node7)
	fmt.Println(root)
}

func TestInsertWithRecursion(t *testing.T) {
	node1 := &Node{
		val:  15,
		time: 1,
	}
	node2 := &Node{
		val:  5,
		time: 1,
	}
	node3 := &Node{
		val:  16,
		time: 1,
	}
	node4 := &Node{
		val:  3,
		time: 1,
	}
	node5 := &Node{
		val:  12,
		time: 1,
	}
	node6 := &Node{
		val:  20,
		time: 1,
	}
	node7 := &Node{
		val:  16,
		time: 1,
	}
	root := InsertWithRecursion(nil, node1)
	InsertWithRecursion(root, node2)
	InsertWithRecursion(root, node3)
	InsertWithRecursion(root, node4)
	InsertWithRecursion(root, node5)
	InsertWithRecursion(root, node6)
	InsertWithRecursion(root, node7)
	fmt.Println(root)
}
