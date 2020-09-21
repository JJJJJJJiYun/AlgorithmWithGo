package tree

import "testing"

func TestInsert(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	HierarchicalTraversalWithBreak(root)
}

func TestDelete(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	Insert(root, 7)
	Insert(root, 13)
	Delete(root, 5)
	HierarchicalTraversalWithBreak(root)
}
