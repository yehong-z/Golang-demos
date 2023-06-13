package lru

type Node struct {
	Key, Value int
	Prev, Next *Node
}

type LRUCache struct {
	Capacity int           // 缓存容量
	Map      map[int]*Node // 哈希表，存储缓存数据
	Head     *Node         // 双向链表头部节点
	Tail     *Node         // 双向链表尾部节点
}

func LRUConstructor(capacity int) LRUCache {
	l := LRUCache{
		Capacity: capacity,
		Map:      make(map[int]*Node)}
	l.Head, l.Tail = &Node{}, &Node{}
	l.Head.Next, l.Tail.Prev = l.Tail, l.Head
	return l
}

func (l *LRUCache) remove(node *Node) {
	node.Prev.Next, node.Next.Prev = node.Next, node.Prev
}

func (l *LRUCache) insert(node *Node) {
	node.Next, node.Prev = l.Head.Next, l.Head
	l.Head.Next, node.Next.Prev = node, node
}

func (l *LRUCache) Get(key int) int {
	if node, ok := l.Map[key]; ok {
		l.remove(node) // 从链表中删除当前节点
		l.insert(node) // 将节点放到链表头部
		return node.Value
	}
	return -1
}

func (l *LRUCache) Put(key, value int) {
	if node, ok := l.Map[key]; ok {
		l.remove(node)     // 从链表中删除当前节点
		node.Value = value // 更新节点的值
		l.insert(node)     // 将节点放到链表头部
	} else {
		node := &Node{Key: key, Value: value}
		l.Map[key] = node // 添加到哈希表中
		l.insert(node)    // 将节点放到链表头部
		if len(l.Map) > l.Capacity {
			delete(l.Map, l.Tail.Prev.Key) // 删除链表尾部的元素
			l.remove(l.Tail.Prev)
		}
	}
}
