package dictation

type Node struct {
	key, value string
	pre, next  *Node
}

type LRUCache struct {
	cap        int
	m          map[string]*Node
	head, tail *Node
}

func NewLRUCache(cap int) *LRUCache {
	Cache := LRUCache{
		cap: cap,
		m:   make(map[string]*Node),
	}
	Cache.head, Cache.tail = &Node{}, &Node{}
	Cache.head.next, Cache.tail.pre = Cache.tail, Cache.head
	return &Cache
}

func (lru *LRUCache) insert(node *Node) {
	node.pre, node.next = lru.head, lru.head.next
}

func (lru *LRUCache) remove(node *Node) {
	node.pre.next, node.next.pre = node.next, node.pre
}

func (lru *LRUCache) Get(key string) string {
	if node, ok := lru.m[key]; ok {
		lru.remove(node)
		lru.insert(node)
		return node.value
	} else {
		return ""
	}
}

func (lru *LRUCache) Put(key, value string) {
	if node, ok := lru.m[key]; ok {
		lru.remove(node)
		node.value = value
		lru.insert(node)
	} else {
		newNode := Node{
			key:   key,
			value: value,
		}
		lru.m[key] = &newNode
		lru.insert(&newNode)
		if len(lru.m) > lru.cap {
			lru.remove(lru.tail.pre)
		}
	}
}
