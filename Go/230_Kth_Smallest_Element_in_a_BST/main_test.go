package main

import (
	"strconv"
	"testing"
)

type inputData struct {
	root *TreeNode
	k    int
}

func TestKthSmallest(t *testing.T) {
	for input, output := range map[*inputData]int{
		{&TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
				},
			},
			Right: &TreeNode{
				Val: 4,
			},
		}, 1}: 1,
		{&TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 6,
			},
		}, 3}: 3,
	} {
		if value := kthSmallest(input.root, input.k); value != output {
			var tree []string
			printTree(input.root, &tree)
			t.Errorf("Expected minWindow(%v, %d) == %d, But got %d", tree, input.k, output, value)
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
