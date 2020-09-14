package tree

type Node struct {
	val   int
	time  int
	left  *Node
	right *Node
}

// Insert 插入（非递归）
func Insert(root, node *Node) *Node {
	if root == nil {
		root = node
		return root
	}
	temp := root
	for true {
		if root.val > node.val {
			if root.left == nil {
				root.left = node
				break
			}
			root = root.left
		} else if root.val == node.val {
			root.time++
			break
		} else {
			if root.right == nil {
				root.right = node
				break
			}
			root = root.right
		}
	}
	return temp
}

// InsertWithRecursion 插入（递归）
func InsertWithRecursion(root, node *Node) *Node {
	if root == nil {
		root = node
		return root
	}
	if root.val > node.val {
		root.left = InsertWithRecursion(root.left, node)
	} else if root.val == node.val {
		root.time++
	} else {
		root.right = InsertWithRecursion(root.right, node)
	}
	return root
}
