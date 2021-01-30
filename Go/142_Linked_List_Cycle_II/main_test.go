package main

import "testing"

func buildTests() map[*ListNode]*ListNode {
	res := make(map[*ListNode]*ListNode, 5)
	a, b, c, d := &ListNode{3, nil}, &ListNode{2, nil}, &ListNode{0, nil}, &ListNode{4, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = b
	res[a] = b
	a, b = &ListNode{1, nil}, &ListNode{2, nil}
	a.Next = b
	b.Next = a
	res[a] = a
	a = &ListNode{1, nil}
	res[a] = nil
	res[nil] = nil
	a, b = &ListNode{1, nil}, &ListNode{2, nil}
	a.Next = b
	res[a] = nil
	return res
}

func TestDetectCycle(t *testing.T) {
	for input, output := range buildTests() {
		if value := detectCycle(input); value != output {
			t.Errorf("Expected detectCycle(%v) == %v, But got %v", input, output, value)
		}
	}
}
