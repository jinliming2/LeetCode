package main

import "testing"

type inputData struct {
	head *ListNode
	k    int
}

func TestReverseKGroup(t *testing.T) {
	for input, output := range map[*inputData][]int{
		{
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							&ListNode{
								5,
								nil,
							},
						},
					},
				},
			},
			2,
		}: {2, 1, 4, 3, 5},
		{
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							&ListNode{
								5,
								nil,
							},
						},
					},
				},
			},
			3,
		}: {3, 2, 1, 4, 5},
		{
			&ListNode{
				1,
				&ListNode{
					2,
					&ListNode{
						3,
						&ListNode{
							4,
							&ListNode{
								5,
								nil,
							},
						},
					},
				},
			},
			1,
		}: {1, 2, 3, 4, 5},
		{&ListNode{1, nil}, 1}: {1},
	} {
		value := reverseKGroup(input.head, input.k)
		circleDetect := value
		answer := []int{}
		for value != nil {
			answer = append(answer, value.Val)
			value = value.Next
			if circleDetect != nil && circleDetect.Next != nil {
				circleDetect = circleDetect.Next.Next
				if value == circleDetect {
					t.Errorf("reverseKGroup(%v, %d) returns circle link", input.head, input.k)
					break
				}
			}
		}
		if value != nil {
			continue
		}
		if len(output) != len(answer) {
			t.Errorf("Expected len(reverseKGroup(%v, %d)) == %d , But got %d", input.head, input.k, len(output), len(answer))
			continue
		}
		for i, num := range output {
			if answer[i] != num {
				t.Errorf("Expected reverseKGroup(%v, %d) == %v, But got %v", input.head, input.k, output, answer)
				break
			}
		}
	}
}
