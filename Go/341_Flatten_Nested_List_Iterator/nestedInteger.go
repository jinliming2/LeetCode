package main

// NestedInteger : This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	val  int
	list []*NestedInteger
}

// IsInteger : Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return n.list == nil
}

// GetInteger : Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return n.val
}

// SetInteger : Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.val = value
	n.list = nil
}

// Add : Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.list = append(n.list, &elem)
}

// GetList : Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return n.list
}
