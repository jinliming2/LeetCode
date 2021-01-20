package main

import "testing"

func TestCountNodes(t *testing.T) {
	for input, output := range map[*TreeNode]int{
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 6,
				},
			},
		}: 6,
		nil:      0,
		{Val: 1}: 1,
	} {
		if value := countNodes(input); value != output {
			t.Errorf("Expected countNodes(%v) == %d, But got %d", input, output, value)
		}
	}
}
