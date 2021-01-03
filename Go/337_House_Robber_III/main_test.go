package main

import (
	"strconv"
	"testing"
)

func TestRob(t *testing.T) {
	for input, output := range map[*TreeNode]int{
		{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 1,
				},
			},
		}: 7,
		{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 1,
				},
			},
		}: 9,
		{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
		}: 7,
		{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		}: 7,
	} {
		if value := rob(input); value != output {
			tree := []string{}
			printTree(input, &tree)
			t.Errorf("Expected rob(%v) == %d, But got %d", tree, output, value)
		}
	}
}

func printTree(root *TreeNode, res *[]string) {
	if root == nil {
		*res = append(*res, "null")
		return
	}
	*res = append(*res, strconv.Itoa(root.Val))
	printTree(root.Left, res)
	printTree(root.Right, res)
}
