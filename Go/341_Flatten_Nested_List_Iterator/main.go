package main

// Runtime: 4 ms, faster than 97.67% of Go online submissions for Flatten Nested List Iterator.
// Memory Usage: 5.3 MB, less than 69.77% of Go online submissions for Flatten Nested List Iterator.
// https://leetcode.com/submissions/detail/445075427/

// NestedIterator :
type NestedIterator struct {
	stack []*NestedInteger
}

// Constructor :
func Constructor(nestedList []*NestedInteger) *NestedIterator {
	res := &NestedIterator{make([]*NestedInteger, 0, len(nestedList))}
	for i := len(nestedList) - 1; i >= 0; i-- {
		res.stack = append(res.stack, nestedList[i])
	}
	return res
}

// Next :
func (ni *NestedIterator) Next() int {
	length := len(ni.stack)
	last := ni.stack[length-1]
	ni.stack = ni.stack[:length-1]
	return last.GetInteger()
}

// HasNext :
func (ni *NestedIterator) HasNext() bool {
	length := len(ni.stack)
	if length == 0 {
		return false
	}
	last := ni.stack[length-1]
	if last.IsInteger() {
		return true
	}
	ni.stack = ni.stack[:length-1]
	list := last.GetList()
	for i := len(list) - 1; i >= 0; i-- {
		ni.stack = append(ni.stack, list[i])
	}
	return ni.HasNext()
}
