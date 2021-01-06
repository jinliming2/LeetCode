package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestInvertTree(t *testing.T) {
	for input, output := range map[*TreeNode]string{
		{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 1,
				},
				Right: &TreeNode{
					Val: 3,
				},
			},
			Right: &TreeNode{
				Val: 7,
				Left: &TreeNode{
					Val: 6,
				},
				Right: &TreeNode{
					Val: 9,
				},
			},
		}: "4 7 9 null null 6 null null 2 3 null null 1 null null",
	} {
		tree := invertTree(input)
		var value []string
		if printTree(tree, &value); strings.Join(value, " ") != output {
			t.Errorf("minDepth got wrong answer %v", value)
		}
	}
}

// TreeNode : Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
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
