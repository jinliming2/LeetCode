package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestInvertTree(t *testing.T) {
	for input, output := range map[*TreeNode]string{
		{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 3,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 5,
				Right: &TreeNode{
					Val: 6,
				},
			},
		}: "1 null 2 null 3 null 4 null 5 null 6 null null",
	} {
		flatten(input)
		var value []string
		if printTree(input, &value); strings.Join(value, " ") != output {
			t.Errorf("flatten got wrong answer %v", value)
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
