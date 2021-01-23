package main

// Runtime: 152 ms, faster than 20.47% of Go online submissions for LRU Cache.
// Memory Usage: 13.4 MB, less than 90.69% of Go online submissions for LRU Cache.
// https://leetcode.com/submissions/detail/446744226/

type node struct {
	key, val   int
	prev, next *node
}

type linkedList struct {
	size       int
	head, tail *node
}

func newLinkedList() linkedList {
	n := node{}
	return linkedList{
		size: 0,
		head: &n,
		tail: &n,
	}
}

func (ll *linkedList) append(key, value int) {
	n := node{key: key, val: value, prev: ll.tail}
	ll.tail, ll.tail.next = &n, &n
	ll.size++
}

func (ll *linkedList) remove(n *node) {
	if n == nil {
		return
	}
	if n.next == nil {
		n.prev.next, ll.tail = nil, n.prev
	} else {
		n.prev.next, n.next.prev = n.next, n.prev
	}
	ll.size--
}

func (ll *linkedList) removeHead() *node {
	head := ll.head.next
	ll.remove(head)
	return head
}

func (ll *linkedList) moveToTail(n *node) {
	if n.next == nil {
		return
	}
	ll.remove(n)
	ll.tail, ll.tail.next, n.prev, n.next = n, n, ll.tail, nil
	ll.size++
}

// LRUCache ...
type LRUCache struct {
	capacity int
	hash     map[int]*node
	list     linkedList
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		hash:     make(map[int]*node),
		list:     newLinkedList(),
	}
}

// Get ...
func (lru *LRUCache) Get(key int) int {
	if n, ok := lru.hash[key]; ok {
		lru.list.moveToTail(n)
		return n.val
	}
	return -1
}

// Put ...
func (lru *LRUCache) Put(key int, value int) {
	if n, ok := lru.hash[key]; ok {
		lru.list.moveToTail(n)
		n.val = value
		return
	}
	if lru.list.size == lru.capacity {
		n := lru.list.removeHead()
		delete(lru.hash, n.key)
	}
	lru.list.append(key, value)
	lru.hash[key] = lru.list.tail
}
