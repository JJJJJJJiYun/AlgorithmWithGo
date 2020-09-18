package bst

import (
	"testing"
)

func TestRecursivePreOrderTraversal(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	RecursivePreOrderTraversal(root)
}

func TestNonRecursivePreOrderTraversal(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	NonRecursivePreOrderTraversal(root)
}

func TestRecursiveInOrderTraversal(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	RecursiveInOrderTraversal(root)
}

func TestNonRecursiveInOrderTraversal(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	NonRecursiveInOrderTraversal(root)
}

func TestRecursivePostOrderTraversal(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	RecursivePostOrderTraversal(root)
}

func TestNonRecursivePostOrderTraversal(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	NonRecursivePostOrderTraversal(root)
}

func TestNonRecursivePostOrderTraversalSimple(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	NonRecursivePostOrderTraversalSimple(root)
}

func TestHierarchicalTraversal(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	HierarchicalTraversal(root)
}

func TestHierarchicalTraversalWithBreak(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	HierarchicalTraversalWithBreak(root)
}

func TestHierarchicalTraversalWithBreakAndReverse(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	Insert(root, 10)
	Insert(root, 13)
	Insert(root, 6)
	Insert(root, 18)
	Insert(root, 23)
	HierarchicalTraversalWithBreakAndReverse(root)
}
