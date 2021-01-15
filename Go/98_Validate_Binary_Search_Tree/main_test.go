package main

import "testing"

func TestIsValidBST(t *testing.T) {
	for input, output := range map[*TreeNode]bool{
		{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 3,
			},
		}: true,
		{
			Val: 5,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 6,
				},
			},
		}: false,
		{
			Val: 1,
			Left: &TreeNode{
				Val: 1,
			},
		}: false,
		{
			Val: 2147483647,
		}: true,
		{
			Val: -2147483648,
		}: true,
	} {
		if value := isValidBST(input); value != output {
			t.Errorf("Expected isValidBST(%v) == %t, But got %t", input, output, value)
		}
	}
}
