package tree

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	fmt.Println(root)
}

func TestDelete(t *testing.T) {
	root := Insert(nil, 15)
	Insert(root, 5)
	Insert(root, 16)
	Insert(root, 3)
	Insert(root, 12)
	Insert(root, 20)
	Insert(root, 16)
	Insert(root, 4)
	Delete(root, 5)
	fmt.Println(root)
}
