package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestBstToGst(t *testing.T) {
	for input, output := range map[*TreeNode]string{
		{
			Val: 4,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 0,
				},
				Right: &TreeNode{
					Val: 2,
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			Right: &TreeNode{
				Val: 6,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 7,
					Right: &TreeNode{
						Val: 8,
					},
				},
			},
		}: "30 36 36 null null 35 null 33 null null 21 26 null null 15 null 8 null null",
		{
			Val: 0,
			Right: &TreeNode{
				Val: 1,
			},
		}: "1 null 1 null null",
		{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 2,
			},
		}: "3 3 null null 2 null null",
		{
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
		}: "7 9 10 null null null 4 null null",
	} {
		value := bstToGst(input)
		var answer []string
		printTree(value, &answer)
		if value := strings.Join(answer, " "); value != output {
			t.Errorf("Expected bstToGst(%v) == %s, But got %s", input, output, value)
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
