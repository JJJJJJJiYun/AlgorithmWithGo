package microsoft

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

type LRUCache struct {
	size, capacity int
	cache          map[int]*DLinkedNode
	head, tail     *DLinkedNode
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    map[int]*DLinkedNode{},
		head:     &DLinkedNode{},
		tail:     &DLinkedNode{},
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRUCache) Get(key int) int {
	node, ok := l.cache[key]
	// 如果不存在关键字，返回-1
	if !ok {
		return -1
	}
	// 存在关键字，将关键字移到队首
	l.moveToHead(node)
	return node.value
}

func (l *LRUCache) moveToHead(node *DLinkedNode) {
	l.removeNode(node)
	l.addToHead(node)
}

func (l *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (l *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = l.head
	node.next = l.head.next
	l.head.next.prev = node
	l.head.next = node
}

func (l *LRUCache) Put(key int, value int) {
	if n, ok := l.cache[key]; !ok {
		node := &DLinkedNode{
			key:   key,
			value: value,
		}
		l.cache[key] = node
		l.addToHead(node)
		l.size++
		if l.size > l.capacity {
			removed := l.removeTail()
			delete(l.cache, removed.key)
			l.size--
		}
	} else {
		n.value = value
		l.moveToHead(n)
	}
}

func (l *LRUCache) removeTail() *DLinkedNode {
	node := l.tail.prev
	l.removeNode(node)
	return node
}
