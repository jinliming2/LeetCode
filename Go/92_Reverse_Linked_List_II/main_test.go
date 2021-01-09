package main

import "testing"

type inputData struct {
	head *ListNode
	m    int
	n    int
}

func TestReverseBetween(t *testing.T) {
	for input, output := range map[*inputData][]int{
		{
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
			m: 2,
			n: 4,
		}: {1, 4, 3, 2, 5},
	} {
		value := reverseBetween(input.head, input.m, input.n)
		circleDetect := value
		answer := []int{}
		for value != nil {
			answer = append(answer, value.Val)
			value = value.Next
			if circleDetect != nil && circleDetect.Next != nil {
				circleDetect = circleDetect.Next.Next
				if value == circleDetect {
					t.Errorf("reverseBetween(%v, %d, %d) returns circle link", input.head, input.m, input.n)
					break
				}
			}
		}
		if value != nil {
			continue
		}
		if len(output) != len(answer) {
			t.Errorf("Expected len(reverseBetween(%v, %d, %d)) == %d , But got %d", input.head, input.m, input.n, len(output), len(answer))
			continue
		}
		for i, num := range output {
			if answer[i] != num {
				t.Errorf("Expected reverseBetween(%v, %d, %d) == %v, But got %v", input.head, input.m, input.n, output, answer)
				break
			}
		}
	}
}
