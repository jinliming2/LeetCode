package main

// Runtime: 136 ms, faster than 56.59% of Go online submissions for LFU Cache.
// Memory Usage: 18.4 MB, less than 16.28% of Go online submissions for LFU Cache.
// https://leetcode.com/submissions/detail/447113358/

type node struct {
	val        int
	prev, next *node
}

type linkHashList struct {
	hash       map[int]*node
	head, tail *node
}

func newLinkHashList() *linkHashList {
	return &linkHashList{
		hash: make(map[int]*node),
	}
}

func (q *linkHashList) addTail(val int) {
	n := &node{val: val, prev: q.tail}
	if q.tail == nil {
		q.head, q.tail = n, n
	} else {
		q.tail.next, q.tail = n, n
	}
	q.hash[val] = n
}

func (q *linkHashList) removeHead() int {
	var n *node
	n, q.head = q.head, q.head.next
	if q.head == nil {
		q.tail = nil
	} else {
		q.head.prev = nil
	}
	delete(q.hash, n.val)
	return n.val
}

func (q *linkHashList) remove(val int) {
	n := q.hash[val]
	if n.prev == nil {
		q.removeHead()
	} else if n.next == nil {
		q.tail, q.tail.prev.next = q.tail.prev, nil
		delete(q.hash, val)
	} else {
		n.prev.next, n.next.prev = n.next, n.prev
		delete(q.hash, val)
	}
}

func (q *linkHashList) Empty() bool {
	return len(q.hash) == 0
}

// LFUCache ...
type LFUCache struct {
	kv       map[int]int
	kf       map[int]int
	fk       map[int]*linkHashList
	min      int
	capacity int
}

// Constructor ...
func Constructor(capacity int) LFUCache {
	return LFUCache{
		kv:       make(map[int]int),
		kf:       make(map[int]int),
		fk:       make(map[int]*linkHashList),
		min:      0,
		capacity: capacity,
	}
}

// Get ...
func (lfu *LFUCache) Get(key int) int {
	v, ok := lfu.kv[key]
	if !ok {
		return -1
	}
	f := lfu.kf[key]
	lfu.kf[key]++
	lfu.fk[f].remove(key)
	if lfu.fk[f].Empty() {
		delete(lfu.fk, f)
		if f == lfu.min {
			lfu.min++
		}
	}
	if lhl, ok := lfu.fk[f+1]; ok {
		lhl.addTail(key)
	} else {
		lhl = newLinkHashList()
		lfu.fk[f+1] = lhl
		lhl.addTail(key)
	}
	return v
}

// Put ...
func (lfu *LFUCache) Put(key int, value int) {
	if lfu.capacity == 0 {
		return
	}
	if v := lfu.Get(key); v == -1 {
		if len(lfu.kv) == lfu.capacity {
			lhl := lfu.fk[lfu.min]
			k := lhl.removeHead()
			delete(lfu.kv, k)
			delete(lfu.kf, k)
			if lhl.Empty() {
				delete(lfu.fk, lfu.min)
			}
		}
		lfu.kf[key] = 0
		lfu.min = 0
		if lhl, ok := lfu.fk[0]; ok {
			lhl.addTail(key)
		} else {
			lhl = newLinkHashList()
			lfu.fk[0] = lhl
			lhl.addTail(key)
		}
	}
	lfu.kv[key] = value
}
