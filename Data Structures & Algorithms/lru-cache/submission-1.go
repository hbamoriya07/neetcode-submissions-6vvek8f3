type Node struct {
    key int
    value int
    next, prev *Node
}

type LRUCache struct {
    capacity int
    cache map[int]*Node
    head, tail *Node
}

func Constructor(capacity int) *LRUCache {
    head := &Node{}
    tail := &Node{}

    head.next = tail
    tail.prev = head

    return &LRUCache{
        capacity: capacity,
        cache: make(map[int]*Node),
        head: head,
        tail: tail,
    }
}

// 4 - 1 - 2 - 3

func(this *LRUCache) remove(node *Node) {
    node.next.prev = node.prev
    node.prev.next = node.next
}

func(this *LRUCache) insert(node *Node) {
    oldFirst := this.head.next
    node.next  = oldFirst
    node.prev = this.head
    this.head.next = node
    oldFirst.prev = node
}
 
func (this *LRUCache) Get(key int) int {
    if node, ok := this.cache[key]; ok {
        this.remove(node)
        this.insert(node)

        return node.value
    }

    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, ok := this.cache[key]; ok {
		node.value = value
		this.remove(node)
		this.insert(node)

		return
	}

	if len(this.cache) == this.capacity {
		lru := this.tail.prev
		this.remove(lru)
		delete(this.cache, lru.key)
	}

	newNode := &Node{key: key, value: value}
	this.insert(newNode)

	this.cache[key] = newNode
}
