package main

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement Queue using Stacks.
// Memory Usage: 2 MB, less than 100.00% of Go online submissions for Implement Queue using Stacks.
// https://leetcode.com/submissions/detail/448980255/

type stack struct {
	data []int
}

func (s *stack) push(x int) {
	s.data = append(s.data, x)
}

func (s *stack) pop() (res int) {
	length := len(s.data)
	if length > 0 {
		res = s.data[length-1]
		s.data = s.data[:length-1]
	}
	return
}

func (s stack) peek() (res int) {
	length := len(s.data)
	if length > 0 {
		res = s.data[length-1]
	}
	return
}

func (s stack) empty() bool {
	return len(s.data) == 0
}

// MyQueue ...
type MyQueue struct {
	stack1, stack2 stack
}

// Constructor : Initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{
		stack1: stack{},
		stack2: stack{},
	}
}

// Push : Push element x to the back of queue.
func (mq *MyQueue) Push(x int) {
	mq.stack1.push(x)
}

// Pop : Removes the element from in front of queue and returns that element.
func (mq *MyQueue) Pop() int {
	mq.Peek()
	return mq.stack2.pop()
}

// Peek : Get the front element.
func (mq *MyQueue) Peek() int {
	if mq.stack2.empty() {
		for !mq.stack1.empty() {
			mq.stack2.push(mq.stack1.pop())
		}
	}
	return mq.stack2.peek()
}

// Empty : Returns whether the queue is empty.
func (mq *MyQueue) Empty() bool {
	return mq.stack1.empty() && mq.stack2.empty()
}
