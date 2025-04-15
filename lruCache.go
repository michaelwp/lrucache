package lrucache

type LRUCache interface {
	Get(key string) any
	Put(key string, value any)
	ViewList() []*data
}

type data struct {
	Key   string
	Value any
}

type node struct {
	key        string
	value      any
	prev, next *node
}

type lruCache struct {
	capacity int
	cache    map[string]*node
	head     *node
	tail     *node
}

func NewLRUCache(capacity int) LRUCache {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head
	return &lruCache{
		capacity: capacity,
		cache:    make(map[string]*node),
		head:     head,
		tail:     tail,
	}
}

func (c *lruCache) remove(node *node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *lruCache) insert(node *node) {
	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

func (c *lruCache) Get(key string) any {
	if node, exists := c.cache[key]; exists {
		c.remove(node)
		c.insert(node)
		return node.value
	}
	return nil
}

func (c *lruCache) Put(key string, value any) {
	if node, exists := c.cache[key]; exists {
		c.remove(node)
	}

	newNode := &node{key: key, value: value}
	c.insert(newNode)
	c.cache[key] = newNode

	if len(c.cache) > c.capacity {
		lru := c.tail.prev
		c.remove(lru)
		delete(c.cache, lru.key)
	}
}

func (c *lruCache) ViewList() []*data {
	currentNode := c.head
	list := make([]*data, c.capacity)
	index := 0

	for currentNode.next != c.tail {
		list[index] = &data{
			Key:   currentNode.next.key,
			Value: currentNode.next.value,
		}
		currentNode = currentNode.next
		index++
	}

	return list
}
