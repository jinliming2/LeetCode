package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestFindDuplicateSubtrees(t *testing.T) {
	for input, output := range map[*TreeNode][]string{
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 4,
					},
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
		}: {"2 4 null null null", "4 null null"},
		{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
			},
			Right: &TreeNode{
				Val: 1,
			},
		}: {"1 null null"},
		{
			Val: 2,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
			},
		}: {"2 3 null null null", "3 null null"},
		{
			Val: 2,
			Left: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 11,
				},
			},
			Right: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 1,
				},
			},
		}: {},
	} {
		var tree []string
		printTree(input, &tree)
		value := findDuplicateSubtrees(input)
		if len(output) != len(value) {
			t.Errorf("Expected len(findDuplicateSubtrees(%v)) == %d , But got %d", tree, len(output), len(value))
			continue
		}
		answerList := make([]string, 0, len(value))
		for _, answer := range value {
			var res []string
			printTree(answer, &res)
			answerList = append(answerList, strings.Join(res, " "))
		}
	itemLoop:
		for _, item := range output {
		answerLoop:
			for _, answer := range answerList {
				if answer != item {
					continue answerLoop
				}
				continue itemLoop
			}
			t.Errorf("Expected solveNQueens(%v) == %v, But got %v", tree, output, value)
			break itemLoop
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
