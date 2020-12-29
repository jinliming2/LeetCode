package main

import (
	"strconv"
	"testing"
)

func TestMinDepth(t *testing.T) {
	for input, output := range map[*TreeNode]int{
		{
			Val: 3,
			Left: &TreeNode{
				Val: 9,
			},
			Right: &TreeNode{
				Val: 20,
				Left: &TreeNode{
					Val: 15,
				},
				Right: &TreeNode{
					Val: 7,
				},
			},
		}: 2,
		{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
				Right: &TreeNode{
					Val: 4,
					Right: &TreeNode{
						Val: 5,
						Right: &TreeNode{
							Val: 6,
						},
					},
				},
			},
		}: 5,
		nil: 0,
	} {
		if value := minDepth(input); value != output {
			tree := []string{}
			printTree(input, &tree)
			t.Errorf("Expected minDepth(%v) == %d, But got %d", tree, output, value)
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
