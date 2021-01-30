package main

import "testing"

func buildTests() map[*ListNode]bool {
	res := make(map[*ListNode]bool, 3)
	a, b, c, d := &ListNode{3, nil}, &ListNode{2, nil}, &ListNode{0, nil}, &ListNode{4, nil}
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = b
	res[a] = true
	a, b = &ListNode{1, nil}, &ListNode{2, nil}
	a.Next = b
	b.Next = a
	res[a] = true
	a = &ListNode{1, nil}
	res[a] = false
	res[nil] = false
	a, b = &ListNode{1, nil}, &ListNode{2, nil}
	a.Next = b
	res[a] = false
	return res
}

func TestHasCycle(t *testing.T) {
	for input, output := range buildTests() {
		if value := hasCycle(input); value != output {
			t.Errorf("Expected hasCycle(%v) == %t, But got %t", input, output, value)
		}
	}
}
