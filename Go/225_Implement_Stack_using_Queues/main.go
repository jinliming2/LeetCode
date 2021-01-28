package main

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement Stack using Queues.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Implement Stack using Queues.
// https://leetcode.com/submissions/detail/448986850/

type node struct {
	val  int
	next *node
}

type queue struct {
	head, tail *node
	size       int
}

func (q *queue) enqueue(x int) {
	n := &node{val: x}
	if q.head == nil {
		q.head, q.tail = n, n
	} else {
		q.tail.next, q.tail = n, n
	}
	q.size++
}

func (q *queue) dequeue() (res int) {
	res, q.head = q.head.val, q.head.next
	q.size--
	return
}

func (q queue) peek() int {
	return q.head.val
}

func (q queue) empty() bool {
	return q.head == nil
}

// MyStack :
type MyStack struct {
	q queue
}

// Constructor : Initialize your data structure here.
func Constructor() MyStack {
	return MyStack{
		q: queue{},
	}
}

// Push :  Push element x onto stack.
func (ms *MyStack) Push(x int) {
	ms.q.enqueue(x)
}

// Pop : Removes the element on top of the stack and returns that element.
func (ms *MyStack) Pop() int {
	for size := ms.q.size; size > 1; size-- {
		ms.q.enqueue(ms.q.dequeue())
	}
	return ms.q.dequeue()
}

// Top : Get the top element.
func (ms *MyStack) Top() int {
	for size := ms.q.size; size > 1; size-- {
		ms.q.enqueue(ms.q.dequeue())
	}
	res := ms.q.dequeue()
	ms.q.enqueue(res)
	return res
}

// Empty : Returns whether the stack is empty.
func (ms *MyStack) Empty() bool {
	return ms.q.empty()
}
