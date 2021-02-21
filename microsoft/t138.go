package microsoft

// 创建一个拷贝节点
func copyRandomList(head *Node) *Node {
	// 如果给定链表为空，返回空
	if head == nil {
		return nil
	}
	// 对于每一个链表节点，在它后面创建一个它自身的拷贝
	p := head
	for p != nil {
		temp := &Node{
			Val:  p.Val,
			Next: p.Next,
		}
		p.Next = temp
		p = p.Next.Next
	}
	// 对于每一个拷贝节点，设置random节点
	p = head
	q := head.Next
	for {
		if p.Random != nil {
			q.Random = p.Random.Next
		}
		p = q.Next
		if p == nil {
			break
		}
		q = p.Next
	}
	// 重新连接拷贝节点并还原原链表
	newHead := head.Next
	p, q = head, head.Next
	for p != nil && q != nil {
		p.Next = q.Next
		if p.Next == nil {
			break
		}
		q.Next = p.Next.Next
		p = p.Next
		q = q.Next
	}
	return newHead
}
