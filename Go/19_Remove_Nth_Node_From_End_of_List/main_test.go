package main

import "testing"

type inputData struct {
	head *ListNode
	n    int
}

func TestRemoveNthFromEnd(t *testing.T) {
testCase:
	for input, output := range map[*inputData][]int{
		{&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2}: {1, 2, 3, 5},
		{&ListNode{1, nil}, 1}:               {},
		{&ListNode{1, &ListNode{2, nil}}, 1}: {1},
	} {
		value := removeNthFromEnd(input.head, input.n)
		for index, n := range output {
			if value == nil {
				t.Errorf("Expected removeNthFromEnd(%v, %d)[%d] == %d , But got nil", input.head, input.n, index, n)
				continue testCase
			}
			if value.Val != n {
				t.Errorf("Expected removeNthFromEnd(%v, %d)[%d] == %d , But got %d", input.head, input.n, index, n, value.Val)
				continue testCase
			}
			value = value.Next
		}
		if value != nil {
			t.Errorf("Expected removeNthFromEnd(%v, %d) == nil , But got %d", input.head, input.n, value.Val)
		}
	}
}
