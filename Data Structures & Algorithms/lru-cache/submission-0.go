type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}

	head.next = tail
	tail.prev = head

	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

// Remove node from DLL
func (this *LRUCache) remove(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// Insert node right after head (MRU position)
func (this *LRUCache) insert(node *Node) {
	node.next = this.head.next
	node.prev = this.head

	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}

	// Move to front (mark as recently used)
	this.remove(node)
	this.insert(node)

	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		// Update existing key
		node.value = value
		this.remove(node)
		this.insert(node)
		return
	}

	// Evict LRU if capacity full
	if len(this.cache) == this.capacity {
		lru := this.tail.prev
		this.remove(lru)
		delete(this.cache, lru.key)
	}

	// Insert new node
	newNode := &Node{key: key, value: value}
	this.cache[key] = newNode
	this.insert(newNode)
}