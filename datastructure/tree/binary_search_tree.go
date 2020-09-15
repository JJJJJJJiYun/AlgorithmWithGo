package tree

type Node struct {
	val   int
	time  int
	left  *Node
	right *Node
}

// Insert 插入
func Insert(root *Node, val int) *Node {
	if root == nil {
		return &Node{
			val:  val,
			time: 1,
		}
	}
	if root.val > val {
		root.left = Insert(root.left, val)
	} else if root.val == val {
		root.time++
	} else {
		root.right = Insert(root.right, val)
	}
	return root
}

// Delete 删除
func Delete(root *Node, val int) *Node {
	if root == nil {
		return root
	}
	if root.val > val {
		root.left = Delete(root.left, val)
	} else if root.val < val {
		root.right = Delete(root.right, val)
	} else {
		if root.time > 1 {
			root.time--
		} else {
			if root.left != nil && root.right != nil {
				temp := root.left
				for temp.right != nil {
					temp = temp.right
				}
				root = Delete(root, temp.val)
				root.val = temp.val
			} else {
				if root.left != nil {
					root = root.left
				} else {
					root = root.right
				}
			}
		}
	}
	return root
}
