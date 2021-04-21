package LC_Go

type ListNode struct {
	key, value int
	prev, next *ListNode
}

type LRUCache struct {
	size, capacity int
	head, tail     *ListNode
	cacheMap       map[int]*ListNode
}

func LruConstructor(capacity int) LRUCache {
	return LRUCache{
		size:     0,
		capacity: capacity,
		head:     nil,
		tail:     nil,
		cacheMap: make(map[int]*ListNode),
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cacheMap[key]
	if !ok {
		return -1
	}
	this.moveToBegin(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.cacheMap[key]
	if ok {
		node.value = value
		this.moveToBegin(node)
		return
	}
	// key doesn't exist
	cur := &ListNode{
		key:   key,
		value: value,
	}
	this.cacheMap[key] = cur
	// two things:
	// 1. prepend cur to ListNode
	// 2. check size
	// case 1: if it's an empty ListNode
	if this.head == nil {
		this.head = cur
		this.tail = cur
		this.size++
		return
	}
	// case 2: it's NOT an empty ListNode
	cur.next = this.head
	this.head.prev = cur
	this.head = cur
	if this.size < this.capacity {
		this.size++
	} else {
		delete(this.cacheMap, this.tail.key)
		this.tail = this.tail.prev
		this.tail.next = nil
	}
}

func (this *LRUCache) moveToBegin(node *ListNode) {
	if node == nil || node == this.head {
		return
	}
	// node is NOT the 1st node
	if node == this.tail {
		this.tail = this.tail.prev
		this.tail.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	node.next = this.head
	this.head.prev = node
	node.prev = nil
	this.head = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
