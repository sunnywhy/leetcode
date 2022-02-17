package main

// LRU Cache
type LRUCache struct {
	data     map[int]*Node
	head     *Node
	tail     *Node
	capacity int
	size     int
}

type Node struct {
	key  int
	val  int
	prev *Node
	next *Node
}

func Constructor(capacity int) LRUCache {
	m := make(map[int]*Node)
	head, tail := &Node{}, &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		data:     m,
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if e, ok := this.data[key]; ok {
		this.MoveToFront(e)
		return e.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.data[key]; ok {
		e.val = value
		this.MoveToFront(e)
	} else {
		e := &Node{key: key, val: value}
		this.PushFront(e)
		this.data[key] = e
		this.size++
	}

	if this.size > this.capacity {
		key := this.RemoveBack()
		delete(this.data, key)
		this.size--
	}
}

func (this *LRUCache) MoveToFront(node *Node) {
	this.Delete(node)
	this.PushFront(node)
}

func (this *LRUCache) RemoveBack() int {
	node := this.tail.prev
	this.Delete(node)
	return node.key
}

func (this *LRUCache) Delete(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next
	node.prev = nil
	node.next = nil
}

func (this *LRUCache) PushFront(node *Node) {
	this.head.next.prev = node
	node.next = this.head.next
	this.head.next = node
	node.prev = this.head
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
