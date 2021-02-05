package main

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	for input, output := range map[*ListNode][]int{
		{1, &ListNode{1, &ListNode{2, nil}}}:                             {1, 2},
		{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}: {1, 2, 3},
	} {
		value := deleteDuplicates(input)
		if len(output) == 0 && value != nil {
			t.Errorf("Expected len(deleteDuplicates(%v)) == 0", input)
			continue
		}
		p := value
		for index, item := range output {
			if p.Val != item {
				t.Errorf("Expected deleteDuplicates(%v)[%d] == %d, But got %d", input, index, item, p.Val)
				p = nil
				break
			}
			p = p.Next
		}
		if p != nil {
			t.Errorf("Expected deleteDuplicates(%v)[length] == nil, But got %d", input, p.Val)
		}
	}
}
